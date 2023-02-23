package productservice

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

func (service service) FindById(response http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	productID, err := strconv.ParseInt(parametros["product_id"], 10, 64)
	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	product, err := service.usecase.FindById(productID)
	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusOK, product)

}
