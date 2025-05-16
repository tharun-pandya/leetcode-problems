package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head, new, start *ListNode
	var count, data int
	fmt.Println("Enter length")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Println("enter data for", i, "in list")
		fmt.Scan(&data)
		head = &ListNode{Val: data, Next: nil}
		if i == 0 {
			start = head
		} else {
			new.Next = head
		}
		new = head
	}
	head = start
	for i := 0; i < count; i++ {
		fmt.Print(head, "->\t")
		head = head.Next
	}
	fmt.Println("")
	res := middleNode(start)
	for res != nil {
		fmt.Print(res, "->\t")
		res = res.Next
	}
}
func middleNode(head *ListNode) *ListNode {
	count := 0
	start := head
	for head != nil {
		count++
		head = head.Next
	}
	head = start
	count = int(count / 2)
	for i := 0; i < count; i++ {
		head = head.Next
	}
	return head
}
