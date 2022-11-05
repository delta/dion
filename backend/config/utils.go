package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func unmarshallYaml(file string) (map[string]interface{}, error) {
	yamlContent, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	yamlContent = []byte(os.ExpandEnv(string(yamlContent)))
	res := make(map[string]interface{})
	if err := yaml.Unmarshal(yamlContent, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func combine(baseConf interface{}, extraConf interface{}) interface{} {
	extraConfAsMap, _ := extraConf.(map[string]interface{})
	baseConfAsMap, ok := baseConf.(map[string]interface{})
	if !ok {
		return baseConfAsMap
	} else {
		return mergeConfigs(baseConfAsMap, extraConfAsMap)
	}
}

func mergeConfigs(baseConfig map[string]interface{}, extraConfig map[string]interface{}) map[string]interface{} {
	for key, val := range extraConfig {
		_, ok := baseConfig[key]
		if ok {
			_, ok := val.(map[string]interface{})
			if ok {
				baseConfig[key] = combine(baseConfig[key], val)
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

	env := os.Getenv("ENV")

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

	envSpecifiConfiFileName := filepath.Join("config", fmt.Sprintf("config.%s.yaml", env))
	baseConfigFileName := filepath.Join("config", "config.base.yaml")
	// fmt.Println(os.Getwd())
	// fmt.Println(envSpecifiConfiFileName, baseConfigFileName)
	// envSpecifiConfiFileName = fmt.Sprintf("%s/config/config.%s.yaml", basePath, env)
	// baseConfigFileName := fmt.Sprintf("%s/config/config.base.yaml", basePath)

	baseConfig, err := unmarshallYaml(baseConfigFileName)
	if err != nil {
		return nil, err
	}
	envSpecifiConfig, err := unmarshallYaml(envSpecifiConfiFileName)
	if err != nil {
		return nil, err
	}
	mergedConfig := mergeConfigs(baseConfig, envSpecifiConfig)
	conf := Config{}
	marshalledMergedConf, err := yaml.Marshal(mergedConfig)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(marshalledMergedConf, &conf)

	return &conf, nil
}
