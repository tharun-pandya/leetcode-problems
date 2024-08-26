package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	var result, s string
	s = fmt.Sprint(x)
	for _, i := range s {
		result = string(i) + result
	}
	num, _ := strconv.Atoi(result)
	if num == x {
		return true
	}
	return false
}
func main() {
	var num int
	fmt.Println("Enter a number to check that number Palindrome or not")
	fmt.Scan(&num)
	if isPalindrome(num) {
		fmt.Println(num, "is a Palindrome Number ")
	} else {
		fmt.Println(num, "is not Palindrome Number")
	}
}
