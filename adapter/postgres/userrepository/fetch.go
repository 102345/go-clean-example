package userrepository

import (
	"context"
	"log"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination, error) {
	ctx := context.Background()
	users := []domain.User{}
	total := int32(0)

	query, queryCount, _ := paginate.Paginate("SELECT * FROM user_api").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "email").
		Query()

	log.Printf("Query formada: %s", *query)

	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			log.Printf("Error Query in repository User: %s", err)
			return nil, err
		}

		for rows.Next() {
			user := domain.User{}

			rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Password,
				&user.CreatedAt,
			)

			users = append(users, user)
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			log.Printf("Error Query Row in repository User: %s", err)
			return nil, err
		}
	}

	return &domain.Pagination{
		Items: users,
		Total: total,
	}, nil
}
