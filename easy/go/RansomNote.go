package main

import (
	"strings"
)

/*
Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 {
		return true
	}
	if len(magazine) == 0 {
		return false
	}
	for _, s := range ransomNote {
		findIndex := strings.Index(magazine, string(s))
		if findIndex == -1 {
			return false
		}
		magazine = magazine[:findIndex] + magazine[findIndex+1:]
	}
	return true
}

/*別解
func canConstruct(ransomNote string, magazine string) bool {
    for _, char := range ransomNote {
        if strings.Count(ransomNote, string(char)) > strings.Count(magazine, string(char)) {
            return false
        }
    }

    return true
}
*/
