package main

import "fmt"

func uniqueOccurrences(arr []int) bool {

	hash := make(map[int]int)

	for _, v := range arr {
		hash[v]++
	}

	temp := make(map[int]int)
	for _, v := range hash {
		if _, ok := temp[v]; ok {
			return false
		}
		temp[v] = 1
	}

	return true
}

func main() {
	num1 := []int{1, 2}
	fmt.Println(uniqueOccurrences(num1))
}
