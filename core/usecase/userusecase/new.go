package userusecase

import (
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

type usecase struct {
	repository domain.UserRepository
}

// Delete implements domain.UserUseCase
func (*usecase) Delete(id uint64) error {
	panic("unimplemented")
}

// Fetch implements domain.UserUseCase
func (*usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	panic("unimplemented")
}

// Update implements domain.UserUseCase
func (*usecase) Update(userRequest *dto.UpdateUserRequest) (*domain.User, error) {
	panic("unimplemented")
}

// New returns contract implementation of UserUseCase
func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase{
		repository: repository,
	}
}
