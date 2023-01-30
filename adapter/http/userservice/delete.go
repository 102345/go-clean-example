package userservice

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user_id query integer true "1"
// @Success 204 {object} domain.User
// @Router /user/{user_id} [delete]
func (service service) Delete(response http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)

	id, erro := strconv.ParseUint(parametros["user_id"], 10, 64)
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
