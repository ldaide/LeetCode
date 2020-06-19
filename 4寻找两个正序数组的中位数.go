package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num []int

	i := 0
	j := 0
	for i < len(nums1) || j < len(nums2) {
		fmt.Println(i, ",", j)
		if nums1[i] < nums2[j] {
			num = append(num, nums1[i])
			i++
		} else {
			num = append(num, nums2[j])
			j++
		}
		fmt.Println(num)
	}

	l := len(num)
	if l%2 == 1 {
		return float64(num[l%2])
	} else {
		mid := l / 2
		return float64((num[mid] + num[mid+1]) / 2)
	}

}

func main() {
	num1 := []int{1, 3}
	num2 := []int{2}

	fmt.Println(findMedianSortedArrays(num1, num2))
}
