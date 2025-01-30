package main

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true
Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false


Constraints:

0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
*/

func isSubsequence(s string, t string) bool {
	tRunes := []rune(t)
	for _, str := range s {
		index := findIndex(tRunes, rune(str))
		if index == -1 {
			return false
		} else {
			tRunes = tRunes[index+1:]
		}
	}
	return true
}

/*別解
func isSubsequence(s string, t string) bool {
    sIdx, tIdx := 0, 0

    for sIdx < len(s) && tIdx < len(t) {
        if s[sIdx] == t[tIdx] {
            sIdx++
        }

        tIdx++
    }

    return sIdx == len(s)
}
*/
