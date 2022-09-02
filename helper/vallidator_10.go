package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// func FormatValidationError(err error) []string {
// 	var errors []string
// 	for _, e := range err.(validator.ValidationErrors) {
// 		errors = append(errors, e.Error())
// 	}
// 	return errors
// }
func FormatValidationError(err error) string {
	var errors string
	for _, e := range err.(validator.ValidationErrors) {
		errors = fmt.Sprintf("Field is %s", e.ActualTag())
	}
	return errors
}