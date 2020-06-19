package main

import (
	"fmt"
	"sort"
)

func main() {

	//nums1 := []int{1, 2, 3, 0, 0, 0}
	//m := 3
	//nums2 := []int{2, 5, 6}
	//n := 3

	//nums1 := []int{1, 2, 3, 0, 0, 0, 0, 0}
	//m := 3
	//nums2 := []int{-2, -1, 2, 5, 6}
	//n := 5

	//nums1 := []int{2, 0}
	//m := 1
	//nums2 := []int{1}
	//n := 1

	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1

	merge2(nums1, m, nums2, n)

}

//合并后排序
func merge(nums1 []int, m int, nums2 []int, n int) {

	//for i := 0; i < n; i++ {
	//	nums1[m+i] = nums2[i]
	//}

	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)

	fmt.Println(nums1)
}

//双指针 / 从前往后
func merge1(nums1 []int, m int, nums2 []int, n int) {
	num1_copy := make([]int, m)
	copy(num1_copy, nums1[:m])
	p1 := 0
	p2 := 0
	p := 0

	for p1 < m && p2 < n {
		if num1_copy[p1] < nums2[p2] {
			nums1[p] = num1_copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
		p++
	}

	if p1 < m {
		nums1 = append(nums1[:p], num1_copy[p1:m]...)
	}
	if p2 < n {
		nums1 = append(nums1[:p], nums2[p2:n]...)
	}
	fmt.Println(nums1)
}

//双指针 / 从后往前

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 && m > 0 {
		p1 := m - 1
		p2 := n - 1
		p := m + n - 1

		for p1 >= 0 && p2 >= 0 {
			if nums1[p1] < nums2[p2] {
				nums1[p] = nums2[p2]
				p2--
			} else {
				nums1[p] = nums1[p1]
				p1--
			}
			p--
		}

		nums1 = append(nums2[:p2+1], nums1[p2+1:]...)

	} else if m == 0 && n > 0 {
		copy(nums1, nums2)
	}
	fmt.Println(nums1)
}
