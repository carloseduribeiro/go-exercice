package usecase

import "testing"

func TestGetFibonacciUseCase(t *testing.T) {
	t.Run("should calculate the fibonacci sequence and return it", func(t *testing.T) {
		// given
		input := 10
		expected := 55
		usecase := NewGetFibonacci()
		// when
		obtained := usecase.Execute(input)
		// then
		if expected != obtained.Sequence {
			t.Fatalf("expected %d but got %d", expected, obtained)
		}
	})
}
