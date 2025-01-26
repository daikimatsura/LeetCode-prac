package main

/*
jsの汎用関数をGoで実装
*/
func includes(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func findIndex(targetStr string, searchStr string) int {
	for i, str := range targetStr {
		if string(str) == string(searchStr) {
			return i
		}
	}
	return -1
}
