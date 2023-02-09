package config

import (
	configdev "PTH-IT/api_golang/config/dev"
	configLocal "PTH-IT/api_golang/config/local"
	configprod "PTH-IT/api_golang/config/production"
	configStg "PTH-IT/api_golang/config/stag"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env    string       `json:"env"`
	Port   string       `json:"port"`
	Mysql  MysqlConfig  `json:"mysql"`
	Monggo MonggoConfig `json:"monggodb"`
	Redis  RedisConfig  `json:"redis"`
}
type MysqlConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"password"`
	Db   string `json:"db"`
}
type MonggoConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Pass string `json:"password"`
	Db   string `json:"db"`
}
type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Pass string `json:"password"`
	Db   int    `json:"db"`
}

func Getconfig() AppConfig {
	var appConfig AppConfig
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		err := json.Unmarshal([]byte(configLocal.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "dev" {
		err := json.Unmarshal([]byte(configdev.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "stg" {
		err := json.Unmarshal([]byte(configStg.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	} else if env == "prod" {
		err := json.Unmarshal([]byte(configprod.ConfigApp), &appConfig)
		if err != nil {
			panic(err)
		}
	}
	if appConfig.Mysql.User == "" {
		appConfig.Mysql.User = os.Getenv("DB_USER")

	}
	if appConfig.Mysql.Pass == "" {
		appConfig.Mysql.Pass = os.Getenv("DB_PASSWORD")

	}
	if appConfig.Mysql.Host == "" {
		appConfig.Mysql.Host = os.Getenv("DB_HOST")

	}
	if appConfig.Mysql.Port == "" {
		appConfig.Mysql.Port = os.Getenv("DB_PORT")

	}
	if appConfig.Mysql.Db == "" {
		appConfig.Mysql.Db = os.Getenv("DB_NAME")

	}
	return appConfig
}
