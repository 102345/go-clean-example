package productrepository_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/adapter/postgres/productrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupUpdate() ([]string, dto.UpdateProductRequest, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"name", "price", "description", "id"}
	fakeProductRequest := dto.UpdateProductRequest{}
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeProductRequest, fakeProductDBResponse, mock
}

func TestUpdate(t *testing.T) {
	cols, fakeProductRequest, fakeProductDBResponse, mock := setupUpdate()
	defer mock.Close()

	//mock.ExpectExec("UPDATE baskets").WithArgs(newProp, id).WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery("UPDATE product (.+) WHERE (.+)").WithArgs(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
		fakeProductRequest.ID,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
		fakeProductRequest.ID,
	))

	sut := productrepository.New(mock)

	product, err := sut.Update(&fakeProductRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, product.ID, fakeProductDBResponse.ID)
	require.Equal(t, product.Name, fakeProductDBResponse.Name)
	require.Equal(t, product.Price, fakeProductDBResponse.Price)
	require.Equal(t, product.Description, fakeProductDBResponse.Description)
}

// func TestUpdate_DBError(t *testing.T) {
// 	_, fakeProductRequest, _, mock := setupCreate()
// 	defer mock.Close()

// 	mock.ExpectQuery("UPDATE product SET (.+) WHERE (.+)").WithArgs(
// 		fakeProductRequest.Name,
// 		fakeProductRequest.Price,
// 		fakeProductRequest.Description,
// 	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

// 	sut := productrepository.New(mock)
// 	product, err := sut.Create(&fakeProductRequest)

// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	require.NotNil(t, err)
// 	require.Nil(t, product)
// }
