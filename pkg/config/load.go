package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/joho/godotenv"
)

var Config *config

func init() {
	godotenv.Load()
	var err error
	Config, err = LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
}

func LoadConfig() (*config, error) {
	var configBuffer []byte
	var err error

	configBuffer, err = os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Failed to Read Config:" + err.Error())
		return nil, err
	}
	var config config
	err = yaml.Unmarshal(configBuffer, &config)
	if err != nil {
		log.Fatal("Failed to Unmarshal Config:" + err.Error())
		return nil, err
	}

	return &config, nil
}
