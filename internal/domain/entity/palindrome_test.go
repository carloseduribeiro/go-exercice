package entity

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input       string
		expected    bool
		description string
	}{
		{input: "racecar", expected: true, description: "must return true when given input is racecar"},
		{input: "madam", expected: true, description: "must return true when given input is madam"},
		{input: "carlos", expected: false, description: "must return false when given input is carlos"},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			obtained := IsPalindrome(tt.input)
			if tt.expected != obtained {
				t.Fatalf("should return %t but got %t", tt.expected, obtained)
			}
		})
	}
}
