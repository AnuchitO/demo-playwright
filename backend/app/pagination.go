package app

// totalPages is a function to calculate total totalPages
// perPage	The number of items per page.
// total	The total number of items.
func totalPages(perPage uint, total uint) uint {
	pages := total / perPage
	if total%perPage > 0 {
		return pages + 1
	}
	return pages
}

// totalCount is a function to calculate total items since the first page
// page	The index of the current page (starting at 1).
// perPage	The number of items per page.
// total	The total number of items.
func totalCount(page, perPage, total uint) uint {
	if total == 0 {
		return 0
	}

	if total <= (page * perPage) {
		return total
	}

	return page * perPage
}

// hasPage is a function to calculate next and previous page
// page	The index of the current page (starting at 1).
// perPage	The number of items per page.
// total	The total number of items.
func hasPage(page, perPage, total uint) (prevPage, nextPage uint) {
	if page > 1 {
		prevPage = page - 1
	}

	if total > page*perPage {
		nextPage = page + 1
	}

	return
}

// paging is a function to calculate paging
// page	The index of the current page (starting at 1).
// perPage	The number of items per page.
// total	The total number of items.
func paging(page, perPage, total uint) Paging {
	prevPage, nextPage := hasPage(page, perPage, total)
	return Paging{
		Page:       page,
		Total:      total,
		TotalPages: totalPages(perPage, total),
		TotalCount: totalCount(page, perPage, total),
		NextPage:   nextPage,
		PrevPage:   prevPage,
	}
}
