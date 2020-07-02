package main

type Node struct {
	val  int
	pre  *Node
	next *Node
}

type MyCircularDeque struct {
	head *Node //头节点
	tail *Node //尾节点
	len  int
	cap  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor2(k int) MyCircularDeque {
	head := Node{
		val:  -1,
		pre:  nil,
		next: nil,
	}
	tail := Node{
		val:  -1,
		pre:  nil,
		next: nil,
	}
	head.next = &tail
	tail.pre = &head
	return MyCircularDeque{
		head: &head,
		tail: &tail,
		len:  0,
		cap:  k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	node := &Node{
		val:  value,
		pre:  this.head,
		next: this.head.next,
	}
	this.head.next.pre = node
	this.head.next = node

	this.len++

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	node := &Node{
		val:  value,
		pre:  this.tail.pre,
		next: this.tail,
	}
	this.tail.pre.next = node
	this.tail.pre = node
	this.len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	temp := this.head.next
	temp.next.pre = this.head
	this.head.next = temp.next
	temp.pre, temp.next = nil, nil
	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	temp := this.tail.pre

	temp.pre.next = this.tail
	this.tail.pre = temp.pre
	temp.pre, temp.next = nil, nil
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.head.next.val
	}

	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.tail.pre.val
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len >= this.cap
}

func main() {
	circularDeque := Constructor2(3)
	circularDeque.InsertLast(1)  // 返回 true
	circularDeque.InsertLast(2)  // 返回 true
	circularDeque.InsertFront(3) // 返回 true
	circularDeque.InsertFront(4) // 已经满了，返回 false
	circularDeque.GetRear()      // 返回 2
	circularDeque.IsFull()       // 返回 true
	circularDeque.DeleteLast()   // 返回 true
	circularDeque.InsertFront(4) // 返回 true
	circularDeque.GetFront()
}
