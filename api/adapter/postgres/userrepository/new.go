package userrepository

import (
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of UserRepository
func New(db postgres.PoolInterface) domain.UserRepository {
	return &repository{
		db: db,
	}
}
