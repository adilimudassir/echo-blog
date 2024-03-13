package helpers

import (
	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors formats validation errors into a JSON object with each field mapped to an array of error messages
func FormatValidationErrors(err error) map[string][]string {
	errMsg := make(map[string][]string)
	for _, e := range err.(validator.ValidationErrors) {
		errMsg[e.Field()] = append(errMsg[e.Field()], e.Tag())
	}
	return errMsg
}
