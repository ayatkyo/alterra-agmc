package utils

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}

	return nil
}

var EchoCustomValidator *CustomValidator = &CustomValidator{
	validator: validator.New(),
}
