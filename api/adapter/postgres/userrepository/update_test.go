package userrepository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/adapter/postgres/userrepository"
	"github.com/marc/go-clean-example/core/domain"
	"github.com/marc/go-clean-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupUpdate() (dto.UpdateUserRequestDTO, domain.User, pgxmock.PgxPoolIface) {
	fakeUserRequest := dto.UpdateUserRequestDTO{}
	fakeUserDBResponse := domain.User{}
	faker.FakeData(&fakeUserRequest)
	faker.FakeData(&fakeUserDBResponse)

	mock, _ := pgxmock.NewPool()

	return fakeUserRequest, fakeUserDBResponse, mock
}

func TestUpdate(t *testing.T) {
	fakeUserRequest, _, mock := setupUpdate()
	defer mock.Close()
	mock.ExpectExec(regexp.QuoteMeta("Update user_api set name = $2, email = $3, password= $4 where id = $1")).WithArgs(
		fakeUserRequest.ID,
		fakeUserRequest.Name,
		fakeUserRequest.Email,
		fakeUserRequest.Password,
	).WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	sut := userrepository.New(mock)

	User, err := sut.Update(&fakeUserRequest)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, User.ID, fakeUserRequest.ID)
	require.Equal(t, User.Name, fakeUserRequest.Name)
	require.Equal(t, User.Email, fakeUserRequest.Email)
	require.Equal(t, User.Password, fakeUserRequest.Password)
}

func TestUpdate_DBError(t *testing.T) {
	_, _, mock := setupUpdate()

	fakeUserRequestUpdate := dto.UpdateUserRequestDTO{}
	defer mock.Close()

	mock.ExpectExec(regexp.QuoteMeta("Update user_api set name = $2, email = $3, password= $4 where id = $1")).WithArgs(
		fakeUserRequestUpdate.ID,
		fakeUserRequestUpdate.Name,
		fakeUserRequestUpdate.Email,
		fakeUserRequestUpdate.Password,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := userrepository.New(mock)
	User, err := sut.Update(&fakeUserRequestUpdate)

	require.NotNil(t, err)
	require.Nil(t, User)
}
