package dto

import (
	"encoding/json"
	"io"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// CreateProductRequest is an representation request body to create a new User
type CreateUserRequestDTO struct {
	Name     string    `json:name`
	Email    string    `json:email`
	Password string    `json:password`
	CreateAt time.Time `json:createdAt`
}

// ValidateCreateUserRequest valid the rules on propertys
func (u CreateUserRequestDTO) ValidateCreateUserRequest() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Email, validation.Required, is.Email))
}

// UpdateUserRequestDTO is an representation request body to update a User
type UpdateUserRequestDTO struct {
	ID       int32     `json: id`
	Name     string    `json:name,omitempty`
	Email    string    `json:email,omitempty`
	Password string    `json:password,omitempty`
	CreateAt time.Time `json:createdAt,omitempty`
}

// ValidateUpdateUserRequestDTO valid the rules on propertys
func (u UpdateUserRequestDTO) ValidateUpdateUserRequest() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.ID, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&u.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 20)),
		validation.Field(&u.Email, validation.Required, is.Email))
}

// FromJSONCreateUserRequest converts json body request to a CreateUserRequestDTO struct
func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequestDTO, error) {
	createUserRequest := CreateUserRequestDTO{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}

// FromJSONUpdateUserRequestDTO converts json body request to a UpdateUserRequestDTODTO struct
func FromJSONUpdateUserRequestDTO(body io.Reader) (*UpdateUserRequestDTO, error) {
	UpdateUserRequestDTO := UpdateUserRequestDTO{}
	if err := json.NewDecoder(body).Decode(&UpdateUserRequestDTO); err != nil {
		return nil, err
	}

	return &UpdateUserRequestDTO, nil
}
