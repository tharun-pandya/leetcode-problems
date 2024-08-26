package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter string")
	fmt.Scan(&str)
	fmt.Println(minimumLength(str))
}
func minimumLength(s string) int {
	right := len(s) - 1
	left := 0
	//l_ele := 0
	//r_ele := 0
	for left = 0; left <= right; left++ {
		if left == right {
			return 0
		} else if s[left] == s[right] {
			if s[left] == s[right-1] {
				right -= 2
			} else if s[left+1] == s[right] {
				right--
			}
		} else {
			left++
			break
		}
	}
	result := left - right
	return result
}
