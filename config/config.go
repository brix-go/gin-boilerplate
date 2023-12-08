package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ConfigModel struct {
	AppConfig struct {
		Name   string `envconfig:"APP_NAME"`
		Host   string `envconfig:"APP_HOST"`
		Port   string `envconfig:"APP_PORT"`
		Secret string `envconfig:"APP_SECRET"`
	} ``

	Database struct {
		Driver   string `envconfig:"DB_DRIVER"`
		Host     string `envconfig:"DB_HOST"`
		Port     int    `envconfig:"DB_PORT"`
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
		Database string `envconfig:"DB_NAME"`
	} ``

	Log struct {
		Host     string `envconfig:"LOG_HOST"`
		Port     int    `envconfig:"LOG_PORT"`
		Username string `envconfig:"LOG_USER"`
		Password string `envconfig:"LOG_PASS"`
		Index    string `envconfig:"LOG_INDEX"`
	} `mapstructure:"log"`

	ErrorContract struct {
		JSONPathFile string `envconfig:"ERROR_CONTRACT_JSON_PATH"`
	} `mapstructure:"errorContract"`

	Kafka struct {
		Host    string `envconfig:"KAFKA_HOST"`
		GroupId string `envconfig:"KAFKA_GROUP_ID"`
		Reset   string `envconfig:"KAFKA_RESET"`
	}
	Redis struct {
		Host string `envconfig:"REDIS_HOST"`
		Port int    `envconfig:"REDIS_PORT"`
	}
}

var Config ConfigModel

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = envconfig.Process(".env", &Config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("APP ADDRESS : ", Config.AppConfig.Host)
}
