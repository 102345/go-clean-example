package userusecase

import "github.com/marc/go-clean-example/core/domain"

func (usecase usecase) SearchByEmail(email string) (domain.User, error) {

	user, err := usecase.repository.SearchByEmail(email)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
