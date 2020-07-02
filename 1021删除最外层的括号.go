package main

import "fmt"

/**
思路：
1. 循环遍历字符
2. 如果是'(',则入栈
3. 如果是')'，判断栈顶元素是否是')',如果是，弹出元素，如果不是，将')'压入栈
4. 如果栈为空时，获取s+1到i的元素，并更新s=i+1
*/
func removeOuterParentheses(S string) string {

	stack := make([]byte, 0)
	ret := ""
	start := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, S[i])
		} else {
			e := stack[len(stack)-1]
			if e == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, S[i])
			}
		}

		if len(stack) == 0 {
			ret += S[start+1 : i]
			start = i + 1
		}
	}
	return ret
}

func main() {
	s := "()()"
	fmt.Println(removeOuterParentheses(s))
}

func removeOuterParentheses1(S string) string {

	stack := NewStack()
	ret := ""
	start := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack.Push(S[i])
		} else {
			if stack.GetTopE() == '(' {
				stack.Pop()
			} else {
				stack.Push(S[i])
			}
		}

		if stack.IsEmpty() {
			ret += S[start+1 : i]
			start = i + 1
		}
	}
	return ret
}

type Stack struct {
	data []byte
	size int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]byte, 0),
		size: 0,
	}
}

func (this *Stack) GetTopE() byte {
	return this.data[this.size-1]
}

func (this *Stack) IsEmpty() bool {
	return this.size == 0
}

func (this *Stack) Push(e byte) {
	this.data = append(this.data, e)
	this.size++
}

func (this *Stack) Pop() byte {
	this.size--
	e := this.data[this.size]
	this.data = this.data[:this.size]
	return e
}
