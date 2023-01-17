package dto

import (
	"encoding/json"
	"io"
)

// CreateProductRequest is an representation request body to create a new Product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

// UpdateProductRequest is an representation request body to update a Product
type UpdateProductRequest struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

// FromJSONCreateProductRequest converts json body request to a CreateProductRequest struct
func FromJSONCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}
	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	return &createProductRequest, nil
}

// FromJSONUpdateProductRequest converts json body request to a UpdateProductRequest struct
func FromJSONUpdateProductRequest(body io.Reader) (*UpdateProductRequest, error) {
	updateProductRequest := UpdateProductRequest{}
	if err := json.NewDecoder(body).Decode(&updateProductRequest); err != nil {
		return nil, err
	}

	return &updateProductRequest, nil
}
