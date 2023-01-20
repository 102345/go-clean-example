package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/usecase/productusecase"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {

	product_id := 1

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Delete(uint64(product_id)).Return(nil)

	sut := productusecase.New(mockProductRepository)
	err := sut.Delete(uint64(product_id))

	require.Nil(t, err)
}

func TestDelete_Error(t *testing.T) {

	product_id := 1

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Delete(uint64(product_id)).Return(fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	err := sut.Delete(uint64(product_id))

	require.NotNil(t, err)

}
