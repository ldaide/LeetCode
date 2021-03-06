package main

import "fmt"

/**
题目：
给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。

示例：
输入：A = [4,5,0,-2,-3,1], K = 5
输出：7
解释：
有 7 个子数组满足其元素之和可被 K = 5 整除：
[4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]

*/

func subarraysDivByK(A []int, K int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range A {
		sum += elem
		modulus := (sum%K + K) % K
		record[modulus]++
	}
	fmt.Println(record)
	for _, cx := range record {
		ans += cx * (cx - 1) / 2
	}
	return ans

}

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5
	fmt.Println(-2 % 5)
	fmt.Println(subarraysDivByK(A, K))

}
