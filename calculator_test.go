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

func TestSubtract(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		got := Subtract(5, 3)
		want := 2
		assert.Equal(t, want, got)
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		got := Subtract(5.5, 2.0)
		want := 3.5
		assert.Equal(t, want, got)
	})
}

func TestMultiply(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		got := Multiply(4, 3)
		want := 12
		assert.Equal(t, want, got)
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		got := Multiply(2.5, 4.0)
		want := 10.0
		assert.Equal(t, want, got)
	})
}

func TestDivide(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		got := Divide(10, 2)
		want := 5
		assert.Equal(t, want, got)
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		got := Divide(10.0, 4.0)
		want := 2.5
		assert.Equal(t, want, got)
	})
}
