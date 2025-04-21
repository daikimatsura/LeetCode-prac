package main

import (
	"strings" 
	"unicode" 
)


/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

 

Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
 

Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters. 
*/

func isPalindrome(s string) bool {
	var cleaned strings.Builder 
	for _, r := range s {
		if unicode.IsLetter(r) {
			cleaned.WriteRune(unicode.ToLower(r)) 
		} else if unicode.IsDigit(r) {
			cleaned.WriteRune(r)
		}
	}

	cleanedStr := cleaned.String()

	left, right := 0, len(cleanedStr)-1

	for left < right {
		if cleanedStr[left] != cleanedStr[right] {
			return false
		}
		left++ 
		right--
	}

	return true
}

