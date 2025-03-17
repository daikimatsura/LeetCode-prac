package main

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
Note that the integers in the lists may be returned in any order.

Example 1:

Input: nums1 = [1,2,3], nums2 = [2,4,6]
Output: [[1,3],[4,6]]
Explanation:
For nums1, nums1[1] = 2 is present at index 0 of nums2, whereas nums1[0] = 1 and nums1[2] = 3 are not present in nums2. Therefore, answer[0] = [1,3].
For nums2, nums2[0] = 2 is present at index 1 of nums1, whereas nums2[1] = 4 and nums2[2] = 6 are not present in nums2. Therefore, answer[1] = [4,6].
Example 2:

Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
Output: [[3],[]]
Explanation:
For nums1, nums1[2] and nums1[3] are not present in nums2. Since nums1[2] == nums1[3], their value is only included once and answer[0] = [3].
Every integer in nums2 is present in nums1. Therefore, answer[1] = [].

Constraints:

1 <= nums1.length, nums2.length <= 1000
-1000 <= nums1[i], nums2[i] <= 1000
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	answer := [][]int{[]int{}, []int{}}
	seen1 := make(map[int]bool)
	seen2 := make(map[int]bool)

	for _, num := range nums1 {
		seen1[num] = true
	}
	for _, num := range nums2 {
		seen2[num] = true
	}

	for num := range seen1 {
		if !seen2[num] {
			answer[0] = append(answer[0], num)
		}
	}
	for num := range seen2 {
		if !seen1[num] {
			answer[1] = append(answer[1], num)
		}
	}

	return answer
}

/*別解
func findDifference(nums1 []int, nums2 []int) [][]int {
	uniqTo2 := uniqToArray(nums1, nums2)
	uniqTo1 := uniqToArray(nums2, nums1)
	return [][]int{uniqTo1, uniqTo2}
}

func uniqToArray(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool)
	for i := 0; i < len(nums1); i ++ {
		nums1Map[nums1[i]] = true
	}
	uniqTo2 := map[int]bool{}
	uniqArray := []int{}
	for i := 0; i < len(nums2); i ++ {
		if !nums1Map[nums2[i]] {
			if !uniqTo2[nums2[i]]{
				uniqArray = append(uniqArray, nums2[i])
			}
			uniqTo2[nums2[i]] = true
		}
	}
	return uniqArray
}
*/
