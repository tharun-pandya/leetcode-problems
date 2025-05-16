package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type head struct {
	start *ListNode
}

func main() {
	var l1, l2, new *ListNode
	var s1, s2 head
	var data, in int
	fmt.Println("Enter how much data you want enter for list1")
	fmt.Scan(&in)
	for i := 0; i < in; i++ {
		fmt.Println("Enter data for list1 of", i, "index")
		fmt.Scan(&data)
		l1 = &ListNode{Val: data, Next: nil}
		if i == 0 {
			s1.start = l1
			new = l1
		} else {
			new.Next = l1
			new = l1
		}
	}
	fmt.Println("Enter how much data you want enter for list2")
	fmt.Scan(&in)
	for i := 0; i < in; i++ {
		fmt.Println("Enter data for list2 of", i, "index")
		fmt.Scan(&data)
		l2 = &ListNode{Val: data, Next: nil}
		if i == 0 {
			s2.start = l2
			new = l2
		} else {
			new.Next = l2
			new = l2
		}
	}
	l1 = s1.start
	for l1 != nil {
		fmt.Print(l1.Val, "->")
		l1 = l1.Next
	}
	fmt.Println()
	l2 = s2.start
	for l2 != nil {
		fmt.Print(l2.Val, "->")
		l2 = l2.Next
	}
	fmt.Println()
	l1 = s1.start
	l2 = s2.start
	fmt.Println(addTwoNumbers(l1, l2))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr1, arr2 []int
	pw, j, ans, sum := 1, 0, 0, 0
	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}
	if len(arr2) > len(arr1) {
		count := len(arr1)
		for i := 0; i < len(arr2); i++ {
			for j <= i {
				pw *= 10
				j++
			}
			for k := len(arr1) - 1; k > +0; k-- {

				if count == k {
					sum = arr1[k] + arr2[i]
					sum *= pw
					ans += sum
				}
			}
			count--
			if count == 0 {
				sum = arr2[i] + 0
				sum *= pw
				ans += sum
			}
			fmt.Println(sum)
			fmt.Println(ans)
			fmt.Println("")
		}
	} else {
		for i := len(arr1) - 1; i > +0; i-- {
			sum := arr1[i] + arr2[i]
			fmt.Println(sum)
			sum *= pw
			ans += sum
			fmt.Println(ans)
		}
	}
	ans = ans / 10
	fmt.Println(ans)

	//final linked list
	var final, new *ListNode
	var h head
	var dup int
	count1 := 1
	dup = ans
	for {
		ans /= 10
		if ans == 0 {
			break
		}
		count1++
	}
	ans = dup
	for i := 0; i < count1; i++ {
		dup = ans
		ans %= 10
		final = &ListNode{Val: ans, Next: nil}
		if i == 0 {
			new = final
			h.start = final
		} else {
			new.Next = final
			new = final
		}
		ans = dup / 10
	}
	final = h.start
	return final
}
