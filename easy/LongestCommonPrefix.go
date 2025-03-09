package main

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

func longestCommonPrefix(strs []string) string {
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		s := string(strs[0][i])
		for _, str := range strs {
			if i >= len(str) {
				return result
			}
			if string(str[i]) != s {
				return result
			}
		}
		result += s
	}
	return result
}

/* 別解
func longestCommonPrefix(strs []string) string {
	ret := strs[0]
	for i := 1; i < len(strs); i++ {

		for j := 0; j < len(ret); j++ {

			if j == len(strs[i]) {
				ret = strs[i][:j]
				break

			} else if ret[j] != strs[i][j] {
				ret = strs[i][:j]
				if ret == "" {
					return ret
				}
				break
			}
		}
	}
	return ret
}
*/
