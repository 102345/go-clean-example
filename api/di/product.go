package di

import (
	"github.com/marc/go-clean-example/adapter/http/productservice"
	"github.com/marc/go-clean-example/adapter/postgres"
	"github.com/marc/go-clean-example/adapter/postgres/productrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}
