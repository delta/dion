package config

import (
	"fmt"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type ServerConfig struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"readtimeout"`
	WriteTimeout int `yaml:"writetimeout"`
}
type Config struct {
	Db          DbConfig     `yaml:"db"`
	Environment string       `yaml:"environment"`
	Server      ServerConfig `yaml:"server"`
}

var C Config

func init() {
	fmt.Println("== SETTING UP CONFIG ==")
	conf, err := loadConfig()
	if err != nil {
		errMsg := fmt.Errorf("unable to load config due to %+v", err)
		panic(errMsg)
	}
	C = *conf
}
