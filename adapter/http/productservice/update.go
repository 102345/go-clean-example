package productservice

import (
	"net/http"

	"github.com/marc/go-clean-example/core/dto"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// @Summary Update a product
// @Description Update a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body dto.UpdateProductRequest true "product"
// @Success 200 {object} domain.Product
// @Router /product/{product_id} [put]
func (service service) Update(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONUpdateProductRequest(request.Body)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	if erro := productRequest.ValidateUpdateRequest(); erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	product, err := service.usecase.Update(productRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusOK, product)
}
