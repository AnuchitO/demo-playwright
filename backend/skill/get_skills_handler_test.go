package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct {
	storager
	err     error
	skill   Skill
	orderBy string
	page    uint
	perPage uint
}

func (m *mockStorage) GetAllSkills(orderBy string, page, limit uint) ([]Skill, uint, error) {
	m.orderBy = orderBy
	m.page = page
	m.perPage = limit

	if m.err != nil {
		return nil, 0, m.err
	}

	return []Skill{
		{
			Key:         "key",
			Name:        "name",
			Description: "description",
			Logo:        "logo",
			Levels:      []Level{},
			Tags:        []string{"tag1"},
		},
	}, 1, nil
}

func TestGetAllSkill(t *testing.T) {
	t.Run("should response all skills from storage", func(t *testing.T) {
		h := newHandler(&mockStorage{})

		c := &context{}
		h.GetSkills(c)

		resp, _ := c.ok.([]Skill)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("should response all skills with pagination from storage", func(t *testing.T) {
		spy := &mockStorage{}
		h := newHandler(spy)

		c := &context{q: map[string]string{"page": "2", "perPage": "20", "orderBy": "name"}}
		h.GetSkills(c)

		assert.Equal(t, "name", spy.orderBy)
		assert.EqualValues(t, 2, spy.page)
		assert.EqualValues(t, 20, spy.perPage)
	})

	t.Run("should get all skill by default query pagination", func(t *testing.T) {
		spy := &mockStorage{}
		h := newHandler(spy)

		c := &context{}
		h.GetSkills(c)

		assert.Equal(t, "key", spy.orderBy)
		assert.EqualValues(t, 1, spy.page)
		assert.EqualValues(t, 10, spy.perPage)
	})

	t.Run("should response error when storage error", func(t *testing.T) {
		fake := &mockStorage{err: assert.AnError}
		h := newHandler(fake)

		c := &context{}
		h.GetSkills(c)

		assert.Equal(t, assert.AnError, c.err)
	})
}
