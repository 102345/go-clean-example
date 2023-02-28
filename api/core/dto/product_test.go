package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/go-clean-example/core/dto"

	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	fakeItem := dto.CreateProductRequestDTO{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateProductRequestDTO(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Price, fakeItem.Price)
	require.Equal(t, itemRequest.Description, fakeItem.Description)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateProductRequestDTO(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}

func TestFromJSONUpdateProductRequest(t *testing.T) {
	fakeItem := dto.UpdateProductRequestDTO{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONUpdateProductRequestDTO(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.Equal(t, itemRequest.ID, fakeItem.ID)
	require.Equal(t, itemRequest.Name, fakeItem.Name)
	require.Equal(t, itemRequest.Price, fakeItem.Price)
	require.Equal(t, itemRequest.Description, fakeItem.Description)
}

func TestFromJSONUpdateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONUpdateProductRequestDTO(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}
