package productrepository

import (
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
