package main

import "fmt"

func main() {
	var len, data int
	var array = make([]int, len)
	fmt.Println("enter length of array")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("enter data to array(", i, ")")
		fmt.Scan(&data)
		array = append(array, data)
	}
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if array[i] < array[j] {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
	fmt.Println(sortedSquares(array))
}
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}

// ********----best way---**********
// func sortedSquares(nums []int) []int {
//     squares := make([]int, len(nums))
//     l, r := 0, len(nums) - 1
//     for i := len(nums) - 1; i >= 0; i-- {
//         ls, rs  := nums[l] * nums[l], nums[r] * nums[r]
//         if ls >= rs {
//             squares[i] = ls
//             l++
//             continue
//         }
//         squares[i] = rs
//         r--
//     }
//     return squares
// }
