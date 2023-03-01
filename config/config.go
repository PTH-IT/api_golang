package config

import (
	"encoding/json"
	"log"
	"os"

	configdev "github.com/PTH-IT/api_golang/config/dev"
	configLocal "github.com/PTH-IT/api_golang/config/local"
	configprod "github.com/PTH-IT/api_golang/config/production"
	configStg "github.com/PTH-IT/api_golang/config/stag"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env    string       `json:"env"`
	Port   string       `json:"port"`
	Mysql  MysqlConfig  `json:"mysql"`
	Monggo MonggoConfig `json:"monggodb"`
	Redis  RedisConfig  `json:"redis"`
	Aws    AwsConfig
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
	Db   string `json:"db"`
}
type AwsConfig struct {
	Host   string
	Port   string
	Region string
	Id     string
	Secret string
	Token  string
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
	appConfig.Aws.Host = os.Getenv("AWS_HOST")
	appConfig.Aws.Port = os.Getenv("AWS_PORT")
	appConfig.Aws.Region = os.Getenv("AWS_REGION")
	appConfig.Aws.Id = os.Getenv("AWS_ID")
	appConfig.Aws.Secret = os.Getenv("AWS_SECRET")
	appConfig.Aws.Token = os.Getenv("AWS_TOKEN")
	//MYSQL
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
	//MONGGODB
	if appConfig.Monggo.Host == "" {
		appConfig.Monggo.Host = os.Getenv("MONGGO_HOST")

	}
	if appConfig.Monggo.Host == "" {
		appConfig.Monggo.User = os.Getenv("MONGGO_USER")

	}
	if appConfig.Monggo.Host == "" {
		appConfig.Monggo.Pass = os.Getenv("MONGGO_PASSWORD")

	}
	if appConfig.Monggo.Host == "" {
		appConfig.Monggo.Db = os.Getenv("MONGGO_DB")

	}
	//REDIS
	if appConfig.Redis.Host == "" {
		appConfig.Redis.Host = os.Getenv("REDIS_HOST")

	}
	if appConfig.Redis.Port == "" {
		appConfig.Redis.Port = os.Getenv("REDIS_PORT")

	}
	if appConfig.Redis.Pass == "" {
		appConfig.Redis.Pass = os.Getenv("REDIS_PASSWORD")

	}
	if appConfig.Redis.Db == "" {
		appConfig.Redis.Db = os.Getenv("REDIS_DB")

	}
	return appConfig
}
