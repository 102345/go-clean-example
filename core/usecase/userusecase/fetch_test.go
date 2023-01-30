package userusecase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/marc/go-clean-example/core/usecase/userusecase"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeDBUser := domain.User{}

	faker.FakeData(&fakeDBUser)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(&domain.Pagination{
		Items: []domain.User{fakeDBUser},
		Total: 1,
	}, nil)

	sut := userusecase.New(mockUserRepository)
	Users, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	for _, User := range Users.Items.([]domain.User) {
		require.Nil(t, err)
		require.NotEmpty(t, User.ID)
		require.Equal(t, User.Name, fakeDBUser.Name)
		require.Equal(t, User.Email, fakeDBUser.Email)
		require.Equal(t, User.Password, fakeDBUser.Password)
	}
}

func TestFetch_Error(t *testing.T) {
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Fetch(&fakePaginationRequestParams).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := userusecase.New(mockUserRepository)
	User, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)
	require.Nil(t, User)
}
