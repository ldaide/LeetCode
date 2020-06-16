package main

import (
	"fmt"
)

/**
题目：
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

/**
1. 暴力法，双层循坏遍历

2. 哈希表，将数组索引和值互相交换
*/
func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

//哈希表，通过空间换时间
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if v, ok := m[t]; ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}

func main() {
	//nums := []int{3, 2, 4}
	//target := 6
	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum1(nums, target)
	fmt.Println(res)
}
