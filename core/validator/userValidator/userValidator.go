package uservalidator

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/marc/go-clean-example/core/dto"
)

func ValidateUserInsert(userRequest *dto.CreateUserRequest) error {

	if strings.Trim(userRequest.Name, "") == "" {
		return errors.New("Username is required")
	}

	if strings.Trim(userRequest.Email, "") == "" {
		return errors.New("User email is required")
	}

	if erro := checkmail.ValidateFormat(userRequest.Email); erro != nil {
		return errors.New("The email entered is invalid")
	}

	if strings.Trim(userRequest.Password, "") == "" {
		return errors.New("Password is required")
	}

	return nil

}

func ValidateUserUpdate(userRequest *dto.UpdateUserRequest) error {

	if userRequest.IDUser <= 0 {
		return errors.New("User ID is required")
	}

	if strings.Trim(userRequest.Name, "") == "" {
		return errors.New("Username is required")
	}

	if strings.Trim(userRequest.Email, "") == "" {
		return errors.New("User email is required")
	}

	if erro := checkmail.ValidateFormat(userRequest.Email); erro != nil {
		return errors.New("The email entered is invalid")
	}

	if strings.Trim(userRequest.Password, "") == "" {
		return errors.New("Password is required")
	}

	return nil

}
