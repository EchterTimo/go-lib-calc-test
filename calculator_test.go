package calculator

import "testing"

func TestAdd(t *testing.T) {
	// Test case for integers
	t.Run("Integers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5
		if got != want {
			t.Errorf("Add(2, 3) = %d; want %d", got, want)
		}
	})

	// Test case for floats
	t.Run("Floats", func(t *testing.T) {
		got := Add(1.5, 2.5)
		want := 4.0
		if got != want {
			t.Errorf("Add(1.5, 2.5) = %f; want %f", got, want)
		}
	})
}
