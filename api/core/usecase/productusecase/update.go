package productusecase

import (
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (usecase usecase) Update(productRequest *dto.UpdateProductRequestDTO) (*domain.Product, error) {
	product, err := usecase.repository.Update(productRequest)

	if err != nil {
		return nil, err
	}

	return product, nil
}
