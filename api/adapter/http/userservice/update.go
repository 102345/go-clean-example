package userservice

import (
	"net/http"
	"strings"

	"github.com/marc/go-clean-example/core/dto"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
	"github.com/marc/go-clean-example/infra-structure/middlewares/security"
)

// @Summary Update a User
// @Description Update a User
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.UpdateUserRequestDTO true "user"
// @Success 200 {object} domain.User
// @Router /user [put]
func (service service) Update(response http.ResponseWriter, request *http.Request) {

	userRequest, err := dto.FromJSONUpdateUserRequestDTO(request.Body)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	if erro := userRequest.ValidateUpdateUserRequest(); erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	erro := service.formatUserUpdate(userRequest)
	if erro != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	user, err := service.usecase.Update(userRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusOK, user)
}

func (service) formatUserUpdate(userRequest *dto.UpdateUserRequestDTO) error {

	passwordHash, erro := security.Hash(userRequest.Password)
	if erro != nil {
		return erro
	}
	userRequest.Name = strings.TrimSpace(userRequest.Name)
	userRequest.Email = strings.TrimSpace(userRequest.Email)
	userRequest.Password = string(passwordHash)
	return nil
}
