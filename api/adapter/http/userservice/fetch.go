package userservice

import (
	"net/http"

	"github.com/marc/go-clean-example/core/dto"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
)

// @Summary Fetch users with server pagination
// @Description Fetch users with server pagination
// @Tags user
// @Accept  json
// @Produce  json
// @Param sort query string true "1,2"
// @Param descending query string true "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "value_parameter"
// @Success 200 {object} domain.Pagination
// @Router /user [get]
func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParams(request)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	users, err := service.usecase.Fetch(paginationRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusOK, users)
}
