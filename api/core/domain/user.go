package domain

import (
	"net/http"
	"time"

	"github.com/marc/go-clean-example/core/dto"
)

// User is entity of table product database column
type User struct {
	ID        int32     `json: id,omitempty`
	Name      string    `json:name,omitempty`
	Email     string    `json:email,omitempty`
	Password  string    `json:password,omitempty`
	CreatedAt time.Time `json:createdAt,omitempty`
}

// AuthenticationData is the entity that stores the authentication data
type AuthenticationData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

// UserService is a contract of http adapter layer
type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
	Login(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequestDTO) (*User, error)
	Update(userRequest *dto.UpdateUserRequestDTO) (*User, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*PaginationUsers, error)
	SearchByEmail(email string) (User, error)
}

// UserRepository is a contract of database connection adapter layer
type UserRepository interface {
	Create(userRequest *dto.CreateUserRequestDTO) (*User, error)
	Update(userRequest *dto.UpdateUserRequestDTO) (*User, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*PaginationUsers, error)
	SearchByEmail(email string) (User, error)
}
