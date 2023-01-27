package userrepository

import (
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
)

type repository struct {
	db postgres.PoolInterface
}

// Delete implements domain.UserRepository
func (*repository) Delete(id uint64) error {
	panic("unimplemented")
}

// Fetch implements domain.UserRepository
func (*repository) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	panic("unimplemented")
}

// Update implements domain.UserRepository
func (*repository) Update(userRequest *dto.UpdateUserRequest) (*domain.User, error) {
	panic("unimplemented")
}

// New returns contract implementation of UserRepository
func New(db postgres.PoolInterface) domain.UserRepository {
	return &repository{
		db: db,
	}
}
