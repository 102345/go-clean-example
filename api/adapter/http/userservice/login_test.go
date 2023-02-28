package userservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/adapter/http/userservice"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/marc/go-clean-example/infra-structure/middlewares/security"
)

func setupLogin(t *testing.T) (dto.CreateUserRequestDTO, domain.User, *gomock.Controller) {
	fakeUserRequest := dto.CreateUserRequestDTO{}
	faker.FakeData(&fakeUserRequest)

	fakeUser := domain.User{}
	faker.FakeData(&fakeUser)

	mockCtrl := gomock.NewController(t)

	return fakeUserRequest, fakeUser, mockCtrl
}

// func TestLogin_NotBodyRequestUserError(t *testing.T) {
// 	_, _, mock := setupLogin(t)

// 	defer mock.Finish()
// 	mockUserUseCase := mocks.NewMockUserUseCase(mock)

// 	sut := userservice.New(mockUserUseCase)

// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodPost, "/login", nil)
// 	r.Header.Set("Content-Type", "application/json")
// 	sut.Login(w, r)

// 	res := w.Result()
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusUnprocessableEntity {
// 		t.Errorf("status code is not correct")
// 	}
// }

func TestLogin(t *testing.T) {

	fakeUserRequest, fakeUser, mock := setupLogin(t)

	passwordHash, _ := security.Hash("1234567909090tueyueyeye")

	fakeUserRequest.Email = "teste@teste.com"
	fakeUser.Password = string(passwordHash)
	fakeUserRequest.Password = "1234567909090tueyueyeye"
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().SearchByEmail("teste@teste.com").Return(fakeUser, nil)

	payload, _ := json.Marshal(fakeUserRequest)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Login(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code is not correct")
	}
}

func TestLogin_BodyRequestUserError(t *testing.T) {
	_, _, mock := setupLogin(t)

	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Login(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestLogin_UserError(t *testing.T) {
	fakeUserRequest, _, mock := setupLogin(t)
	fakeUserRequest.Email = "teste@teste.com"
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().SearchByEmail("teste@teste.com").Return(domain.User{}, fmt.Errorf("ANY ERROR"))

	payload, _ := json.Marshal(fakeUserRequest)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Login(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}

func TestLogin_ReturnVerifyPasswordError(t *testing.T) {

	fakeUserRequest, fakeUser, mock := setupLogin(t)

	passwordHash, _ := security.Hash("1234567909090tueyueyeye")

	fakeUserRequest.Email = "teste@teste.com"
	fakeUser.Password = string(passwordHash)
	fakeUserRequest.Password = "1234567909090"
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().SearchByEmail("teste@teste.com").Return(fakeUser, nil)

	payload, _ := json.Marshal(fakeUserRequest)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Login(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("status code is not correct")
	}
}
