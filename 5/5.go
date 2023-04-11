package _5

// Brutal-Force: expand the parlindrom from the center
// Credit: https://www.youtube.com/watch?v=XYQecbcd6_c
// TODO: Dynamic Programming

import "log"

func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		centerLeftIndex, centerRightIndex := i, i
		if i != len(s)-1 && s[i] == s[i+1] {
			centerRightIndex = i + 1
			// for edge case: bb
		}

		for centerLeftIndex >= 0 && centerRightIndex < len(s) {
			log.Printf("centerLeftIndex: %d, centerRightIndex: %d", centerLeftIndex, centerRightIndex)
			if s[centerLeftIndex] == s[centerRightIndex] {
				if len(result) < centerRightIndex-centerLeftIndex+1 {
					result = s[centerLeftIndex : centerRightIndex+1]
				}
			} else {
				break // the palindrome can not grow
				// for example, dabac. d and c are not equal.
			}
			// expand the palindrome
			centerLeftIndex--
			centerRightIndex++
		}

	}

	return result
}
