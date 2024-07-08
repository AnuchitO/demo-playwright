package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalPage(t *testing.T) {
	t.Run("should return 2 when total items is 4 perPage 3 ", func(t *testing.T) {
		assert.EqualValues(t, 2, totalPages(3, 4))
	})

	t.Run("should return 0 when total items is 0 pagePage 10", func(t *testing.T) {
		assert.EqualValues(t, 0, totalPages(10, 0))
	})

	t.Run("should return 1 when total items is 1 pagePage 10", func(t *testing.T) {
		assert.EqualValues(t, 1, totalPages(10, 1))
	})

	t.Run("should return 1 when total items is 10 pagePage 10", func(t *testing.T) {
		assert.EqualValues(t, 1, totalPages(10, 10))
	})

	t.Run("should return 2 when total items is 11 pagePage 10", func(t *testing.T) {
		assert.EqualValues(t, 2, totalPages(10, 11))
	})
}

func TestHasPage(t *testing.T) {
	t.Run("should return 0, 2 when page is 1, perPage is 10, total is 11", func(t *testing.T) {
		var page uint = 1     // The index of the current page (starting at 1).
		var perPage uint = 10 // The number of items per page.
		var total uint = 11   // The total number of items in database.

		prevPage, nextPage := hasPage(page, perPage, total)

		assert.EqualValues(t, 0, prevPage)
		assert.EqualValues(t, 2, nextPage)
	})

	t.Run("should return 1, 0 when page is 2, perPage is 10, total is 11", func(t *testing.T) {
		var page uint = 2
		var perPage uint = 10
		var total uint = 11

		prevPage, nextPage := hasPage(page, perPage, total)

		assert.EqualValues(t, 1, prevPage)
		assert.EqualValues(t, 0, nextPage)
	})

	t.Run("should return 0, 0 when page is 1, perPage is 10, total is 10", func(t *testing.T) {
		var page uint = 1
		var perPage uint = 10
		var total uint = 10

		prevPage, nextPage := hasPage(page, perPage, total)

		assert.EqualValues(t, 0, prevPage)
		assert.EqualValues(t, 0, nextPage)
	})

	t.Run("should return 1, 3 when page is 2, perPage is 10, total is 30", func(t *testing.T) {
		var page uint = 2
		var perPage uint = 10
		var total uint = 30

		prevPage, nextPage := hasPage(page, perPage, total)

		assert.EqualValues(t, 1, prevPage)
		assert.EqualValues(t, 3, nextPage)
	})
}

func TestCountItemsSinceFirstPage(t *testing.T) {
	t.Run("should return 0 when total items is 0", func(t *testing.T) {
		var page uint = 1
		var total uint = 0

		assert.EqualValues(t, 0, totalCount(page, 10, total))
	})

	t.Run("should return 10 when total items is 10 and page is 1", func(t *testing.T) {
		var page uint = 1
		var perPage uint = 10
		var total uint = 10

		assert.EqualValues(t, 10, totalCount(page, perPage, total))
	})

	t.Run("should return 20 when total items is 10 and page is 2", func(t *testing.T) {
		var page uint = 2
		var perPage uint = 10
		var total uint = 20

		assert.EqualValues(t, 20, totalCount(page, perPage, total))
	})

	t.Run("should return 15 when total items is 10 and page is 2", func(t *testing.T) {
		var page uint = 2
		var perPage uint = 10
		var total uint = 15

		assert.EqualValues(t, 15, totalCount(page, perPage, total))
	})

	t.Run("should return 7 when total items is 7 and page is 1", func(t *testing.T) {
		var page uint = 1
		var perPage uint = 10
		var total uint = 7

		assert.EqualValues(t, 7, totalCount(page, perPage, total))
	})
}

func TestPaging(t *testing.T) {
	t.Run("should return Paging struct", func(t *testing.T) {
		var page uint = 1
		var perPage uint = 10
		var total uint = 11

		got := paging(page, perPage, total)

		want := Paging{Page: 1, Total: 11, TotalPages: 2, TotalCount: 10, NextPage: 2, PrevPage: 0}

		assert.Equal(t, want, got)
	})
}
