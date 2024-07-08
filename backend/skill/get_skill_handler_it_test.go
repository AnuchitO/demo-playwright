package skill

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestGetSkillByKeyHandlerIT(t *testing.T) {
	t.Run("should response the skill by key", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestGetSkillByKeyHandlerIT?mode=memory&cache=shared")
		defer db.Close()

		dssql := `INSERT INTO skill (key, name, description) VALUES
			('figma','Figma','Figma is a vector graphics editor and prototyping tool which is primarily web-based, with additional offline features enabled by desktop applications for macOS and Windows.'),
			('html5','HTML5','HTML5 is a markup language used for structuring and presenting content on the World Wide Web.') ON CONFLICT (key) DO NOTHING;`

		db.Exec(skillTable)
		db.QueryRow(dssql)
		s := storage{conn: db}
		h := newHandler(&s)
		c := &context{p: map[string]string{"key": "figma"}}

		h.GetSkillByKey(c)

		resp, _ := c.ok.(Skill)
		assert.Nil(t, c.err)
		assert.Equal(t, "figma", resp.Key)
	})
}
