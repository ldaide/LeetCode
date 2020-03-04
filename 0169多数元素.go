package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(nums))
}

/**
哈希表
	用哈希表来快速统计每个元素出现的次数
*/
func majorityElement(nums []int) int {

	var (
		numsMap map[int]int
		n       int
	)
	n = len(nums) / 2

	numsMap = make(map[int]int)

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			numsMap[num] += 1
			if numsMap[num] > n {
				return num
			}
		} else {
			numsMap[num] = 1
		}

	}

	return nums[0]
}

/**
排序
	将所有数字排序，取出n/2位置的数
*/
func majorityElement1(nums []int) int {

	var length int
	length = len(nums)
	sort.Ints(nums)

	fmt.Println(nums)
	return nums[length/2]
}
