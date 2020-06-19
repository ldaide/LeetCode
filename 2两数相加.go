package main

import "fmt"

/**
题目：
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Val = 0

	p, q, cur := l1, l2, dummyHead

	sum := 0

	for p != nil || q != nil {
		sum /= 10
		if p != nil {
			sum += p.Val
			p = p.Next
		}

		if q != nil {
			sum += q.Val
			q = q.Next
		}
		cur.Next = new(ListNode)
		cur.Next.Val = sum % 10

		cur = cur.Next
	}
	if sum/10 == 1 {
		cur.Next = new(ListNode)
		cur.Next.Val = 1
	}
	return dummyHead.Next
}

func main() {
	l1 := new(ListNode)

	l11 := &ListNode{0, nil}
	l12 := &ListNode{1, nil}

	l11.Next = l12

	ret := addTwoNumbers(l1, l11)

	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}

}
