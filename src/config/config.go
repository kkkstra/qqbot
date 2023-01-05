package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type config struct {
	App    app    `yaml:"app"`
	Cqhttp cqhttp `yaml:"cqhttp"`
	Bot    bot    `yaml:"bot"`
}

type app struct {
	Addr string `yaml:"addr"`
}

type cqhttp struct {
	Addr string `yaml:"addr"`
}

type bot struct {
	Qq   string `yaml:"qq"`
	Name string `yaml:"name"`
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
