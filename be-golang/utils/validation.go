package utils

import (
	"github.com/go-playground/validator/v10"
	"be-golang/app/entity"
)

func ValidateLoginFields(user entity.User) error {
	validate := validator.New()

	if err := validate.Var(user.NIK, "required,len=16"); err != nil {
		return err
	}
	if err := validate.Var(user.Password, "required,min=6"); err != nil {
		return err
	}

	return nil
}
