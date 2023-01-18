package productrepository_test

import (
	"fmt"
	"testing"

	"github.com/marc/go-clean-example/adapter/postgres/productrepository"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupDelete() pgxmock.PgxPoolIface {

	mock, _ := pgxmock.NewPool()

	return mock
}

func TestDelete(t *testing.T) {
	mock := setupDelete()
	defer mock.Close()

	product_id := 1

	mock.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(product_id).WillReturnResult(nil)

	sut := productrepository.New(mock)
	err := sut.Delete(uint64(product_id))

	require.Nil(t, err)
}

func TestDelete_DBError(t *testing.T) {

	mock := setupDelete()
	defer mock.Close()

	product_id := 1

	mock.ExpectExec("DELETE FROM product WHERE id = ?").WithArgs(product_id).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := productrepository.New(mock)
	err := sut.Delete(uint64(product_id))
	require.NotNil(t, err)

}
