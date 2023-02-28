package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/core/dto"

	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateUserRequest(t *testing.T) {
	fakeItem := dto.CreateUserRequestDTO{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateUserRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Email, fakeItem.Email)
	require.Equal(t, itemRequest.Password, fakeItem.Password)
}

func TestFromJSONCreateUserRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateUserRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}

func TestFromJSONUpdateUserRequestDTO(t *testing.T) {
	fakeItem := dto.UpdateUserRequestDTO{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONUpdateUserRequestDTO(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.ID, fakeItem.ID)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Email, fakeItem.Email)
	require.Equal(t, itemRequest.Password, fakeItem.Password)
}

func TestFromJSONUpdateUserRequestDTO_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONUpdateUserRequestDTO(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
