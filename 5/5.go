package _5

import (
	"strings"
)

/*----- Brutal-force -----*/
// O(n^2)
// Grow the surrounding palindrome("bab", "bbb") or two("bb") centroid characters
func brutalForceLongestPalindromeSubstring(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		// for even numbered palindrome like "aa"
		result = growPalindrome1(i, i+1, s, result)
		// for odd numbered palindrome like "aaa", "bab"
		result = growPalindrome1(i, i, s, result)
	}
	return result
}

func growPalindrome1(centerLeftIndex int, centerRightIndex int, s string, result string) string {
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

/*----- Manachar's Algorithm -----*/
func manacharLongestPalindromeSubstring(s string) string {
	stubString := generateStubString(s)
	palindromeLengths := make([]int, len(stubString))
	palindromeLengths[0] = 1 // special char "@"
	palindromeLengths[1] = 1 // first stub char "#"
	pivot, pivotRightLimit := 1, 1
	for center := 2; center <= len(stubString)-3; center++ {
		// check if the pivotRightLimit Can't hold
		// this palindrome substring
		centerMirrored := mirrorPosition(pivot, center)
		centerRightLimit := getCenterRightLimit(center, palindromeLengths[centerMirrored])
		if centerRightLimit >= pivotRightLimit {
			// grow the palindrome
			centerRightIndex := pivotRightLimit + 1 // IMPORTANT: the chars beyond pivotRightLimit + 1  are not symmetric
			centerLeftIndex := mirrorPosition(center, centerRightIndex)
			centerRightLimit = growPalindrome(centerLeftIndex, centerRightIndex, stubString)
			palindromeLength := (centerRightLimit-center)*2 + 1
			palindromeLengths[center] = palindromeLength

			// update the pivot
			if pivotRightLimit <= centerRightLimit {
				pivot = center
				pivotRightLimit = centerRightLimit
			}
		} else {
			// just mirror
			palindromeLengths[center] = palindromeLengths[centerMirrored]
			// no need to compute by symmetry of palindrome
		}
	}

	// linear search for LPS
	maxLength := 0
	center := 0
	for index, length := range palindromeLengths {
		if length > maxLength {
			maxLength = length
			center = index
		}
	}

	// strip the stub characters
	// TODO: calculate the relation of stub string index and original string index
	LPSStubString := stubString[center-maxLength/2 : center+maxLength/2]
	LPS := strings.Split(LPSStubString, "#")

	return strings.Join(LPS, "")
}

/*
*
growPalindrome grow the palindrome and returns the central right limit
*/
func growPalindrome(centerLeftIndex int, centerRightIndex int, stubString string) int {
	centerRightLimit := centerRightIndex - 1
	for centerLeftIndex >= 1 && centerRightIndex <= len(stubString)-2 && stubString[centerRightIndex] == stubString[centerLeftIndex] {
		centerRightLimit = centerRightIndex
		centerRightIndex++
		centerLeftIndex--

	}
	return centerRightLimit
}

func mirrorPosition(center int, position int) int {
	return center - (position - center)
}

/*
palindromeSubstringLength must be odd int.
*/
func getCenterRightLimit(center int, palindromeSubstringLength int) int {
	return center + palindromeSubstringLength/2 // integer division only returns the quotient
}

/*
*
"ababac" => "@#a#b#a#b#a#c#$"
"#", the stub character is where the even length palindrome grows
*/
func generateStubString(s string) string {
	// Add two distinct characters at the start and end of the stub string
	// so that index is always in bound
	stubString := "@#"

	chars := strings.Split(s, "")
	stubString += strings.Join(chars, "#")
	stubString += "#$"
	return stubString
}
