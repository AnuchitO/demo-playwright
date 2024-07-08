package skill

import (
	"errors"

	"demo/app"
)

func (h Handler) GetSkillByKey(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	skill, err := h.skill.GetSkillByKey(key)
	if err != nil {
		c.InternalServerError(err)
		return
	}

	c.OK(skill)
}
