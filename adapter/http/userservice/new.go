package userservice

import (
	"net/http"

	"github.com/marc/go-clean-example/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

// Delete implements domain.UserService
func (*service) Delete(response http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// Fetch implements domain.UserService
func (*service) Fetch(response http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// Update implements domain.UserService
func (*service) Update(response http.ResponseWriter, request *http.Request) {
	panic("unimplemented")
}

// New returns contract implementation of ProductService
func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
