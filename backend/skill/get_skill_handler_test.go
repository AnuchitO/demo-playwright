package skill

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *mockStorage) GetSkillByKey(key string) (Skill, error) {
	return m.skill, m.err
}

func TestGetSkillByKey(t *testing.T) {
	t.Run("should response error when key is empty", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{}

		h.GetSkillByKey(c)

		assert.Equal(t, errors.New("key is required"), c.err)
	})

	t.Run("should response error when storage return error", func(t *testing.T) {
		mock := &mockStorage{err: errors.New("error")}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": "html5"}}

		h.GetSkillByKey(c)

		assert.Equal(t, errors.New("error"), c.err)
	})

	t.Run("should response all skills from storage", func(t *testing.T) {
		mock := &mockStorage{
			skill: Skill{
				Key: "html5", Name: "HTML5", Description: "description",
				Logo: "logo", Levels: []Level{}, Tags: []string{"tag"},
			},
		}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": "html5"}}

		h.GetSkillByKey(c)

		resp, _ := c.ok.(Skill)
		assert.Equal(t, "html5", resp.Key)
	})
}
