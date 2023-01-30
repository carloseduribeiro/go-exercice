package entity

// IsPalindrome parses a string and return true if it's a palindrome
func IsPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			return false
		}
	}
	return true
}
