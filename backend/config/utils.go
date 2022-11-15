package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

type envAndDefault struct {
	envVarName        string
	envVarVal         string
	defaultVal        string
	setDefaultOnEmpty bool
	notDefined        bool
}

type parseState struct {
	src       []byte
	curLineNo int
	curCharNo int
	curIndex  int
}

// After we encounter a ${ in the yaml, we let this function take over.
// This will take in the current state, consume the yaml until we reach the '}'.
// Semantics:
//
//   - ${ENV} => This means we'll check environment variable "ENV" and set the value to whatever it is.
//     If it is empty or unset, it'll be set to that
//
//   - ${ENV:dev} => This means if environment variable "ENV" is not set, we'll use the default
//     value i.e "dev"
//
//   - ${ENV:-dev} => This means if environment variable "ENV" is not set or is set to
//     empty string,we'll use the default value "dev"
//
//     For further information ,see the tests or understand the source code
func parseEnvVarAndDefault(state *parseState) (*envAndDefault, error) {
	startLine := state.curLineNo
	startChar := state.curCharNo
	insideBraces := []byte{}
	broken := false
	for ; state.curIndex < len(state.src); state.curIndex++ {
		if state.src[state.curIndex] == '}' {
			state.curCharNo++
			broken = true
			break
		} else {
			if state.src[state.curIndex] == '\n' {
				state.curLineNo++
				state.curCharNo = -1
			}
			state.curCharNo++
			insideBraces = append(insideBraces, state.src[state.curIndex])
		}
	}
	if !broken {
		return nil, fmt.Errorf("unclosed '{' brace[line:%d,char:%d]", startLine, startChar)
	}
	envVarName, defaultValue, found := bytes.Cut(insideBraces, []byte(":"))

	envVal, foundEnv := os.LookupEnv(string(envVarName))
	if !found {
		return &envAndDefault{
			envVarName:        string(envVarName),
			envVarVal:         envVal,
			defaultVal:        "",
			setDefaultOnEmpty: true,
			notDefined:        !foundEnv,
		}, nil
	} else {
		defaultOnEmpty := false
		if len(defaultValue) > 0 && defaultValue[0] == '-' {
			defaultValue = defaultValue[1:]
			defaultOnEmpty = true
		}
		return &envAndDefault{
			envVarName:        string(envVarName),
			envVarVal:         envVal,
			defaultVal:        string(defaultValue),
			setDefaultOnEmpty: defaultOnEmpty,
			notDefined:        !foundEnv,
		}, nil
	}
}

// This function can panic
func handleEnvVarAndDefaultValues(src []byte) []byte {
	result := []byte{}
	state := parseState{
		src:       src,
		curLineNo: 1,
		curCharNo: 0,
		curIndex:  0,
	}
	for ; state.curIndex < len(state.src); state.curIndex++ {
		if state.src[state.curIndex] == '\n' {
			state.curLineNo++
			state.curCharNo = -1
		}
		state.curCharNo += 1
		if state.src[state.curIndex] == '$' {
			// For escaping the $,it must be \$
			if state.curIndex != 0 && state.src[state.curIndex-1] == '\\' {
				result[len(result)-1] = '$'
				continue
			}
			state.curIndex++
			if len(state.src) == state.curIndex || state.src[state.curIndex] != '{' {
				panic("Expected { after $, If you want $, escape it")
			}
			state.curIndex++
			envAndDef, err := parseEnvVarAndDefault(&state)
			if err != nil {
				panic(err)
			}
			if len(envAndDef.envVarVal) == 0 && len(envAndDef.defaultVal) == 0 {
				log.New(os.Stderr, "[Config Warning]", log.LstdFlags).Printf(
					"Env %s has both env variable empty and default value empty",
					envAndDef.envVarName,
				)
			}
			if envAndDef.notDefined {
				result = append(result, []byte(envAndDef.defaultVal)...)
			} else if len(envAndDef.envVarVal) == 0 && envAndDef.setDefaultOnEmpty {
				result = append(result, []byte(envAndDef.defaultVal)...)
			} else {
				result = append(result, []byte(envAndDef.envVarVal)...)
			}
		} else {
			result = append(result, state.src[state.curIndex])
		}
	}
	return result
}

func panicIfNotValidENV() {
	//
	// ENV should be only dev|prod|test
	possible := []string{"dev", "prod", "test", ""}

	env := strings.ToLower(os.Getenv("ENV"))

	valid := false
	for _, elem := range possible {
		valid = (valid || (env == elem))
	}
	if !valid {
		panic("ENV environment variable should be one of prod|dev|test")
	}
}

func loadConfig() (*Config, error) {
	panicIfNotValidENV()

	_, filename, _, _ := runtime.Caller(1)
	configFileName := path.Join(path.Dir(filename), "config.yaml")

	yamlContent, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	yamlContent = handleEnvVarAndDefaultValues(yamlContent)

	conf := Config{}
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(yamlContent))
	decoder.KnownFields(true)
	if err = decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
