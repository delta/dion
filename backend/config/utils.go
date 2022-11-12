package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
)

type envAndDefault struct {
	envName        string
	envVal         string
	defaultVal     string
	defaultOnEmpty bool
	notDefined     bool
}

func getEnvAndDefaultValue(s []byte, lineNo *int, charNo *int, i *int) (*envAndDefault, error) {
	startLine := *lineNo
	startChar := *charNo
	insideBraces := []byte{}
	broken := false
	for ; *i < len(s); *i++ {
		if s[*i] == '}' {
			*charNo++
			broken = true
			break
		} else {
			if s[*i] == '\n' {
				*lineNo++
				*charNo = -1
			}
			*charNo++
			insideBraces = append(insideBraces, s[*i])
		}
	}
	if !broken {
		return nil, fmt.Errorf("unclosed '{' brace[line:%d,char:%d]", startLine, startChar)
	}
	envVarName, defaultValue, found := bytes.Cut(insideBraces, []byte(":"))

	envVal, foundEnv := os.LookupEnv(string(envVarName))
	if !found {
		return &envAndDefault{
			envName:        string(envVarName),
			envVal:         envVal,
			defaultVal:     "",
			defaultOnEmpty: true,
			notDefined:     !foundEnv,
		}, nil
	} else {
		defaultOnEmpty := false
		if len(defaultValue) > 0 && defaultValue[0] == '-' {
			defaultValue = defaultValue[1:]
			defaultOnEmpty = true
		}
		return &envAndDefault{
			envName:        string(envVarName),
			envVal:         envVal,
			defaultVal:     string(defaultValue),
			defaultOnEmpty: defaultOnEmpty,
			notDefined:     !foundEnv,
		}, nil
	}
}

func handleEnvVarAndDefaultValues(s []byte) []byte {
	result := []byte{}
	lineNo := 1
	charNo := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lineNo++
			charNo = -1
		}
		charNo += 1
		if s[i] == '$' {
			// For escaping the $,it must be \$
			if i != 0 && s[i-1] == '\\' {
				result[len(result)-1] = '$'
				continue
			}
			i++
			if len(s) == i || s[i] != '{' {
				panic("Expected { after $, If you want $, escape it")
			}
			i++
			envAndDef, err := getEnvAndDefaultValue(s, &lineNo, &charNo, &i)
			if err != nil {
				panic(err)
			}
			if len(envAndDef.envVal) == 0 && len(envAndDef.defaultVal) == 0 {
				log.New(os.Stderr, "[Config Warning]", log.LstdFlags).Printf(
					"Env %s has both env variable empty and default value empty",
					envAndDef.envName,
				)
			}
			if envAndDef.notDefined {
				result = append(result, []byte(envAndDef.defaultVal)...)
			} else if len(envAndDef.envVal) == 0 && envAndDef.defaultOnEmpty {
				result = append(result, []byte(envAndDef.defaultVal)...)
			} else {
				result = append(result, []byte(envAndDef.envVal)...)
			}
		} else {
			result = append(result, s[i])
		}
	}
	return result
}

func unmarshallYaml(file string) (map[string]interface{}, error) {
	yamlContent, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// yamlContent = []byte(os.ExpandEnv(string(yamlContent)))
	yamlContent = handleEnvVarAndDefaultValues(yamlContent)
	res := make(map[string]interface{})
	if err := yaml.Unmarshal(yamlContent, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// if both are maps, it'll perform merging,else it'll return extraConf.
// This is because we want the more specific configuration to override
// the base config
func mergeIfBothMaps(baseConf interface{}, extraConf interface{}) interface{} {
	extraConfAsMap, _ := extraConf.(map[string]interface{})
	baseConfAsMap, ok := baseConf.(map[string]interface{})
	if !ok {
		return extraConfAsMap
	} else {
		return merge(baseConfAsMap, extraConfAsMap)
	}
}

func merge(baseConfig map[string]interface{}, extraConfig map[string]interface{}) map[string]interface{} {
	for key, val := range extraConfig {
		_, ok := baseConfig[key]
		if ok {
			_, ok := val.(map[string]interface{})
			if ok {
				baseConfig[key] = mergeIfBothMaps(baseConfig[key], val)
			} else {
				baseConfig[key] = val
			}
		} else {
			baseConfig[key] = val
		}
	}
	return baseConfig
}

func loadConfig() (*Config, error) {
	// ENV should be only dev|prod|test
	possible := []string{"dev", "prod", "test", ""}

	env := string(bytes.ToLower([]byte(os.Getenv("ENV"))))

	valid := false
	for _, elem := range possible {
		valid = (valid || (env == elem))
	}
	if !valid {
		panic("ENV environment variable should be one of prod|dev|test")
	}

	if len(env) == 0 {
		env = "dev"
	}

	_, filename, _, _ := runtime.Caller(1)
	envSpecifiConfiFileName := path.Join(path.Dir(filename), fmt.Sprintf("config.%s.yaml", env))
	baseConfigFileName := path.Join(path.Dir(filename), "config.base.yaml")

	baseConfig, err := unmarshallYaml(baseConfigFileName)
	if err != nil {
		return nil, err
	}
	envSpecifiConfig, err := unmarshallYaml(envSpecifiConfiFileName)
	if err != nil {
		return nil, err
	}
	mergedConfig := merge(baseConfig, envSpecifiConfig)
	conf := Config{}
	marshaledMergedConf, err := yaml.Marshal(mergedConfig)
	if err != nil {
		return nil, err
	}
	// decoder := yaml.NewDecoder(strings.NewReader(string(marshaledMergedConf)))
	decoder := yaml.NewDecoder(bytes.NewReader(marshaledMergedConf))
	decoder.KnownFields(true)
	if err = decoder.Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
