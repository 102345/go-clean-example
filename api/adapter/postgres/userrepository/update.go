package userrepository

import (
	"context"

	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (repository repository) Update(userRequest *dto.UpdateUserRequestDTO) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}

	_, err := repository.db.Exec(
		ctx,
		"Update user_api set name = $2, email = $3, password= $4 where id = $1",
		userRequest.ID,
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
	)

	if err != nil {
		return nil, err
	}

	user.ID = userRequest.ID
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	return &user, nil
}
