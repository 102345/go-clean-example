package userrepository

import (
	"context"

	"github.com/marc/go-clean-example/core/domain"
)

func (repository repository) SearchByEmail(email string) (domain.User, error) {

	ctx := context.Background()

	row, erro := repository.db.Query(
		ctx,
		"select id, password from user_api where email = ?",
		email,
	)

	if erro != nil {
		return domain.User{}, erro
	}

	defer row.Close()

	var user domain.User

	if row.Next() {
		if erro = row.Scan(
			&user.ID,
			&user.Password,
		); erro != nil {
			return domain.User{}, erro
		}
	}

	return user, nil
}
