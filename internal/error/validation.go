package error

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type ErrorField struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

func (e ErrorField) Error() string {
	return fmt.Sprintf("error on field %s with %v", e.Field, e.Errors)
}

// Validate returns a list of invalid fields or nil
func Validate(entity interface{}) []ErrorField {
	err := validate.Struct(entity)
	invalidFields := make([]ErrorField, 0)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			fieldWithError := ErrorField{}
			fieldWithError.Field = strings.ToLower(validationError.StructField())
			fieldWithError.Errors = []string{processError(validationError.Error())}
			invalidFields = append(invalidFields, fieldWithError)
		}
	}
	return invalidFields
}

func processError(err string) string {
	parts := strings.Split(err, "Error:")
	return parts[1]
}
