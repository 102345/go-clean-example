package productrepository_test

import (
	"fmt"
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
	_, fakeProductRequest, _, mock := setupUpdate()
	defer mock.Close()

	//mock.ExpectExec("UPDATE baskets").WithArgs(newProp, id).WillReturnResult(mock.NewResult(0, 1))
	//Update product set name = $2, price = $3, description= $4 where id = $1 returning *

	//sql := "Update product set name = $2, price = $3, description= $4 where id = $1 returning *"

	// rows := pgxmock.NewRows(cols).AddRow(
	// 	fakeProductDBResponse.ID,
	// 	fakeProductDBResponse.Name,
	// 	fakeProductDBResponse.Price,
	// 	fakeProductDBResponse.Description,
	// )

	//mock.ExpectExec("UPDATE products").WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	//UPDATE product (.+) WHERE (.+)

	mock.ExpectExec("Update product set name = $2, price = $3, description= $4 where id = $1").WithArgs(
		fakeProductRequest.ID,
		fakeProductRequest.Name,
		fakeProductRequest.Price,
		fakeProductRequest.Description,
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	// mock.ExpectQuery("Update product set name = $2, price = $3, description= $4 where id = $1 returning *").WithArgs(
	// 	fakeProductRequest.Name,
	// 	fakeProductRequest.Price,
	// 	fakeProductRequest.Description,
	// 	fakeProductRequest.ID,
	// ).WillReturnRows(pgxmock.NewRows(cols).AddRow(
	// 	fakeProductRequest.Name,
	// 	fakeProductRequest.Price,
	// 	fakeProductRequest.Description,
	// 	fakeProductRequest.ID,
	// ))

	sut := productrepository.New(mock)

	product, err := sut.Update(&fakeProductRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, product.ID, fakeProductRequest.ID)
	require.Equal(t, product.Name, fakeProductRequest.Name)
	require.Equal(t, product.Price, fakeProductRequest.Price)
	require.Equal(t, product.Description, fakeProductRequest.Description)
}

func TestUpdate_DBError(t *testing.T) {
	_, _, _, mock := setupCreate()

	fakeProductRequestUpdate := dto.UpdateProductRequest{}
	defer mock.Close()

	// mock.ExpectQuery("UPDATE product SET (.+) WHERE (.+)").WithArgs(
	// 	fakeProductRequestUpdate.ID,
	// 	fakeProductRequestUpdate.Name,
	// 	fakeProductRequestUpdate.Price,
	// 	fakeProductRequestUpdate.Description,
	// ).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	mock.ExpectExec("Update product set name = $2, price = $3, description= $4 where id = $1").WithArgs(
		fakeProductRequestUpdate.ID,
		fakeProductRequestUpdate.Name,
		fakeProductRequestUpdate.Price,
		fakeProductRequestUpdate.Description,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := productrepository.New(mock)
	product, err := sut.Update(&fakeProductRequestUpdate)

	require.NotNil(t, err)
	require.Nil(t, product)
}
