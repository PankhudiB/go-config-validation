package configuration

import (
	"fmt"
	"github.com/go-playground/validator"
	"strings"
)

type Config struct {
	ServerUrl string `validate:"is_https"`
	AppPort   int
}

func CustomValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if !strings.HasPrefix(value, "https://") {
		return false
	}
	return true
}

func Validate(config Config) bool {
	validate := validator.New()

	err := validate.RegisterValidation("is_https", CustomValidation)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	validationErr := validate.Struct(config)
	if validationErr != nil {
		for _, validationErr := range validationErr.(validator.ValidationErrors) {
			fmt.Println(validationErr.StructNamespace() + " violated " + validationErr.Tag() + " validation.")
		}
		return false
	}
	return true
}
