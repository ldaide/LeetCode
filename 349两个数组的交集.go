package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	temp := make(map[int]int)
	for _, v := range nums1 {
		for _, v1 := range nums2 {
			if v == v1 {
				temp[v] = 1
			}
		}
	}
	ret := make([]int, 0)
	for k, _ := range temp {
		ret = append(ret, k)
	}
	return ret
}

func intersection1(nums1 []int, nums2 []int) []int {
	temp1 := make(map[int]int)
	temp2 := make(map[int]int)
	for _, v := range nums1 {
		temp1[v] = 1
	}
	for _, v := range nums2 {
		temp2[v] = 1
	}
	ret := make([]int, 0)
	for k, _ := range temp1 {
		if _, ok := temp2[k]; ok {
			ret = append(ret, k)
		}
	}
	return ret
}

func intersection2(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	ret := make([]int, 0)
	for _, v := range nums1 {
		hash[v] = true
	}
	for _, v := range nums2 {
		if hash[v] == true {
			ret = append(ret, v)
			hash[v] = false
		}
	}
	return ret
}

func main() {
	num1 := []int{1, 2, 2, 1}
	num2 := []int{2, 2}
	fmt.Println(intersection2(num1, num2))
}
