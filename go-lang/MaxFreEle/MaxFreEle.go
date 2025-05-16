package main

import (
	"fmt"
	"sort"
)

func main() {
	var array []int
	var data, count int
	fmt.Println("Enter length of elements")
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Println("Enter data for array", i)
		fmt.Scan(&data)
		array = append(array, data)
	}
	fmt.Println(maxFrequencyElements(array))
}

func maxFrequencyElements(nums []int) int {
	var result int
	var count int
	var bucket map[int]int
	bucket = make(map[int]int)
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	bucket[nums[0]] = 1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		bucket[nums[i]] = count
		count = 0
	}
	for _, value1 := range bucket {
		for _, value2 := range bucket {
			if value2 >= value1 {
				result = value2
			}
		}
	}
	for key, value := range bucket {
		if value == result {
			delete(bucket, key)
			break
		}
	}
	res := result
	for _, value := range bucket {
		if value == result {
			res = value + res
		}
	}
	fmt.Println(bucket, nums)
	return res
}

// func maxFrequencyElements(nums []int) int {
//     mapCount := make(map[int]int)

//     for _, val := range nums {
//         mapCount[val]++
//     }

//     count := 0
//     max := -1
//     for _, freq := range mapCount {
//         if freq > max {
//             max = freq
//         }
//     }

//     for _, freq := range mapCount {
//         if freq == max {
//             count += max
//         }
//     }
//     return count
// }
