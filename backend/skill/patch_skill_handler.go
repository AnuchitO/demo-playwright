package skill

import (
	"errors"

	"demo/app"
)

func (h Handler) PatchName(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var name struct {
		Name string `json:"name" binding:"required"` //TODO: solution: the binding tag not able to test from integration test here because it need gin framework context
	}
	if err := c.Bind(&name); err != nil {
		c.BadRequest(err)
		return
	}

	if err := h.skill.PatchName(key, name.Name); err != nil {
		c.InternalServerError(err)
		return
	}

	c.UPDATED()
}

func (h Handler) PatchDescription(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var desc struct {
		Description string `json:"description" binding:"required"`
	}
	if err := c.Bind(&desc); err != nil {
		c.BadRequest(err)
		return
	}

	if err := h.skill.PatchDescription(key, desc.Description); err != nil {
		c.InternalServerError(err)
		return
	}

	c.UPDATED()
}

func (h Handler) PatchLogo(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var logo struct {
		Logo string `json:"logo" binding:"required"`
	}
	if err := c.Bind(&logo); err != nil {
		c.BadRequest(err)
		return
	}

	if err := h.skill.PatchLogo(key, logo.Logo); err != nil {
		c.InternalServerError(err)
		return
	}

	c.UPDATED()
}

func (h Handler) PatchLevels(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var lv struct {
		Levels []Level `json:"levels" binding:"required"`
	}
	if err := c.Bind(&lv); err != nil {
		c.BadRequest(err)
		return
	}

	if err := h.skill.PatchLevels(key, levels(lv.Levels)); err != nil {
		c.InternalServerError(err)
		return
	}

	c.UPDATED()

}

func (h Handler) PatchTags(c app.Context) {
	key := c.Param("key")
	if key == "" {
		c.BadRequest(errors.New("key is required"))
		return
	}

	var tg struct {
		Tags []string `json:"tags" binding:"required"`
	}
	if err := c.Bind(&tg); err != nil {
		c.BadRequest(err)
		return
	}

	if err := h.skill.PatchTags(key, tags(tg.Tags)); err != nil {
		c.InternalServerError(err)
		return
	}

	c.UPDATED()
}
