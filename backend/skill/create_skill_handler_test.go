package skill

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (m *mockStorage) CreateSkill(key string, skill record) error {
	return m.err
}

func TestCreateSkillHandler(t *testing.T) {
	t.Run("should response error when name is empty", func(t *testing.T) {
		mock := &mockStorage{}
		h := newHandler(mock)
		c := &context{payload: []byte(`{bad json}`)}

		h.CreateSkill(c)

		assert.NotNil(t, c.err)
		assert.Equal(t, http.StatusBadRequest, c.status)
	})
}

func TestGenerateKeyUsingName(t *testing.T) {
	t.Run("generate key name field to lowercase and repace space with underscore", func(t *testing.T) {
		name := "Figma"

		key := generateSkillKey(name)

		assert.Equal(t, "figma", key)
	})

	t.Run("generate key name field to lowercase and repace space with underscore", func(t *testing.T) {
		name := "Architecture Design"

		key := generateSkillKey(name)

		assert.Equal(t, "architecture_design", key)
	})

	t.Run("generate key name field and remove special character '&'", func(t *testing.T) {
		name := "Container & Kubernetes"

		key := generateSkillKey(name)

		assert.Equal(t, "container_kubernetes", key)
	})

	t.Run("generate key name field and remove special character '!'", func(t *testing.T) {
		name := "Container! & Kubernetes"

		key := generateSkillKey(name)

		assert.Equal(t, "container_kubernetes", key)
	})

	t.Run("generate key name field 'e2e skill created 9x23oi'", func(t *testing.T) {
		name := "e2e Skill Created 9x23oi"

		key := generateSkillKey(name)

		assert.Equal(t, "e2e_skill_created_9x23oi", key)
	})

}
