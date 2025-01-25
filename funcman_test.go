package funcman_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/mishankov/funcman"
	"github.com/mishankov/testman/assert"
)

var errSample = errors.New("sample error")

func TestMap(t *testing.T) {
	got := funcman.Map([]int{1, 2, 3}, func(i int) string { return strconv.Itoa(i * 2) })
	assert.DeepEqual(t, got, []string{"2", "4", "6"})
}

func TestMapWithError(t *testing.T) {
	t.Run("test positive", func(t *testing.T) {
		got, err := funcman.MapWithError([]int{1, 2, 3}, func(i int) (string, error) { return strconv.Itoa(i * 2), nil })
		assert.NoError(t, err)
		if assert.NotEmptySlice(t, got) {
			assert.DeepEqual(t, got, []string{"2", "4", "6"})
		}
	})

	t.Run("test error", func(t *testing.T) {
		got, err := funcman.MapWithError([]int{1, 2, 3}, func(i int) (string, error) { return "", errSample })
		assert.Error(t, err)
		assert.EmptySlice(t, got)
	})
}
