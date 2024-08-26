package main

// after exprience

import "fmt"

func main() {
	var ele int
	fmt.Println("Enter number of nodes")
	fmt.Scan(&ele)
	fmt.Println("number of binary search trees formed with", ele, " nodes is", numTrees(ele))
}
func numTrees(n int) int {
	// quetion number 96
	var result int
	//var arr []int
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	//for i := 0; i < count; i++ {

	//}
	return result
}
