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
)

func setupUpdate(t *testing.T) (dto.UpdateUserRequestDTO, domain.User, *gomock.Controller) {
	fakeUserRequest := dto.UpdateUserRequestDTO{}
	fakeUser := domain.User{}
	faker.FakeData(&fakeUserRequest)
	faker.FakeData(&fakeUser)

	mockCtrl := gomock.NewController(t)

	return fakeUserRequest, fakeUser, mockCtrl
}

func TestUpdate(t *testing.T) {
	fakeUserRequest, fakeUser, mock := setupUpdate(t)
	fakeUserRequest.Email = "marc@teste.com"
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Update(&fakeUserRequest).Return(&fakeUser, nil)

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code is not correct")
	}
}

func TestUpdate_JsonErrorFormater(t *testing.T) {
	_, _, mock := setupUpdate(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestUpdate_ErrorValidate(t *testing.T) {

	fakeUserRequest, _, mock := setupUpdate(t)

	fakeUserRequest.Name = ""
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestUpdate_UserError(t *testing.T) {
	fakeUserRequest, _, mock := setupUpdate(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Update(&fakeUserRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := userservice.New(mockUserUseCase)

	payload, _ := json.Marshal(fakeUserRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/user/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
