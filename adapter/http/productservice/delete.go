package productservice

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

func (service service) Delete(response http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	id, erro := strconv.ParseUint(parametros["product_id"], 10, 64)
	if erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	err := service.usecase.Delete(id)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	infrastructure.JSON(response, http.StatusNoContent, nil)
}
