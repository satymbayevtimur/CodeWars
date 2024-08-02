package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	left := 0
	right := len(str) - 1

	for left < right {
		for left < right && !isAlphaNumeric(str[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(str[right]) {
			right--
		}

		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}
