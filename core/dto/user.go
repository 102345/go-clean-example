package dto

import (
	"encoding/json"
	"io"
	"time"
)

// CreateProductRequest is an representation request body to create a new User
type CreateUserRequest struct {
	Name     string    `json:name`
	Email    string    `json:email`
	Password string    `json:password`
	CreateAt time.Time `json:createdAt`
}

// UpdateUserRequest is an representation request body to update a User
type UpdateUserRequest struct {
	ID       int32     `json: idUser`
	Name     string    `json:name,omitempty`
	Email    string    `json:email,omitempty`
	Password string    `json:password,omitempty`
	CreateAt time.Time `json:createdAt,omitempty`
}

// FromJSONCreateUserRequest converts json body request to a CreateUserRequest struct
func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}

// FromJSONUpdateUserRequest converts json body request to a UpdateUserRequest struct
func FromJSONUpdateUserRequest(body io.Reader) (*UpdateUserRequest, error) {
	updateUserRequest := UpdateUserRequest{}
	if err := json.NewDecoder(body).Decode(&updateUserRequest); err != nil {
		return nil, err
	}

	return &updateUserRequest, nil
}
