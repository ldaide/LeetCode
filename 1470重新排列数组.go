package main

import "fmt"

/**
题目：
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。

示例 1：
输入：nums = [2,5,1,3,4,7], n = 3
输出：[2,3,5,4,1,7]
解释：由于 x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 ，所以答案为 [2,3,5,4,1,7]
*/

//方法一
func shuffle(nums []int, n int) []int {
	ret := make([]int, 2*n)
	fmt.Println(ret)
	for i := 0; i < n; i++ {
		ret[2*i] = nums[i]
		ret[2*i+1] = nums[n+i]
	}
	return ret

}

//方法二，这种方法比较耗时，没有方法一效率高
func shuffle1(nums []int, n int) []int {
	ret := make([]int, 0, 2*n)
	fmt.Println(ret)
	for i := 0; i < n; i++ {
		ret = append(ret, nums[i], nums[n+i])
	}
	return ret

}

func main() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))
}
