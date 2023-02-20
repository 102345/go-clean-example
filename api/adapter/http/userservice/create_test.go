package userservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/adapter/http/userservice"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/marc/go-clean-example/infra-structure/middlewares/security"
)

func setupCreate(t *testing.T) (dto.CreateUserRequest, domain.User, *gomock.Controller) {
	fakeUserRequest := dto.CreateUserRequest{}

	fakeUser := domain.User{}

	faker.FakeData(&fakeUserRequest)
	faker.FakeData(&fakeUser)

	mockCtrl := gomock.NewController(t)

	return fakeUserRequest, fakeUser, mockCtrl
}

func TestCreate(t *testing.T) {

	fakeUserRequest, fakeUser, mock := setupCreate(t)

	passwordHash, _ := security.Hash("1234567909090tueyueyeye")

	fakeUserRequest.Email = "teste@teste.com"
	fakeUser.Password = string(passwordHash)
	fakeUser.CreatedAt = time.Now()
	fakeUserRequest.Password = fakeUser.Password
	fakeUserRequest.CreateAt = fakeUser.CreatedAt

	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Create(&fakeUserRequest).Return(&fakeUser, nil)

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupCreate(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_ErrorValidate(t *testing.T) {

	fakeUserRequest, _, mock := setupCreate(t)

	fakeUserRequest.Name = ""
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_UserError(t *testing.T) {
	fakeUserRequest, _, mock := setupCreate(t)
	fakeUserRequest.Email = "teste@teste.com"
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Create(&fakeUserRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
