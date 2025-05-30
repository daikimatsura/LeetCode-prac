package main

/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


Follow-up: Could you solve the problem in linear time and in O(1) space?
*/

func majorityElement(nums []int) int {
	elem := make(map[int]int)
	for _, num := range nums {
		elem[num]++
	}
	for num, count := range elem {
		if count > len(nums)/2 {
			return num
		}
	}
	return 0
}

/* 別解
func majorityElement(nums []int) int {
	// ボイヤー・ムーア多数決アルゴリズムを使用
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	return candidate
}
*/
