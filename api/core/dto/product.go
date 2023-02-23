package dto

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateProductRequest is an representation request body to create a new Product
type CreateProductRequest struct {
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

// ValidateCreateRequest valid the rules on propertys
func (p CreateProductRequest) ValidateCreateRequest() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&p.Price, validation.Required, validation.Min(1.0)),
		validation.Field(&p.Description, validation.Required, validation.Length(5, 500)))

}

// UpdateProductRequest is an representation request body to update a Product
type UpdateProductRequest struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

// UpdateProductRequest valid the rules on propertys
func (p UpdateProductRequest) ValidateUpdateRequest() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ID, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&p.Price, validation.Required, validation.Min(1.0)),
		validation.Field(&p.Description, validation.Required, validation.Length(5, 500)))

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
