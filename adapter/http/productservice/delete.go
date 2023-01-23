package productservice

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// @Summary Delete a product
// @Description Delete a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product_id query integer true "1"
// @Success 204 {object} domain.Product
// @Router /product/{product_id} [delete]
func (service service) Delete(response http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	id, erro := strconv.ParseUint(parametros["product_id"], 10, 64)
	if erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	err := service.usecase.Delete(id)

	if err != nil {
		infrastructure.JSON(response, http.StatusInternalServerError, nil)
		return
	}

	infrastructure.JSON(response, http.StatusNoContent, nil)
}
