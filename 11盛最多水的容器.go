package main

import (
	"fmt"
)

//枚举，双层循环,逐个计算，找出最大值
func maxArea(height []int) int {
	maxArea := 0

	for i := 0; i < len(height)-1; i++ {
		for j := 1; j < len(height); j++ {
			minHeight := min(height[i], height[j])
			area := (j - i) * minHeight
			maxArea = max(maxArea, area)
		}
	}
	return maxArea
}

func maxArea1(height []int) int {
	maxArea := 0

	i := 0
	j := len(height) - 1

	for i < j {
		minHeight := 0
		if height[i] < height[j] {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}
		area := (j - i + 1) * minHeight
		maxArea = max(maxArea, area)
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea1(height))
}
