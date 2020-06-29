package main

import "fmt"

/*
题目：
给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
以数组形式返回答案。

示例 1：
输入：nums = [8,1,2,2,3]
输出：[4,0,1,1,3]
解释：
对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
对于 nums[1]=1 不存在比它小的数字。
对于 nums[2]=2 存在一个比它小的数字：（1）。
对于 nums[3]=2 存在一个比它小的数字：（1）。
对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。

*/

//双重循环
func smallerNumbersThanCurrent(nums []int) []int {
	l := len(nums)
	ret := make([]int, l)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				ret[i]++
			} else if nums[i] < nums[j] {
				ret[j]++
			}
		}
	}
	return ret
}

func smallerNumbersThanCurrent1(nums []int) []int {
	cnt := make(map[int]int)
	l := 0
	//求出每个元素出现的次数
	for _, v := range nums {
		cnt[v]++
		l++
	}

	for i := 1; i <= 100; i++ {
		cnt[i] += cnt[i-1]
	}

	vec := make([]int, l)
	for i, _ := range nums {
		vec[i] = cnt[nums[i]-1]
	}
	return vec
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent1(nums))

}
