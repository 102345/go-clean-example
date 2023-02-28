package userusecase

import (
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequestDTO) (*domain.User, error) {
	user, err := usecase.repository.Create(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}
