package config

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
	"github.com/ogunb/matchday-functions/fixture/model"
)

func GetConfig() *model.Config {
	yamlFile, err := ioutil.ReadFile("env.yaml")

	if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
	}

	config := &model.Config{}
	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
			log.Fatalf("Unmarshal: %v", err)
	}

	return config
}