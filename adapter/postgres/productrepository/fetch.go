package productrepository

import (
	"context"
	"log"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	products := []domain.Product{}
	pageData := domain.Page{}
	pageData.PageNumber = int32(pagination.Page)
	pageData.Quantity = int32(pagination.ItemsPerPage)

	query, queryCount, _ := paginate.Paginate("SELECT * FROM product").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Query()

	log.Printf("Query formada: %s", *query)

	log.Printf("QueryCount formada: %s", *queryCount)

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			product := domain.Product{}

			rows.Scan(
				&product.ID,
				&product.Name,
				&product.Price,
				&product.Description,
			)

			products = append(products, product)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&pageData.Total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination{
		Items: products,
		Page:  pageData,
	}, nil
}
