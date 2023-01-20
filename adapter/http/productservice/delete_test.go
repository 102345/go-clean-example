package productservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/adapter/http/productservice"
	"github.com/marc/go-clean-example/core/domain/mocks"
)

func setupDelete(t *testing.T) *gomock.Controller {

	mockCtrl := gomock.NewController(t)

	return mockCtrl
}

func TestDelete(t *testing.T) {
	mock := setupDelete(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Delete(1).Return(nil)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/product/1", nil)
	//r = SetURLVars(r, map[string]string{"id":"1"})
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code is not correct")
	}
}

func TestDelete_ErrorValidate(t *testing.T) {

	mock := setupDelete(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/product/", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestDelete_ProductError(t *testing.T) {

	mock := setupDelete(t)
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Delete(1).Return(fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodDelete, "/product/1", nil)
	r.Header.Set("Content-Type", "application/json")
	sut.Delete(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
