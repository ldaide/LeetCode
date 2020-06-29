package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
题目：
给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。

示例 1：
输入：nums = [12,345,2,6,7896]
输出：2
解释：
12 是 2 位数字（位数为偶数）
345 是 3 位数字（位数为奇数）
2 是 1 位数字（位数为奇数）
6 是 1 位数字 位数为奇数）
7896 是 4 位数字（位数为偶数）
因此只有 12 和 7896 是位数为偶数的数字
*/

//先把数字转化为字符串，然后统计字符串长度
func findNumbers(nums []int) int {
	n := 0
	for _, v := range nums {
		s := strconv.Itoa(v) //耗时4ms
		//s := fmt.Sprintf("%d", v) //耗时8ms
		if len(s)%2 == 0 {
			n++
		}
	}
	return n
}

//求10的对数
//math.Log10(10.2) = 1.xxx
//math.Log10(112.2) = 2.xxx
// 10^(k-1)< x < 10^k ==> k = log10()
func findNumbers1(nums []int) int {
	n := 0
	for _, v := range nums {
		l := int(math.Log10(float64(v))) + 1
		if l%2 == 0 {
			n++
		}
	}
	return n
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers1(nums))

	fmt.Println(math.Log10(1024))
}
