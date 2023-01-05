package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type config struct {
	App        app        `yaml:"app"`
	Cqhttp     cqhttp     `yaml:"cqhttp"`
	Postgresql postgresql `yaml:"postgresql"`
	Bot        bot        `yaml:"bot"`
	ChatGpt    chatgpt    `yaml:"chatgpt"`
}

type app struct {
	Addr string `yaml:"addr"`
}

type cqhttp struct {
	Addr string `yaml:"addr"`
}

type postgresql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Port     string `yaml:"port"`
}

type bot struct {
	Qq   string `yaml:"qq"`
	Name string `yaml:"name"`
}

type chatgpt struct {
	Token   string `yaml:"token"`
	Cfvalue string `json:"cfvalue"`
}

var C *config

func initConfig() {
	configFile := "config.yaml"
	r, err := os.ReadFile(fmt.Sprintf("./env/config/%s", configFile))
	if err != nil {
		log.Println(err)
		return
	}
	con := &config{}
	err = yaml.Unmarshal(r, con)
	if err != nil {
		log.Println(err)
		return
	}
	C = con
}
