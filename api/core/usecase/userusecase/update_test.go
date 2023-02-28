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

func TestUpdate(t *testing.T) {
	fakeRequestUser := dto.UpdateUserRequestDTO{}
	fakeDBUser := domain.User{}
	faker.FakeData(&fakeRequestUser)
	faker.FakeData(&fakeDBUser)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Update(&fakeRequestUser).Return(&fakeDBUser, nil)

	sut := userusecase.New(mockUserRepository)
	User, err := sut.Update(&fakeRequestUser)

	require.Nil(t, err)
	require.NotEmpty(t, User.ID)
	require.Equal(t, User.Name, fakeDBUser.Name)
	require.Equal(t, User.Email, fakeDBUser.Email)
	require.Equal(t, User.Password, fakeDBUser.Password)
}

func TestUpdate_Error(t *testing.T) {
	fakeRequestUser := dto.UpdateUserRequestDTO{}
	faker.FakeData(&fakeRequestUser)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Update(&fakeRequestUser).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := userusecase.New(mockUserRepository)
	User, err := sut.Update(&fakeRequestUser)

	require.NotNil(t, err)
	require.Nil(t, User)
}
