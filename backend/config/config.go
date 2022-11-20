//		config.yaml variable meanings
//
//	 - ${ENV} => This means we'll check environment variable "ENV" and set the value to whatever it is.
//	   If it is empty or unset, it'll be set to that
//
//	 - ${ENV:dev} => This means if environment variable "ENV" is not set, we'll use the default
//	   value i.e "dev"
//
//	 - ${ENV:-dev} => This means if environment variable "ENV" is not set or is set to
//	   empty string,we'll use the default value "dev"
//
//	   For further information ,see the tests or understand the source code
//	  or go through the config_test.go to understand how config works
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

type SessionConfig struct {
	SecretKey  string `yaml:"secret_key"`
	ExpiryTime int    `yaml:"expiry_time"`
	Domain     string `yaml:"domain"`
}

type DAuthConfig struct {
	ClientSecret string `yaml:"client_secret"`
	ClientId     string `yaml:"client_id"`
	RedirectURL  string `yaml:"redirect_url"`
}

type Config struct {
	Db          DbConfig      `yaml:"db"`
	Environment string        `yaml:"environment"`
	FrontendUrl string        `yaml:"frontend_url"`
	Server      ServerConfig  `yaml:"server"`
	OauthConfig DAuthConfig   `yaml:"oauth"`
	Session     SessionConfig `yaml:"session"`
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
	fmt.Printf("%#v\n", C)
}
