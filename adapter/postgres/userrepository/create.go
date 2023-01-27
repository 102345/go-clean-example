package userrepository

import (
	"context"

	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (repository repository) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}

	// log.Printf("password: %s", userRequest.Password)
	// log.Printf("passwordCrypto: %b", passwordCrypto)

	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO user_api (name, email, password) VALUES ($1, $2, $3) returning *",
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
	).Scan(
		&user.IDUser,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
