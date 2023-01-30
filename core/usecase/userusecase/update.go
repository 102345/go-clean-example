package userusecase

import (
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

// Update implements domain.UserUseCase
func (usecase usecase) Update(userRequest *dto.UpdateUserRequest) (*domain.User, error) {

	user, err := usecase.repository.Update(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}
