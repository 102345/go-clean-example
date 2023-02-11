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

func setupCreate(t *testing.T) (dto.CreateProductRequest, domain.Product, *gomock.Controller) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	mockCtrl := gomock.NewController(t)

	return fakeProductRequest, fakeProduct, mockCtrl
}

func TestCreate(t *testing.T) {
	fakeProductRequest, fakeProduct, mock := setupCreate(t)
	fakeProductRequest.Price = "100"
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(&fakeProduct, nil)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
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
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_ErrorValidate(t *testing.T) {

	fakeProductRequest, _, mock := setupCreate(t)

	fakeProductRequest.Name = ""
	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("status code is not correct")
	}
}

func TestCreate_ProductError(t *testing.T) {
	fakeProductRequest, _, mock := setupCreate(t)
	fakeProductRequest.Price = "100"

	defer mock.Finish()
	mockProductUseCase := mocks.NewMockProductUseCase(mock)
	mockProductUseCase.EXPECT().Create(&fakeProductRequest).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productservice.New(mockProductUseCase)

	payload, _ := json.Marshal(fakeProductRequest)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	r.Header.Set("Content-Type", "application/json")
	sut.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
