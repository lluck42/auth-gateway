package bean

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Config struct {
	App struct {
		Addr    string `yaml:"addr"`
		SetMode string `yaml:"setmode"`
	} `yaml:"app"`
	Mysql struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Database  string `yaml:"database"`
		ParseTime string `yaml:"parsetime"`
		Debug     bool   `yaml:"debug"`
	} `yaml:"mysql"`
	ReverseProxy struct {
		Scheme string `yaml:"Scheme"`
		Host   string `yaml:"Host"`
	} `yaml:"ReverseProxy"`
}

func ConfigInit() {

	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err.Error())
	}
}
