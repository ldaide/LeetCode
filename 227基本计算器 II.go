package main

import "fmt"

/**
给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:
输入: "3+2*2"
输出: 7
*/

/**
乘法、除法优先计算，然后把结果压入栈
减法，转化成负数，然后压入栈
加法，直接压入栈
最后把栈里所有数字相加
*/
func calculate1(s string) int {

	stack := make([]int, 0)
	strlen := len(s)
	num := 0
	var op byte = '+'
	for i := 0; i < strlen; i++ {
		//数字
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}

		if (s[i] != ' ' && (s[i] < '0' || s[i] > '9')) || i == strlen-1 {
			switch op {
			case '-':
				num = -num
			case '*':
				stacklen := len(stack)
				last := stack[stacklen-1]
				stack = stack[:stacklen-1]
				num = last * num
			case '/':
				stacklen := len(stack)
				last := stack[stacklen-1]
				stack = stack[:stacklen-1]
				num = last / num
			}
			stack = append(stack, num)
			num = 0
			op = s[i]
		}
	}

	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}

func main() {
	//s := "(1+(4+5+2)-3)+(6+8)"
	s := "2048"
	ret := calculate1(s)

	fmt.Println(ret)
}
