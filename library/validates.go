package library

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

//ValidateData func
func ValidateData(data interface{}) (HTTPError, error) {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(data)
	errData := HTTPError{}

	return errData, err
}

//CustomValidateMessage func
func CustomValidateMessage(report HTTPError, err error) HTTPError {
	errorMessage := make(map[string]interface{}, 0)
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				errorMessage[err.Field()] = fmt.Sprintf("%s is required", err.Field())
				report.Message = errorMessage
			case "email":
				errorMessage[err.Field()] = fmt.Sprintf("%s is not valid email", err.Field())
				report.Message = errorMessage
			case "gte":
				errorMessage[err.Field()] = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				report.Message = errorMessage
			case "lte":
				errorMessage[err.Field()] = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				report.Message = errorMessage
			}
		}
	}

	return report
}
