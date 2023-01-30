package entity

import "testing"

func TestGetFibonacci(t *testing.T) {
	t.Run("must return 0 when the given input is 0", func(t *testing.T) {
		// given
		input := 0
		expected := input
		// when
		obtained := Fibonacci(input)
		// then
		if expected != obtained {
			t.Fatalf("expected %d but got %d", expected, obtained)
		}
	})

	t.Run("must return 1 when the given input is 1", func(t *testing.T) {
		// given
		input := 1
		expected := input
		// when
		obtained := Fibonacci(input)
		// then
		if expected != obtained {
			t.Fatalf("expected %d but got %d", expected, obtained)
		}
	})

	t.Run("must calculate the nth fibonacci number", func(t *testing.T) {
		// given
		input := 10
		expected := 55
		// when
		obtained := Fibonacci(input)
		// then
		if expected != obtained {
			t.Fatalf("expected %d but got %d", expected, obtained)
		}
	})

}
