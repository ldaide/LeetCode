package main

import "fmt"

/**
题目：
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
请返回 nums 的动态和。

示例 1：
输入：nums = [1,2,3,4]
输出：[1,3,6,10]
解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。

*/

/**
思路：第i个位置的和，等于 前一个位置的和 + 当前位置元素
nums = [1,2,3,4]
[1, 1+2, (1+2)+3, (1+2+3)+4]
[1, nums[0]+2, nums[1]+3, nums[2]+4]
*/

func runningSum(nums []int) []int {
	for i := range nums {
		if i != 0 {
			nums[i] += nums[i-1]
		}
	}

	return nums

}

func main() {
	nums := []int{1, 2, 3, 4}

	ret := runningSum(nums)

	fmt.Println(ret)
}
