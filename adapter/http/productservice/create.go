package productservice

import (
	"net/http"

	"github.com/marc/go-clean-example/core/dto"
	productvalidator "github.com/marc/go-clean-example/core/validator/productValidator"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// @Summary Create new product
// @Description Create new product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductRequest true "product"
// @Success 201 {object} domain.Product
// @Router /product [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {

	productRequest, err := dto.FromJSONCreateProductRequest(request.Body)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	if erro := productvalidator.ValidateInsert(productRequest); erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	product, err := service.usecase.Create(productRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusCreated, product)
}
