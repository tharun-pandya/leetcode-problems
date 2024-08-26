package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter how many elements you want to enter in list1")

	fmt.Println("Enter how many elements you want to enter in list2")

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode
