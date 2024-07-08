package skill

import (
	"bytes"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type mockRow struct {
	data record
}

func (m *mockRow) Scan(dest ...any) error {
	*(dest[0].(*string)) = m.data.Key
	*(dest[1].(*string)) = m.data.Name
	*(dest[2].(*string)) = m.data.Description
	*(dest[3].(*string)) = m.data.Logo
	*(dest[4].(*[]byte)) = bytes.Clone(m.data.Levels)
	*(dest[5].(*pq.StringArray)) = m.data.Tags
	return nil
}

type mockRows struct {
	data record
}

func (m *mockRows) Scan(dest ...any) error {
	*(dest[0].(*string)) = m.data.Key
	*(dest[1].(*string)) = m.data.Name
	*(dest[2].(*string)) = m.data.Description
	*(dest[3].(*string)) = m.data.Logo
	*(dest[4].(*[]byte)) = bytes.Clone(m.data.Levels)
	*(dest[5].(*pq.StringArray)) = m.data.Tags
	*(dest[6].(*uint)) = 1
	return nil
}

func TestRecord(t *testing.T) {
	t.Run("San should scan data from Record to Skill correctly", func(t *testing.T) {
		rws := &mockRows{
			data: record{
				Key:         "go",
				Name:        "Go",
				Description: "description",
				Logo:        "http://logo.com/logo.png",
				Levels:      []byte(`[{"name":"Beginner","level":1},{"name":"Intermediate","level":2}]`),
				Tags:        []string{"tag1"},
			},
		}

		rec := &record{}
		var totalItems uint
		s, _ := rec.Scan(rws, &totalItems)

		sk := Skill{Key: "go", Name: "Go", Description: "description",
			Logo:   "http://logo.com/logo.png",
			Levels: []Level{{Name: "Beginner", Level: 1}, {Name: "Intermediate", Level: 2}},
			Tags:   []string{"tag1"},
		}

		assert.Equal(t, sk, s)
		assert.EqualValues(t, 1, totalItems)
	})

	t.Run("toSkill should return skill", func(t *testing.T) {
		ss := record{
			Key:         "java_script",
			Name:        "JavaScript",
			Description: "description",
			Logo:        "http://logo.com/logo.png",
			Tags:        []string{"tag1", "tag2"},
		}
		levels := []Level{{Level: 1, Descriptions: []string{"level description"}}}

		s := ss.toSkill(levels)

		assert.Equal(t, s.Key, "java_script")
		assert.Equal(t, s.Name, "JavaScript")
		assert.Equal(t, s.Description, "description")
		assert.Equal(t, s.Logo, "http://logo.com/logo.png")
		assert.Len(t, s.Levels, 1)
		assert.Len(t, s.Tags, 2)
	})

	t.Run("ScanRow: should scan data from Record to Skill correctly", func(t *testing.T) {
		rws := &mockRow{
			data: record{
				Key:         "go",
				Name:        "Go",
				Description: "description",
				Logo:        "http://logo.com/logo.png",
				Levels:      []byte(`[{"name":"Beginner","level":1},{"name":"Intermediate","level":2}]`),
				Tags:        []string{"tag1"},
			},
		}

		raw := &record{}
		s, _ := raw.ScanRow(rws)

		sk := Skill{Key: "go", Name: "Go", Description: "description",
			Logo:   "http://logo.com/logo.png",
			Levels: []Level{{Name: "Beginner", Level: 1}, {Name: "Intermediate", Level: 2}},
			Tags:   []string{"tag1"},
		}

		assert.Equal(t, sk, s)
	})

	t.Run("unmarshalLevel: should unmarshal levels", func(t *testing.T) {
		ss := record{
			Levels: []byte(`[{"name":"Beginner","level":1},{"name":"Intermediate","level":2}]`),
		}

		levels, _ := ss.unmarshalLevel()

		assert.Len(t, levels, 2)
		assert.Equal(t, levels[0].Name, "Beginner")
		assert.Equal(t, levels[0].Level, 1)
		assert.Equal(t, levels[1].Name, "Intermediate")
		assert.Equal(t, levels[1].Level, 2)
	})
}
