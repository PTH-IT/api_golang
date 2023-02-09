package config

import (
	configdev "PTH-IT/api_golang/config/dev"
	configLocal "PTH-IT/api_golang/config/local"
	configprod "PTH-IT/api_golang/config/production"
	configStg "PTH-IT/api_golang/config/stag"
	"encoding/json"
	"fmt"
)

type AppConfig struct {
	Env  string `json:"env"`
	Port string `json:"port"`
}

func Getconfig() AppConfig {
	var appConfig AppConfig
	env := "local"
	if env == "local" {
		err := json.Unmarshal([]byte(configLocal.ConfigApp), &appConfig)
		if err != nil {
			fmt.Println(err)
		}
	} else if env == "dev" {
		err := json.Unmarshal([]byte(configdev.ConfigApp), &appConfig)
		if err != nil {
			fmt.Println(err)
		}
	} else if env == "stg" {
		err := json.Unmarshal([]byte(configStg.ConfigApp), &appConfig)
		if err != nil {
			fmt.Println(err)
		}
	} else if env == "prod" {
		err := json.Unmarshal([]byte(configprod.ConfigApp), &appConfig)
		if err != nil {
			fmt.Println(err)
		}
	}
	return appConfig
}
