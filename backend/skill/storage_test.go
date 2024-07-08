package skill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOffset(t *testing.T) {
	t.Run("should return offset", func(t *testing.T) {
		assert.EqualValues(t, 0, offset(1, 10))
		assert.EqualValues(t, 10, offset(2, 10))
		assert.EqualValues(t, 20, offset(3, 10))
	})
}
