package main

import "fmt"

/**
思路：
使用一个栈stack来模拟操作
1. 将pushed数组中的元素依次入栈
2. 同时判断入栈的这个元素是不是popped要出栈的数
3. 如果是，将这个数从popped中出栈
4. 最后判断栈是否为空
*/

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	//循环遍历压入栈
	for _, v := range pushed {
		//入栈元素与popped栈顶元素是否相等，如果不相等，则入栈
		if v != popped[0] {
			stack = append(stack, v)
		} else {
			popped = popped[1:]
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == popped[0] {
					popped = popped[1:]
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
	}
	return len(stack) == 0
}

func validateStackSequences1(pushed []int, popped []int) bool {
	stack := []int{}

	k := 0 //标识popped移动标记
	//循环遍历压入栈
	for _, v := range pushed {
		stack = append(stack, v)

		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i] == popped[k] {
				stack = stack[:len(stack)-1]
				k++
			} else {
				break
			}
		}
	}
	return len(stack) == 0
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}

	fmt.Println(validateStackSequences(pushed, popped))
}
