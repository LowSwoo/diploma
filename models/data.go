package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Data struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	AccessKeyID     string `yaml:"username"`
	SecretAccessKey string `yaml:"password"`
	UseSSL          bool   `yaml:"useSSL"`
}

func (d *Data) GetData() *Data {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		log.Fatal(err.Error())
	}
	return d
}
