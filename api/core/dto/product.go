package dto

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation"
)

// CreateProductRequestDTO is an representation request body to create a new Product
type CreateProductRequestDTO struct {
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

// ValidateCreateRequest valid the rules on propertys
func (p CreateProductRequestDTO) ValidateCreateRequest() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&p.Price, validation.Required, validation.Min(1.0)),
		validation.Field(&p.Description, validation.Required, validation.Length(5, 500)))

}

// UpdateProductRequestDTO is an representation request body to update a Product
type UpdateProductRequestDTO struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}

// UpdateProductRequestDTO valid the rules on propertys
func (p UpdateProductRequestDTO) ValidateUpdateRequest() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.ID, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&p.Price, validation.Required, validation.Min(1.0)),
		validation.Field(&p.Description, validation.Required, validation.Length(5, 500)))

}

// FromJSONCreateProductRequestDTO converts json body request to a CreateProductRequestDTO struct
func FromJSONCreateProductRequestDTO(body io.Reader) (*CreateProductRequestDTO, error) {
	CreateProductRequestDTO := CreateProductRequestDTO{}
	if err := json.NewDecoder(body).Decode(&CreateProductRequestDTO); err != nil {
		return nil, err
	}

	return &CreateProductRequestDTO, nil
}

// FromJSONUpdateProductRequestDTO converts json body request to a UpdateProductRequestDTO struct
func FromJSONUpdateProductRequestDTO(body io.Reader) (*UpdateProductRequestDTO, error) {
	UpdateProductRequestDTO := UpdateProductRequestDTO{}
	if err := json.NewDecoder(body).Decode(&UpdateProductRequestDTO); err != nil {
		return nil, err
	}

	return &UpdateProductRequestDTO, nil
}
