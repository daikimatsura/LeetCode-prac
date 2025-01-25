package main

/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

Example 1:

Input: s = "IceCreAm"

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"

Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
*/
var vowels = []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

func reverseVowels(s string) string {
	sBytes := []byte(s)
	targetVowels := getVowels(s)
	targetVowelsId := 0
	for i, str := range s {
		if includes(vowels, string(str)) {
			sBytes[i] = targetVowels[len(targetVowels)-1-targetVowelsId]
			targetVowelsId++
		}
	}
	return string(sBytes)
}

func includes(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func getVowels(s string) string {
	var result string
	for _, str := range s {
		if includes(vowels, string(str)) {
			result += string(str)
		}
	}
	return result
}

/* もっと実行速度が速い例
func reverseVowels(s string) string {
	word := []rune(s)
	start := 0
	end := len(word) - 1
	vowels := "aeiouAEIOU"

	for start < end {
		for start < end && !strings.ContainsRune(vowels, word[start]) {
			start++
		}
		for start < end && !strings.ContainsRune(vowels, word[end]) {
			end--
		}

		word[start], word[end] = word[end], word[start]

		start++
		end--
	}

	return string(word)
}
*/
