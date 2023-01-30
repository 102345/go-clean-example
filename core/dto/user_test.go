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
	fakeItem := dto.CreateUserRequest{}
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

func TestFromJSONUpdateUserRequest(t *testing.T) {
	fakeItem := dto.UpdateUserRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONUpdateUserRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.ID, fakeItem.ID)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Email, fakeItem.Email)
	require.Equal(t, itemRequest.Password, fakeItem.Password)
}

func TestFromJSONUpdateUserRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONUpdateUserRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
