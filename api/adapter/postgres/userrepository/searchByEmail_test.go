package userrepository_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/adapter/postgres/userrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupSearchByEmail() ([]string, domain.User, pgxmock.PgxPoolIface) {

	cols := []string{"id", "password"}

	fakeUserDBResponse := domain.User{}
	faker.FakeData(&fakeUserDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeUserDBResponse, mock
}

func TestSearchByEmail(t *testing.T) {
	cols, fakeUserDBResponse, mock := setupSearchByEmail()
	defer mock.Close()

	email := "teste@teste.com"

	rows := pgxmock.NewRows(cols).AddRow(
		fakeUserDBResponse.ID,
		fakeUserDBResponse.Password,
	)

	mock.ExpectQuery("select id, password from user_api where email = \\$1").
		WithArgs(email).
		WillReturnRows(rows)

	sut := userrepository.New(mock)
	user, err := sut.SearchByEmail(email)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, user.ID, fakeUserDBResponse.ID)
	require.Equal(t, user.Password, fakeUserDBResponse.Password)

}

func TestSearchByEmail_QueryError(t *testing.T) {
	_, _, mock := setupSearchByEmail()
	defer mock.Close()

	email := "teste@teste.com"

	mock.ExpectQuery("select id, password from user_api where email = \\$1").
		WithArgs(email).
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := userrepository.New(mock)
	_, err := sut.SearchByEmail(email)

	require.NotNil(t, err)

}
