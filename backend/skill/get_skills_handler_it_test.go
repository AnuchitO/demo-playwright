package skill

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestGetSkillsHandlerIT(t *testing.T) {
	t.Run("should return all skills from skill storage database", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:GetAllSkillsHandlerIT?mode=memory&cache=shared")
		defer db.Close()

		gosql := `INSERT INTO skill (key, name, description, levels, tags) VALUES('go','Go','Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.', '[{"name": "Beginner", "description": "I have basic knowledge of Go and can write simple programs."},{"name": "Intermediate", "description": "I can write complex programs and understand the language well."},{"name": "Advanced", "description": "I have deep knowledge of Go and can write complex programs."}]', '{go, golang}');`
		dssql := `INSERT INTO skill (key, name, description) VALUES
			('figma','Figma','Figma is a vector graphics editor and prototyping tool which is primarily web-based, with additional offline features enabled by desktop applications for macOS and Windows.'),
			('html5','HTML5','HTML5 is a markup language used for structuring and presenting content on the World Wide Web.'),
			('negotiation','Negotiation','Negotiation is a dialogue between two or more people or parties intended to reach a beneficial outcome over one or more issues where a conflict exists with respect to at least one of these issues.')  ON CONFLICT (key) DO NOTHING;`

		db.Exec(skillTable)
		db.QueryRow(gosql)
		db.QueryRow(dssql)
		s := storage{conn: db}
		h := newHandler(&s)
		c := &context{q: map[string]string{"page": "2", "perPage": "3", "orderBy": "key"}}

		h.GetSkills(c)

		resp, _ := c.ok.([]Skill)
		assert.Nil(t, c.err)
		assert.Equal(t, 1, len(resp))
		assert.Equal(t, "negotiation", resp[0].Key)
		assert.EqualValues(t, 2, c.page)
		assert.EqualValues(t, 3, c.perPage)
		assert.EqualValues(t, 4, c.totalItems)
	})
}
