package main

import (
	"fmt"
)

func main() {
	var array []int
	var len, data, key int
	fmt.Println("Enter how many elements you want to enter in array")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("Enter data into array at position ", i)
		fmt.Scan(&data)
		array = append(array, data)
	}
	fmt.Println("Enter key value ")
	fmt.Scan(&key)
	fmt.Println(divideArray(array, key))
}
func divideArray(nums []int, k int) [][]int {
	var array [][]int
	var temp []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp

			}
		}
	}
	fmt.Println(nums)
	index := 0
	for i := 0; i < len(nums); i++ {
		if len(nums) == i+1 {
			break
		} else {
			if nums[i+1]-nums[i] <= k {
				if index == 0 {
					temp = append(temp, nums[i], nums[i+1])
					index += 2
				} else if index == 3 {
					array = append(array, []int{temp[0], temp[1], temp[2]})
					index = 0
					temp = append(temp[0:0])
				} else if index != 0 {
					temp = append(temp, nums[i+1])
					index++
				}
			}
		}
	}
	return array
}
