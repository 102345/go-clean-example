package userservice

import (
	"github.com/marc/go-clean-example/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

// New returns contract implementation of UserService
func New(usecase domain.UserUseCase) domain.UserService {
	return &service{
		usecase: usecase,
	}
}
