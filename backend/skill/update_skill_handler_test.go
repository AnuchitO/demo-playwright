package skill

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *mockStorage) UpdateSkill(key string, skill record) error {
	return m.err
}

func TestUpdateSkillHandler(t *testing.T) {
	t.Run("should response error when key is empty", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": ""}}

		h.UpdateSkill(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusBadRequest, c.status)
	})

	t.Run("should response error when bind payload error", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{
			p:       map[string]string{"key": "html5"},
			payload: []byte(`{bad json}`)}

		h.UpdateSkill(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusBadRequest, c.status)
	})

	t.Run("should response error when storage error", func(t *testing.T) {
		mock := &mockStorage{err: errors.New("database error")}
		h := newHandler(mock)
		c := &context{
			p:       map[string]string{"key": "html5"},
			payload: []byte(`{"name": "HTML5"}`)}

		h.UpdateSkill(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusInternalServerError, c.status)
	})

	t.Run("should response ok when update success", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{
			p:       map[string]string{"key": "html5"},
			payload: []byte(`{"name": "HTML5"}`)}

		h.UpdateSkill(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusOK, c.status)
		assert.Equal(t, "html5", c.ok.(Skill).Key)
		assert.Equal(t, "HTML5", c.ok.(Skill).Name)
	})
}
