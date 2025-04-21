package main

/*
jsの汎用関数をGoで実装
*/
func includes[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func findIndex[T comparable](slice []T, searchElement T) int {
	for i, v := range slice {
		if v == searchElement {
			return i
		}
	}
	return -1
}

func findCount[T comparable](arr []T, target T) int {
	count := 0
	for _, num := range arr {
		if target == num {
			count++
		}
	}
	return count
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxBuyArray(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func minBuyArray(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
func findMinIndex(arr []int) int {
	min := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < min {
			min = v
			minIndex = i
		}
	}
	return minIndex
}

func filter[T comparable](arr []T, condition func(T) bool) []T {
	result := []T{}
	for _, v := range arr {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
