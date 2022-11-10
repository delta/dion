package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"gopkg.in/yaml.v3"
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
	ReadTimeout  int    `yaml:"readtimeout"`
	WriteTimeout int    `yaml:"writetimeout"`
}
type Config struct {
	Db          *DbConfig     `yaml:"db"`
	Environment string        `yaml:"environment"`
	Server      *ServerConfig `yaml:"server"`
}

var (
	C Config
)

func init() {
	fmt.Println("==> SETTING UP CONFIG")
	if err := loadConfig(); err != nil {
		errMsg := fmt.Errorf("unable to load config due to %+v."+
			"Check if the you have have created config.yaml with the"+
			" correct keys", err)
		panic(errMsg)
	}
}

func loadConfig() error {
	// fmt.Println(runtime.Caller(1))
	_, filename, _, _ := 	runtime.Caller(1)
	filePath := path.Join(path.Dir(filename), "config.yaml")
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	conf := Config{}
	err = yaml.Unmarshal(yamlFile, &conf)
	C = conf
	return err
}
