package localerrors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		validationError := validationErrors[0]

		lowerField := strings.ToLower(validationError.StructField())
		switch validationError.Tag() {
		case "required":
			return errors.New(fmt.Sprint(lowerField, " is required"))
		case "min":
			return errors.New(fmt.Sprint(lowerField, " requires a minimum of ", validationError.Param()))
		case "max":
			return errors.New(fmt.Sprint(lowerField, " requires a maximum of ", validationError.Param()))
		case "email":
			return errors.New(fmt.Sprint(lowerField, " is not a valid email"))
		}
	}
	return nil
}
