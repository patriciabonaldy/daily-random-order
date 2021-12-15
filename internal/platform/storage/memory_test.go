package storage_test

import (
	"daily-random-order/internal/platform/storage"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryRepository_FindAllCoWorkers_Success(t *testing.T) {
	repository := storage.NewMemoryRepository("data-mock/team.json")
	expected := []string{"Patricia", "Jayson"}
	got, err := repository.FindAllCoWorkers()
	require.NoError(t, err)
	assert.Equal(t, got, expected)
}

func TestMemoryRepository_FindAllCoWorkers_BadPath(t *testing.T) {
	repository := storage.NewMemoryRepository("../data/team.json")
	var expected []string
	got, err := repository.FindAllCoWorkers()
	require.Error(t, err)
	assert.Equal(t, got, expected)
}

func TestMemoryRepository_FindAllCoWorkers_EmptyFile(t *testing.T) {
	repository := storage.NewMemoryRepository("data-mock/empty.json")
	var expected []string
	got, err := repository.FindAllCoWorkers()
	assert.Equal(t, err, storage.ErrorFileDoesntExist)
	assert.Equal(t, got, expected)
}
