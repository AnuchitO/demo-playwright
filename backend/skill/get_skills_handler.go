package skill

import (
	"demo/app"
)

func (h Handler) GetSkills(c app.Context) {
	page := page(c.Query("page"))
	perPage := itemsPerPage(c.Query("perPage"))
	orderBy := orderBy(c.Query("orderBy"))

	skills, totalItems, err := h.skill.GetAllSkills(orderBy, page, perPage)
	if err != nil {
		// TODO: handle db error message
		c.InternalServerError(err)
		return
	}

	c.OkWithPagination(skills, page, perPage, totalItems)
}
