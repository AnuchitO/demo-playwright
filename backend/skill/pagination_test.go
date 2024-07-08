package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	t.Run("should return 1 when page is empty", func(t *testing.T) {
		assert.EqualValues(t, page(""), 1)
	})
	t.Run("should return 1 when page is 0", func(t *testing.T) {
		assert.EqualValues(t, page("0"), 1)
	})

	t.Run("should return 1 when page is -1", func(t *testing.T) {
		assert.EqualValues(t, page("-1"), 1)
	})

	t.Run("should return 2 when page is 2", func(t *testing.T) {
		assert.EqualValues(t, page("2"), 2)
	})

	t.Run("should return 1 when page is not a number", func(t *testing.T) {
		assert.EqualValues(t, page("a"), 1)
	})
}

func TestItemPerPage(t *testing.T) {
	t.Run("should return 1 when item per page is empty", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage(""), 10)
	})

	t.Run("should return 1 when item per page is 0", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage("0"), 10)
	})

	t.Run("should return 1 when item per page is -1", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage("-1"), 10)
	})

	t.Run("should return 1 when item per page is 1", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage("1"), 1)
	})

	t.Run("should return 100 when item per page is 101", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage("101"), 100)
	})

	t.Run("should return 10 when item per page is 10", func(t *testing.T) {
		assert.EqualValues(t, itemsPerPage("10"), 10)
	})
}

func TestOrderBy(t *testing.T) {
	t.Run("should return key when order by is empty", func(t *testing.T) {
		assert.EqualValues(t, orderBy(""), "key")
	})

	t.Run("should return key when order by is not empty", func(t *testing.T) {
		assert.EqualValues(t, orderBy("name"), "name")
	})
}
