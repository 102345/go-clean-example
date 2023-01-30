package userrepository_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/adapter/postgres/userrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupFetch() ([]string, dto.PaginationRequestParms, domain.User, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "email", "password", "createdAt"}
	fakePaginationRequestParams := dto.PaginationRequestParms{
		Page:         1,
		ItemsPerPage: 10,
		Sort:         nil,
		Descending:   nil,
		Search:       "",
	}
	fakeUserDBResponse := domain.User{}
	faker.FakeData(&fakeUserDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakePaginationRequestParams, fakeUserDBResponse, mock
}

func TestFetch(t *testing.T) {
	cols, fakePaginationRequestParams, fakeUserDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM user_api").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeUserDBResponse.ID,
			fakeUserDBResponse.Name,
			fakeUserDBResponse.Email,
			fakeUserDBResponse.Password,
			fakeUserDBResponse.CreatedAt,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM user_api").
		WillReturnRows(pgxmock.NewRows([]string{"count"}).AddRow(int32(1)))

	sut := userrepository.New(mock)
	Users, err := sut.Fetch(&fakePaginationRequestParams)

	require.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for _, User := range Users.Items.([]domain.User) {
		require.Nil(t, err)
		require.NotEmpty(t, User.ID)
		require.Equal(t, User.Name, fakeUserDBResponse.Name)
		require.Equal(t, User.Email, fakeUserDBResponse.Email)
		require.Equal(t, User.Password, fakeUserDBResponse.Password)
		require.Equal(t, User.CreatedAt, fakeUserDBResponse.CreatedAt)
	}
}

func TestFetch_QueryError(t *testing.T) {
	_, fakePaginationRequestParams, _, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM user_api").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := userrepository.New(mock)
	Users, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, Users)
}

func TestFetch_QueryCountError(t *testing.T) {
	cols, fakePaginationRequestParams, fakeUserDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM user_api").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			fakeUserDBResponse.ID,
			fakeUserDBResponse.Name,
			fakeUserDBResponse.Email,
			fakeUserDBResponse.Password,
			fakeUserDBResponse.CreatedAt,
		))

	mock.ExpectQuery("SELECT COUNT(.+) FROM user_api").
		WillReturnError(fmt.Errorf("ANY QUERY COUNT ERROR"))

	sut := userrepository.New(mock)
	Users, err := sut.Fetch(&fakePaginationRequestParams)

	require.NotNil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	require.Nil(t, Users)
}
