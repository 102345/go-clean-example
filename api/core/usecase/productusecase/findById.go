package productusecase

import "github.com/marc/go-clean-example/core/domain"

func (usecase usecase) FindById(id uint64) (domain.Product, error) {

	product, err := usecase.repository.FindById(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil

}
