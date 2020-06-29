package main

import "fmt"

/**
题目：
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
*/

func findKthLargest(nums []int, k int) int {
	data := insertSort(nums)
	fmt.Println(data)
	return data[k-1]
}

func sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var left, mid, right []int

	midTemp := nums[0]

	for _, v := range nums {
		if v < midTemp {
			left = append(left, v)
		} else if v > midTemp {
			right = append(right, v)
		} else {
			mid = append(mid, v)
		}
	}

	right = sort(right)
	left = sort(left)
	ret := append(right, mid...)
	ret = append(ret, left...)
	return ret
}

/**
快速排序
*/
func quitSort(nums []int) []int {

	return nums
}

/**
归并排序
*/
func mergeSort(nums []int) []int {

	return nums
}

/**
堆排序
*/
func heapSort(nums []int) []int {

	return nums
}

/**
希尔排序
*/
func shellSort(nums []int) []int {

	return nums
}

/**
插入排序
核心思想：将一个元素插入到 已排好序的数组中
将第i个元素与i之前的元素逐一比较，如果不符合，将后移一位，直到满足条件为止
*/
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp := nums[i]
			for j := i; j >= 0; j-- {
				if j > 0 && temp > nums[j-1] {
					nums[j] = nums[j-1]
				} else {
					nums[j] = temp
					break
				}

			}
		}
	}

	return nums
}

/**
选择排序
*/
func selectSort(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[min] < nums[j] {
				min = j
			}
		}
		if i != min {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}
	return nums
}

/**
冒泡排序
核心思想：不断的交换，通过交换完成最终的排序
1. 相邻两个元素进行比较
2. 把较大(小)元素交换到数组末尾
*/
func bubbleSort(nums []int) []int {
	l := len(nums)
	flag := true
	for i := 0; i < l && flag; i++ {
		flag = false
		for j := 0; j < l-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
	}
	return nums
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 4
	ret := findKthLargest(nums, k)
	fmt.Println(ret)
}
