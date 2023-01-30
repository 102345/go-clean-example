package userusecase

import (
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	users, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return users, nil
}
