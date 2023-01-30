package usecase

import "testing"

func TestValidatePalindromeUseCase(t *testing.T) {
	t.Run("must return the given input with the response", func(t *testing.T) {
		// given
		input := "racecar"
		usecase := NewValidatePalindrome()
		// when
		obtained := usecase.Execute(input)
		// then
		if input != obtained.Sentence {
			t.Fatalf("the input value must be present in the output, got: '%v'", obtained.Sentence)
		}
	})

	t.Run("must return true in the output if the given input is a valid palindrome", func(t *testing.T) {
		// given
		input := "racecar"
		usecase := NewValidatePalindrome()
		// when
		obtained := usecase.Execute(input)
		// then
		if obtained.IsPalindrome != true {
			t.Fatalf("expected true but got: %t", obtained.IsPalindrome)
		}
	})

	t.Run("must return false in the output if the given input is a non-valid palindrome", func(t *testing.T) {
		// given
		input := "carlos"
		usecase := NewValidatePalindrome()
		// when
		obtained := usecase.Execute(input)
		// then
		if obtained.IsPalindrome != false {
			t.Fatalf("expected false but got: %t", obtained.IsPalindrome)
		}
	})
}
