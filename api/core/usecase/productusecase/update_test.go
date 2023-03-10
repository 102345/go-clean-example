package productusecase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/domain/mocks"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/marc/go-clean-example/core/usecase/productusecase"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	fakeRequestProduct := dto.UpdateProductRequestDTO{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Update(&fakeRequestProduct).Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Update(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)
}

func TestUpdate_Error(t *testing.T) {
	fakeRequestProduct := dto.UpdateProductRequestDTO{}
	faker.FakeData(&fakeRequestProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockProductRepository := mocks.NewMockProductRepository(mockCtrl)
	mockProductRepository.EXPECT().Update(&fakeRequestProduct).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := productusecase.New(mockProductRepository)
	product, err := sut.Update(&fakeRequestProduct)

	require.NotNil(t, err)
	require.Nil(t, product)
}
