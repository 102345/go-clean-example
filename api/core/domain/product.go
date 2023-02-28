package domain

import (
	"net/http"

	"github.com/marc/go-clean-example/core/dto"
)

// Product is entity of table product database column
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
	FindById(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequestDTO) (*Product, error)
	Update(productRequest *dto.UpdateProductRequestDTO) (*Product, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	FindById(id int64) (Product, error)
}

// ProductRepository is a contract of database connection adapter layer
type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequestDTO) (*Product, error)
	Update(productRequest *dto.UpdateProductRequestDTO) (*Product, error)
	Delete(id uint64) error
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination, error)
	FindById(id int64) (Product, error)
}
