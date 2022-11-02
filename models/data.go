package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Data struct {
	Host            string `yaml:"MINIO_HOST"`
	Port            string `yaml:"MINIO_PORT"`
	AccessKeyID     string `yaml:"MINIO_USERNAME"`
	SecretAccessKey string `yaml:"MINIO_PASSWORD"`
	UseSSL          bool   `yaml:"MINIO_USESSL"`
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
