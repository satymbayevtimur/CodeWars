package kata

import (
	"strings"
)

func countVowels(s string) int {
  count := 0
  vowels := "aeiou"
	
	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			count++
		}
	}
  
  return count
}

func GetCount(str string) (count int) {
  return countVowels(str)
}
