package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidations() {
	Validate = validator.New()

	Validate.RegisterValidation("validPass", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 6 || len(password) > 20 {
			return false
		}

		upper := regexp.MustCompile(`[A-Z]`)
		lower := regexp.MustCompile(`[a-z]`)
		number := regexp.MustCompile(`[0-9]`)
		symbol := regexp.MustCompile(`[\W_]`)

		return upper.MatchString(password) && lower.MatchString(password) && number.MatchString(password) && symbol.MatchString(password)
	})
}
