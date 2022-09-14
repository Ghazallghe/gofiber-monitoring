package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) map[string][]string {
	var errors map[string][]string
	err := validate.Struct(s)
	if err != nil {
		errors = make(map[string][]string)
		for _, err := range err.(validator.ValidationErrors) {
			var message string
			if err.Param() == "" {
				message = err.Tag()
			} else {
				message = err.ActualTag() + " Should be " + err.Param()
			}
			field := strings.ToLower(err.Field())
			errors[field] = append(errors[field], message)
		}
	}
	return errors
}
