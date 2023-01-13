package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(config.ServerUrl)
	fmt.Println(config.AppPort)
}

type Config struct {
	ServerUrl string `json:"server_url" validate:"required"`
	AppPort   int    `json:"app_port" validate:"required,numeric,gte=8080,lte=8085"`
}

func LoadConfig() (*Config, error) {
	configFile, err := os.Open("configuration/config.json")
	if err != nil {
		fmt.Errorf("Could not open file config.json : ", err.Error())
	}

	decoder := json.NewDecoder(configFile)
	config := Config{}

	decodeErr := decoder.Decode(&config)
	if decodeErr != nil {
		fmt.Errorf("Could not decode config : ", decodeErr.Error())
	}

	if !Validate(config) {
		return nil, errors.New("Invalid config !")
	}
	return &config, nil
}

func Validate(config Config) bool {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		fmt.Println("Invalid config !")
		for _, validationErr := range err.(validator.ValidationErrors) {
			fmt.Println(validationErr.StructNamespace() + " violated " + validationErr.Tag() + " validation.")
		}
		return false
	}
	return true
}
