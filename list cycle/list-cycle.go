package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head, new *ListNode
	var data, count int
	fmt.Println("length")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&data)
		head = &ListNode{Val: data, Next: new}
		new = head
	}
	fmt.Println(hasCycle(head))
}
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
