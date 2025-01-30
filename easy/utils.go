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
