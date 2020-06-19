package main

import (
	"fmt"
)

/**
题目：
给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
返回一对观光景点能取得的最高分。

示例：
输入：[8,1,5,2,6]
输出：11
解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
*/

/**
1. A[i] + A[j] + i - j ,可以拆分成A[i]+i和A[j]-j两部分，求出两部分的最大值
*/

func maxScoreSightseeingPair(A []int) int {
	max := 0
	leftMax := A[0] + 0

	for i := 1; i < len(A); i++ {
		temp := leftMax + A[i] - i
		if max < temp {
			max = temp
		}

		leftTemp := A[i] + i
		if leftMax < leftTemp {
			leftMax = leftTemp
		}

	}

	return max
}

func maxScoreSightseeingPair1(A []int) int {
	res, cur := 0, 0

	for _, v := range A {
		res = maxA(res, cur+v)
		cur = maxA(cur, v) - 1
	}
	return res
}

func maxA(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{8, 1, 5, 2, 6}

	ret := maxScoreSightseeingPair1(nums)
	fmt.Println(ret)
}
