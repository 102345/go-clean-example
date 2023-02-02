package userservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/marc/go-clean-example/core/domain"
	infrastructure "github.com/marc/go-clean-example/infra-structure"
	"github.com/marc/go-clean-example/infra-structure/middlewares/authentication"
	"github.com/marc/go-clean-example/infra-structure/middlewares/security"
)

func (service service) Login(response http.ResponseWriter, request *http.Request) {

	bodyRequest, err := ioutil.ReadAll(request.Body)
	if err != nil {
		infrastructure.Erro(response, http.StatusUnprocessableEntity, err)
		return
	}

	var user domain.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		infrastructure.Erro(response, http.StatusBadRequest, err)
		return
	}

	userDatabase, err := service.usecase.SearchByEmail(user.Email)

	if err != nil {
		infrastructure.Erro(response, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userDatabase.Password, user.Password); err != nil {
		infrastructure.Erro(response, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.CreateToken(uint64(userDatabase.ID))

	userID := strconv.FormatUint(uint64(userDatabase.ID), 10)

	infrastructure.JSON(response, http.StatusOK, domain.AuthenticationData{ID: userID, Token: token})

}
