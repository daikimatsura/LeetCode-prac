package main

import (
	"strconv"
)

/*
Given an integer x, return true if x is a
palindrome
, and false otherwise.



Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1


Follow up: Could you solve it without converting the integer to a string?
*/

func isPalindrome(x int) bool {
	// 負の数は回文になり得ない
	if x < 0 {
		return false
	}

	// 整数を文字列に変換（strconvパッケージを使用）
	s := strconv.Itoa(x)

	// 文字列の長さを取得
	n := len(s)

	// 文字列の前半と後半を比較
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
