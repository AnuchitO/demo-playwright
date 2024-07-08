package skill

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *mockStorage) PatchName(key string, name string) error {
	return m.err
}

func TestPatchSkillName(t *testing.T) {
	t.Run("should response bad request when key is empty", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": ""}}

		h.PatchName(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusBadRequest, c.status)
		assert.Equal(t, errors.New("key is required"), c.err)
	})

	t.Run("should response bad request when bind name error", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": "html5"}, payload: []byte(`{bad json}`)}

		h.PatchName(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusBadRequest, c.status)
	})

	t.Run("should response internal server error when storage error", func(t *testing.T) {
		mock := &mockStorage{err: errors.New("database error")}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": "html5"}, payload: []byte(`{"name": "HTML5"}`)}

		h.PatchName(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusInternalServerError, c.status)
	})

	t.Run("should response ok with no contect when patch name success", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{p: map[string]string{"key": "html5"}, payload: []byte(`{"name": "HTML5"}`)}

		h.PatchName(c)

		assert.Nil(t, c.err)
		assert.Equal(t, http.StatusNoContent, c.status)
	})
}
