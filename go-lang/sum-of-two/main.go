package main

import (
	"fmt"
)

func main() {
	var nums []int
	var size, target, data int
	fmt.Println("Enter how many values you want to enter to add")
	fmt.Scan(&size)
	for a := 0; a < size; a++ {
		fmt.Println("Enter value into array")
		fmt.Scan(&data)
		nums = append(nums, data)
	}
	fmt.Println("Enter Target value")
	fmt.Scan(&target)
	TwoSum(nums, target)
}
func TwoSum(nums []int, target int) []int {
	var result []int
	for a := 0; a <= len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			if target-nums[b] == nums[a] {
				result = append(result, a, b)
				return result
			}

		}
	}
	return nil
}
