package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/single-number/
*/

func main() {
	var nums []int = []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber3(nums))
}

/**
遍历nums中的每一个元素，将每个元素添加到hashTable中，并统计出现的次数
遍历hashtable，找到出现1次的数
*/
func singleNumber(nums []int) int {
	var numMap map[int]int = make(map[int]int, 0)

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			numMap[num] += 1
		} else {
			numMap[num] = 1
		}
	}

	for k, n := range numMap {
		if n == 1 {
			return k
		}
	}
	return 0
}

/**
方法 1：列表操作
	遍历 nums 中的每一个元素
	如果列表中存在某个数字，那么删除
	如果不存在，那么添加
*/

/**
哈希表
遍历 nums 中的每一个元素
查找 hashtable 中是否有当前元素的键
如果没有，将当前元素作为键插入 hashtable
最后， hash_table 中仅有一个元素，用 popitem 获得它
*/
func singleNumber2(nums []int) int {
	var numMap map[int]int = make(map[int]int)

	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			delete(numMap, num)
		} else {
			numMap[num] = num
		}
	}

	for k := range numMap {
		return k
	}
	return 0
}

/**
位运算：异或操作
*/

func singleNumber3(nums []int) int {
	n := 0

	for _, num := range nums {
		n = n ^ num
		fmt.Println(n)
	}

	return n
}
