package skill

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestPatchSkillHandlerIT(t *testing.T) {
	t.Run("should be able to patch skill name", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillNameHandlerIT?mode=memory&cache=shared")
		defer db.Close()

		figma := `INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`

		db.Exec(skillTable)
		db.Query(figma)

		h := newHandler(newStorage(db))
		c := &context{
			p:       map[string]string{"key": "figma"},
			payload: []byte(`{"name": "New Figma"}`),
		}

		h.PatchName(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)

		var name string
		err := db.QueryRow("SELECT name FROM skill WHERE key = 'figma';").Scan(&name)
		assert.Nil(t, err)
		assert.Equal(t, "New Figma", name)
	})

	t.Run("should be able to patch skill description", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillDesciptionHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)
		db.Query(`INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`)

		h := newHandler(newStorage(db))
		c := &context{
			p:       map[string]string{"key": "figma"},
			payload: []byte(`{"description": "New Figma is description"}`),
		}

		h.PatchDescription(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)

		var description string
		err := db.QueryRow("SELECT description FROM skill WHERE key = 'figma';").Scan(&description)
		assert.Nil(t, err)
		assert.Equal(t, "New Figma is description", description)
	})

	t.Run("should be able to patch skill logo", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillLogoHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)
		db.Query(`INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`)

		h := newHandler(newStorage(db))
		c := &context{
			p:       map[string]string{"key": "figma"},
			payload: []byte(`{"logo": "https://figma.com/logo"}`),
		}

		h.PatchLogo(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)

		var logo string
		err := db.QueryRow("SELECT logo FROM skill WHERE key = 'figma';").Scan(&logo)
		assert.Nil(t, err)
		assert.Equal(t, "https://figma.com/logo", logo)
	})

	t.Run("should be able to patch skill levels", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillLevelsHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)
		db.Query(`INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`)

		h := newHandler(newStorage(db))
		c := &context{
			p:       map[string]string{"key": "figma"},
			payload: []byte(`{"levels": [{"key": "experienced", "name": "Experienced", "brief": "Confident Contributor", "level": 3, "descriptions": ["level description"]}]}`),
		}

		h.PatchLevels(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)

		var levels []byte
		err := db.QueryRow("SELECT levels FROM skill WHERE key = 'figma';").Scan(&levels)
		assert.Nil(t, err)
		assert.Equal(t, `[{"key":"experienced","name":"Experienced","brief":"Confident Contributor","descriptions":["level description"],"level":3}]`, string(levels))
	})

	t.Run("should be able to patch skill tags", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillTagsHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)
		db.Query(`INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`)

		h := newHandler(newStorage(db))
		c := &context{
			p:       map[string]string{"key": "figma"},
			payload: []byte(`{"tags": ["tool", "design"]}`),
		}

		h.PatchTags(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)

		var tags pq.StringArray
		err := db.QueryRow("SELECT tags FROM skill WHERE key = 'figma';").Scan(&tags)
		assert.Nil(t, err)
		assert.Equal(t, pq.StringArray{"tool", "design"}, tags)
	})
}
