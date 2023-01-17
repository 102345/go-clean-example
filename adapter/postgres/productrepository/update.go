package productrepository

import (
	"context"

	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (repository repository) Update(productRequest *dto.UpdateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	product := domain.Product{}

	err := repository.db.QueryRow(
		ctx,
		"Update product set name = $2, price = $3, description= $4 where id = $1 returning *",
		productRequest.ID,
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
