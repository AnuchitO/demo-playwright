package skill

import (
	"strings"
	"unicode"

	"demo/app"
)

func (h Handler) CreateSkill(c app.Context) {
	var skill Skill
	if err := c.Bind(&skill); err != nil {
		c.BadRequest(err)
		return
	}

	key := generateSkillKey(skill.Name)
	if err := h.skill.CreateSkill(key, skill.toRecord()); err != nil {
		c.InternalServerError(err)
		return
	}
	skill.Key = key
	c.CREATED(skill)
}

func generateSkillKey(name string) string {
	name = strings.ReplaceAll(name, "&", "")
	name = strings.ReplaceAll(name, "!", "")

	var parts []string
	start := 0
	for end, r := range name {
		if end != 0 && unicode.IsUpper(r) {
			word := strings.TrimSpace(name[start:end])
			parts = append(parts, word)
			start = end
		}
	}
	if start != len(name) {
		parts = append(parts, name[start:])
	}

	key := strings.ToLower(strings.Join(parts, "_"))
	return strings.ReplaceAll(key, " ", "_")
}
