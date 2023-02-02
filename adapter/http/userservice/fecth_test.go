package userservice_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/adapter/http/userservice"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
)

func setupFetch(t *testing.T) (dto.PaginationRequestParms, domain.User, *gomock.Controller) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         []string{""},
		Descending:   []string{""},
		Search:       "",
	}
	fakeUser := domain.User{}
	faker.FakeData(&fakeUser)

	mockCtrl := gomock.NewController(t)

	return fakePaginationRequestParams, fakeUser, mockCtrl
}

func TestFetch(t *testing.T) {
	fakePaginationRequestParams, fakeUser, mock := setupFetch(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.User{fakeUser},
		Total: 1,
	}, nil)

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/user", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("status code is not correct")
	}
}

func TestFetch_ProductError(t *testing.T) {
	fakePaginationRequestParams, _, mock := setupFetch(t)
	defer mock.Finish()
	mockUserUseCase := mocks.NewMockUserUseCase(mock)
	mockUserUseCase.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := userservice.New(mockUserUseCase)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/user", nil)
	r.Header.Set("Content-Type", "application/json")
	queryStringParams := r.URL.Query()
	queryStringParams.Add("page", "1")
	queryStringParams.Add("itemsPerPage", "10")
	queryStringParams.Add("sort", "")
	queryStringParams.Add("descending", "")
	queryStringParams.Add("search", "")
	r.URL.RawQuery = queryStringParams.Encode()
	sut.Fetch(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("status code is not correct")
	}
}
