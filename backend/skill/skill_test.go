package skill

import (
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestSkill(t *testing.T) {
	t.Run("toRecord: should return Record", func(t *testing.T) {
		s := Skill{
			Key:         "go",
			Name:        "Go",
			Description: "description",
			Logo:        "http://logo.com/logo.png",
			Tags:        []string{"tag1", "tag2"},
			Levels: []Level{{
				Key: "beginner", Name: "Beginner", Brief: "Growing Developer",
				Level: 1, Descriptions: []string{"level description"},
			}},
		}

		ss := s.toRecord()

		assert.Equal(t, ss.Key, "go")
		assert.Equal(t, ss.Name, "Go")
		assert.Equal(t, ss.Description, "description")
		assert.Equal(t, ss.Logo, "http://logo.com/logo.png")
		assert.JSONEq(t, `[{
		"key":"beginner","name":"Beginner","brief":"Growing Developer",
		"descriptions":["level description"],"level":1}]`, string(ss.Levels))
		assert.Equal(t, ss.Tags, pq.StringArray{"tag1", "tag2"})
	})

	t.Run("levels: should return json levels", func(t *testing.T) {
		s := Skill{
			Levels: []Level{{
				Key: "experienced", Name: "Experienced", Brief: "Confident Contributor",
				Level: 3, Descriptions: []string{"level description"}}},
		}

		levels := s.levels()

		assert.JSONEq(t, `[{
			"key":"experienced","name":"Experienced","brief":"Confident Contributor",
			"descriptions":["level description"],"level":3}]`, string(levels))
	})

	t.Run("levels: should return empty array", func(t *testing.T) {
		s := Skill{}

		levels := s.levels()

		assert.JSONEq(t, `[]`, string(levels))
	})

	t.Run("tags: should return json tags", func(t *testing.T) {
		s := Skill{
			Tags: []string{"tag1", "tag2"},
		}

		tags := s.tags()

		assert.Equal(t, pq.StringArray{"tag1", "tag2"}, tags)
	})

	t.Run("tags: should return empty array", func(t *testing.T) {
		s := Skill{}

		tags := s.tags()

		assert.Equal(t, pq.StringArray{}, tags)
	})

	t.Run("toRecord: should return default value", func(t *testing.T) {
		s := Skill{}

		ss := s.toRecord()

		assert.Equal(t, ss.Key, "")
		assert.Equal(t, ss.Name, "")
		assert.Equal(t, ss.Description, "")
		assert.Equal(t, ss.Logo, "")
		assert.JSONEq(t, `[]`, string(ss.Levels))
		assert.Equal(t, ss.Tags, pq.StringArray{})
	})

}
