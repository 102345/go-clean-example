package userrepository_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	userrepository "github.com/marc/go-clean-example/adapter/postgres/userrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupCreate() ([]string, dto.CreateUserRequest, domain.User, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "email", "password", "created_at"}
	fakeUserRequest := dto.CreateUserRequest{}
	fakeUserDBResponse := domain.User{}
	faker.FakeData(&fakeUserRequest)
	faker.FakeData(&fakeUserDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeUserRequest, fakeUserDBResponse, mock
}

func TestCreate(t *testing.T) {
	cols, fakeUserRequest, fakeUserDBResponse, mock := setupCreate()
	defer mock.Close()

	mock.ExpectQuery("INSERT INTO user_api (.+)").WithArgs(
		fakeUserRequest.Name,
		fakeUserRequest.Email,
		fakeUserRequest.Password,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeUserDBResponse.ID,
		fakeUserDBResponse.Name,
		fakeUserDBResponse.Email,
		fakeUserDBResponse.Password,
		fakeUserDBResponse.CreatedAt,
	))

	sut := userrepository.New(mock)
	User, err := sut.Create(&fakeUserRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, User.ID)
	require.Equal(t, User.Name, fakeUserDBResponse.Name)
	require.Equal(t, User.Email, fakeUserDBResponse.Email)
	require.Equal(t, User.Password, fakeUserDBResponse.Password)
}

func TestCreate_DBError(t *testing.T) {
	_, fakeUserRequest, _, mock := setupCreate()
	defer mock.Close()

	mock.ExpectQuery("INSERT INTO user_api (.+)").WithArgs(
		fakeUserRequest.Name,
		fakeUserRequest.Email,
		fakeUserRequest.Password,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := userrepository.New(mock)
	User, err := sut.Create(&fakeUserRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.NotNil(t, err)
	require.Nil(t, User)
}
