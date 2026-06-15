package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		assert.Equal(t, want, got)
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		got := Add(1.5, 2.5)
		want := 4.0
		assert.Equal(t, want, got)
	})
}
