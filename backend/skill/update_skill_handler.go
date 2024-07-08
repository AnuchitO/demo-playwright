package skill

import (
	"errors"

	"demo/app"
)

func (h Handler) UpdateSkill(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var sk Skill
	if err := c.Bind(&sk); err != nil {
		c.BadRequest(err)
		return
	}

	sk.Key = key

	if err := h.skill.UpdateSkill(key, sk.toRecord()); err != nil {
		c.InternalServerError(err)
		return
	}

	c.OK(sk)
}
