package productservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/adapter/http/productservice"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
)

func setupUpdate(t *testing.T) (dto.UpdateProductRequestDTO, domain.Product, *gomock.Controller) {
	fakeProductRequest := dto.UpdateProductRequestDTO{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakeProductRequest, fakeProduct, mockCtrl
}

func TestUpdate(t *testing.T) {
	fakeProductRequest, fakeProduct, mock := setupUpdate(t)
	//fakeProductRequest.Price = "100"
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Update(&fakeProductRequest).Return(&fakeProduct, nil)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product/1", strings.NewReader(string(payload)))
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
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product/1", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestUpdate_ErrorValidate(t *testing.T) {

	fakeProductRequest, _, mock := setupUpdate(t)

	fakeProductRequest.Name = ""
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestUpdate_ProductError(t *testing.T) {
	fakeProductRequest, _, mock := setupUpdate(t)
	//fakeProductRequest.Price = "100"
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Update(&fakeProductRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/product/1", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Update(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
