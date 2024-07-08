package skill

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSkillHandlerIT(t *testing.T) {
	t.Run("should able to update skill succesfully", func(t *testing.T) {
		db, _ := sql.Open("sqlite", "file:TestPatchSkillDesciptionHandlerIT?mode=memory&cache=shared")
		defer db.Close()
		db.Exec(skillTable)
		db.Query(`INSERT INTO skill (key, name, description) VALUES ('figma','Figma','Figma is something');`)

		key := "figma"
		h := newHandler(newStorage(db))
		c := &context{
			p: map[string]string{"key": key},
			payload: []byte(`{
				"name": "New Figma",
				"description": "New Figma is description",
				"logo": "https://figma.com/logo.png",
				"levels": [{"key": "experienced", "name": "Experienced", "brief": "Confident Contributor", "level": 3, "descriptions": ["level description"]}],
				"tags": ["design", "prototyping"]
				}`),
		}

		h.UpdateSkill(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusOK, c.status)

		ss := &record{}
		row := db.QueryRow("SELECT key, name, description, logo, levels, tags FROM skill WHERE key = $1", key)
		sk, err := ss.ScanRow(row)
		assert.Nil(t, err)
		assert.Equal(t, key, sk.Key)
		assert.Equal(t, "New Figma", sk.Name)
		assert.Equal(t, "New Figma is description", sk.Description)
		assert.Equal(t, "https://figma.com/logo.png", sk.Logo)
		assert.Greater(t, len(sk.Levels), 0)
		assert.Greater(t, len(sk.Tags), 1)
	})
}
