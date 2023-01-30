package userrepository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/marc/go-clean-example/adapter/postgres/userrepository"
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

	user_id := 1

	mock.ExpectExec(regexp.QuoteMeta("delete from user_api where id = $1")).
		WithArgs(uint64(user_id)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	sut := userrepository.New(mock)
	err := sut.Delete(uint64(user_id))

	require.Nil(t, err)
}

func TestDelete_DBError(t *testing.T) {

	mock := setupDelete()
	defer mock.Close()

	user_id := 1

	mock.ExpectExec(regexp.QuoteMeta("delete from user_api where id = $1")).
		WithArgs(user_id).
		WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := userrepository.New(mock)
	err := sut.Delete(uint64(user_id))
	require.NotNil(t, err)

}
