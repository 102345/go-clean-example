package userusecase

import (
	"github.com/marc/go-clean-example/core/domain"
)

type usecase struct {
	repository domain.UserRepository
}

// New returns contract implementation of UserUseCase
func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase{
		repository: repository,
	}
}
