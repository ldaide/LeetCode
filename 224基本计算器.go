package main

import "fmt"

/**
题目：
实现一个基本的计算器来计算一个简单的字符串表达式的值。
字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:
输入: "1 + 1"
输出: 2

示例 2:
输入: " 2-1 + 2 "
输出: 3

示例 3:
输入: "(1+(14+5+2)-3)+(6+8)"
输出: 23

*/

func calculate(s string) int {

	stack := make([]int, 10)

	//加减法标识，1表示加号，-1表示减号
	sign := 1
	//结果
	res := 0
	//当前数字
	curNum := 0
	//遍历字符串
	for i := 0; i < len(s); i++ {
		//数字
		if s[i] >= '0' && s[i] <= '9' {
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				curNum = 10*curNum + int(s[i]-'0')
			}
			res += curNum * sign
			curNum = 0
			i--
		} else {
			switch s[i] {
			case '+':
				sign = 1
			case '-':
				sign = -1
			case '(':
				//遇到'('时，保存之前的结果，重新计算括号里的表达式
				//遇到'('时将结果和sign压入栈，重置res和sign
				stack = append(stack, res, sign)
				res = 0
				sign = 1
			case ')':
				//把当前计算的值与 之前保存的结果 再次进行计算
				stackLen := len(stack)
				sign = stack[stackLen-1]
				temp := stack[stackLen-2]
				stack = stack[:stackLen-2]
				res = temp + sign*res
			}
		}
	}

	return res
}

func main() {
	//s := "(1+(4+5+2)-3)+(6+8)"
	s := "30-(9-1)"
	ret := calculate(s)

	fmt.Println(ret)
}
