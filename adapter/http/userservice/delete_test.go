package userservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/marc/go-clean-example/adapter/http/userservice"
	"github.com/marc/go-clean-example/core/domain/mocks"
)

func setupDelete(t *testing.T) *gomock.Controller {

	mockCtrl := gomock.NewController(t)

	return mockCtrl
}

func initDeleteRouter(mockUserUseCase *mocks.MockUserUseCase) http.Handler {

	sut := userservice.New(mockUserUseCase)
	r := mux.NewRouter()

	r.Handle("/user/{user_id}", http.HandlerFunc(sut.Delete)).Methods(http.MethodDelete)
	return r

}

func TestDelete(t *testing.T) {
	mock := setupDelete(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Delete(1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	r.Header.Set("Content-Type", "application/json")

	sut := initDeleteRouter(mockUserUseCase)
	sut.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Errorf("status code is not correct")
	}
}

func TestDelete_ErrorValidate(t *testing.T) {

	mock := setupDelete(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/user/", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestDelete_UserError(t *testing.T) {

	mock := setupDelete(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Delete(1).Return(fmt.Errorf("ANY ERROR"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/user/1", nil)
	r.Header.Set("Content-Type", "application/json")

	sut := initDeleteRouter(mockUserUseCase)
	sut.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
