package config

import (
	"io/ioutil"

	"go.uber.org/fx"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type Config struct {
	Db          *DbConfig `yaml:"db"`
	Environment string    `yaml:"environment"`
}

func New() (*Config, error) {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}
	conf := Config{}
	err = yaml.Unmarshal(yamlFile, &conf)
	return &conf, err
}

var Module = fx.Provide(
	New,
)
