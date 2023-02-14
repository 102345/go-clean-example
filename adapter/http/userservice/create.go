package userservice

import (
	"net/http"
	"strings"

	"github.com/marc/go-clean-example/core/dto"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
	"github.com/marc/go-clean-example/infra-structure/middlewares/security"
)

// @Summary Create new User
// @Description Create new User
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserRequest true "user"
// @Success 201 {object} domain.User
// @Router /user [post]
func (service service) Create(response http.ResponseWriter, request *http.Request) {

	userRequest, err := dto.FromJSONCreateUserRequest(request.Body)

	if err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	if erro := userRequest.ValidateCreateUserRequest(); erro != nil {
		infrastructure.Erro(response, http.StatusBadRequest, erro)
		return
	}

	erro := service.formatUser(userRequest)
	if erro != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	user, err := service.usecase.Create(userRequest)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	infrastructure.JSON(response, http.StatusCreated, user)
}

func (service) formatUser(userRequest *dto.CreateUserRequest) error {

	passwordHash, erro := security.Hash(userRequest.Password)
	if erro != nil {
		return erro
	}
	userRequest.Name = strings.TrimSpace(userRequest.Name)
	userRequest.Email = strings.TrimSpace(userRequest.Email)
	userRequest.Password = string(passwordHash)
	return nil
}
