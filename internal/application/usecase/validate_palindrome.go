package usecase

import "github.com/carloseduribeiro/go-exercise/internal/domain/entity"

type ValidatePalindromeUseCase struct {
}

func NewValidatePalindrome() ValidatePalindromeUseCase {
	return ValidatePalindromeUseCase{}
}

func (v ValidatePalindromeUseCase) Execute(input string) ValidatePalindromeOutputDto {
	isPalindrome := entity.IsPalindrome(input)
	return ValidatePalindromeOutputDto{
		Sentence:     input,
		IsPalindrome: isPalindrome,
	}
}

type ValidatePalindromeOutputDto struct {
	Sentence     string `json:"sentence"`
	IsPalindrome bool   `json:"isPalindrome"`
}
