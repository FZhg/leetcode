package _5

// O(n^2)
// Grow the surrounding palindrome("bab", "bbb") or two("bb") centroid characters
func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		// for even numbered palindrome like "aa"
		result = growPalindrome(i, i+1, s, result)
		// for odd numbered palindrome like "aaa", "bab"
		result = growPalindrome(i, i, s, result)
	}
	return result
}

func growPalindrome(centerLeftIndex int, centerRightIndex int, s string, result string) string {
	var palindromeSubString string
	for centerLeftIndex >= 0 && centerRightIndex < len(s) && s[centerLeftIndex] == s[centerRightIndex] {
		palindromeSubString = s[centerLeftIndex : centerRightIndex+1]
		centerRightIndex++
		centerLeftIndex--
	}

	if len(result) < len(palindromeSubString) {
		return palindromeSubString
	}
	return result
}
