package utils

import "github.com/go-playground/validator"

func ValidateStruct(req interface{}) error {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}
