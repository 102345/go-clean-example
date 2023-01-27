package domain

import (
	"net/http"
	"time"

	"github.com/marc/go-clean-example/core/dto"
)

// User is entity of table product database column
type User struct {
	IDUser   int32     `json: idUser,omitempty`
	Name     string    `json:name,omitempty`
	Email    string    `json:email,omitempty`
	Password string    `json:password,omitempty`
	CreateAt time.Time `json:createdAt,omitempty`
}

// UserService is a contract of http adapter layer
type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Update(userRequest *dto.UpdateUserRequest) (*User, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
}

// UserRepository is a contract of database connection adapter layer
type UserRepository interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Update(userRequest *dto.UpdateUserRequest) (*User, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
}
