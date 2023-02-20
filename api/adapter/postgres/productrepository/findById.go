package productrepository

import (
	"context"
	"log"

	"github.com/marc/go-clean-example/core/domain"
)

func (repository repository) FindById(id uint64) (domain.Product, error) {

	ctx := context.Background()

	row, erro := repository.db.Query(
		ctx,
		"select id, name, price, description from product where id =$1",
		id,
	)

	if erro != nil {
		log.Printf("Error database : %s", erro)
		return domain.Product{}, erro
	}

	defer row.Close()

	var product domain.Product

	if row.Next() {
		if erro = row.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
		); erro != nil {
			return domain.Product{}, erro
		}
	}

	return product, nil
}
