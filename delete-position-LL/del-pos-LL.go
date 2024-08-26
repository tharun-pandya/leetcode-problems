package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var list, new *ListNode
	var len, key, data int
	fmt.Println("Enter length of linked list")
	fmt.Scan(&len)
	if len == 0 {
		return
	}
	fmt.Println("Enter data to the list", 0)
	fmt.Scan(&data)
	list = &ListNode{Val: data, Next: nil}
	new = list
	head := list
	for i := 1; i < len; i++ {
		fmt.Println("Enter data to the list", i)
		fmt.Scan(&data)
		list = &ListNode{Val: data, Next: nil}
		new.Next = list
		new = list
	}
	start := head
	for i := 0; i < len; i++ {
		fmt.Print(start.Val, "->")
		fmt.Println(i)
		start = start.Next
	}
	start = head
	fmt.Println("enter a index value to delete")
	fmt.Scan(&key)
	result := (removeNthFromEnd(head, key))
	for i := 0; i < len-1; i++ {
		fmt.Print(result.Val, "->")
		result = result.Next
	}
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	var ptr *ListNode = head
	var prev, nxt *ListNode
	for ptr != nil {
		count++
		ptr = ptr.Next
	}
	if count == 1 || count == n {
		head = head.Next
		ptr = nil
	} else {
		count = count - n
		ptr = head
		for i := 0; i < count; i++ {
			prev = ptr
			ptr = ptr.Next
		}
		nxt = ptr.Next
		prev.Next = nxt
		ptr = nil
		ptr = nxt
	}
	return head
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
// count := 0
// var nxt, ptr, prev *ListNode
// head = reverse(head, nxt, prev)
// if n == 1 {
// 	head = head.Next
// 	prev = nil
// 	head = reverse(head, nxt, prev)
// 	return head
// } else {
// 	ptr = head
// 	for ptr != nil {
// 		count++
// 		nxt = ptr.Next
// 		if count == n {
// 			prev.Next = nxt
// 			ptr = nil
// 			break
// 		}
// 		prev = ptr
// 		ptr = ptr.Next
// 	}
// 	head = reverse(head, nxt, prev)
// }
// 	return head
// }
func reverse(crr, nxt, prev *ListNode) *ListNode {
	for crr != nil {
		nxt = crr.Next
		crr.Next = prev
		prev = crr
		crr = nxt
	}
	return prev
}
