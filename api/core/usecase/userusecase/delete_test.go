package userusecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/usecase/userusecase"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {

	User_id := 1

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Delete(uint64(User_id)).Return(nil)

	sut := userusecase.New(mockUserRepository)
	err := sut.Delete(uint64(User_id))

	require.Nil(t, err)
}

func TestDelete_Error(t *testing.T) {

	User_id := 1

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserRepository := mocks.NewMockUserRepository(mockCtrl)
	mockUserRepository.EXPECT().Delete(uint64(User_id)).Return(fmt.Errorf("ANY ERROR"))

	sut := userusecase.New(mockUserRepository)
	err := sut.Delete(uint64(User_id))

	require.NotNil(t, err)

}
