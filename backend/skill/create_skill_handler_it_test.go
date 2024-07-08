package skill

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestCreateSkillHandlerIT(t *testing.T) {
	t.Run("should be able to create new skill", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestCreateSkillHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)

		s := storage{conn: db}
		h := newHandler(&s)
		c := &context{
			payload: []byte(`{"key": "figma", "name": "Figma", "description": "Figma is a vector bla bla"}`),
		}

		h.CreateSkill(c)

		resp, _ := c.ok.(Skill)
		assert.Nil(t, c.err)
		assert.Equal(t, "figma", resp.Key)
	})
}
