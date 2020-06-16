package main

import "fmt"

/**
1. for 循环移动非0元素，并统计非0元素个数
2. for 循环补齐0元素
 */
func moveZeroes(nums []int) {
	j := 0
	for i,val := range nums {
		if val != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}


func moveZeroes1(nums []int) {
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
			if i != j {
				nums[i] = 0
			}
		}
	}

	fmt.Println(nums)
}

/**
for循环遍历，非零元素进行交换，j只有非0时才加1，所以j小于等于i
[0 1 0 3 12]
i=0,j=0,nums[0]=0
i=1,j=0,nums[1] = 1,交换：[1 0 0 3 12]，j=1
i=2,j=1,nums[2] = 0
i=3,j=1,nums[3] = 3,交换：[1 3 0 0 12]，j=2
i=4,j=2,nums[4] = 12,交换：[1 3 12 0 0]
 */
func moveZeroes2(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	fmt.Println(nums)
}


/**

 */
func moveZeroes3(nums []int) {

	//每遇到一个0，进行一次减1操作，最终记录非0元素个数
	l := len(nums)
	i := 0

	for {
		//当i大于l时，后面的元素都是0，即可终止循环
		if i > l {
			break
		}
		if nums[i] == 0 {
			//把0前后两部分合并，然后再末尾添加一个0
			//l 减1
			nums = append(nums[0:i], nums[i+1:]...)
			nums = append(nums, 0)
			l--
		} else {
			i++
		}
	}

	fmt.Println(nums)
}

func main() {
	nums := [][]int{
		{0, 1, 0, 3, 12},
		{0,1},
		{},
		{0},
	}
	for _,val :=range nums {
		fmt.Println(val)
		moveZeroes(val)
		fmt.Println("-----------")
	}


}
