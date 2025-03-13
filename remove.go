package main

import "fmt"

func remove(intSlice []int) []int {
	var result []int
	for _, num := range intSlice {
		if !contains(result, num) {
			result = append(result, num)
		}
	}
	return result
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 1, 2, 9}
	fmt.Println("Without duplicates:", remove(nums))
}


