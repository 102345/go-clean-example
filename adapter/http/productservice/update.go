package productservice

import (
	"net/http"

	"github.com/marc/go-clean-example/core/dto"
	productvalidator "github.com/marc/go-clean-example/core/validator/productValidator"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

func (service service) Update(response http.ResponseWriter, request *http.Request) {
	productRequest, err := dto.FromJSONUpdateProductRequest(request.Body)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	if erro := productvalidator.ValidateUpdate(productRequest); erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	product, err := service.usecase.Update(productRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusCreated, product)
}
