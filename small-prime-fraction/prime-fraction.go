package main

import (
	"fmt"
	"sort"
)

func main() {
	var list []int
	var len, data, key int
	fmt.Println("enter length of array")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("enter prime number to array:")
		fmt.Scan(&data)
		list = append(list, data)
	}
	fmt.Println("enter key value")
	fmt.Scan(&key)
	fmt.Println(kthSmallestPrimeFraction(list, key))
}
func kthSmallestPrimeFraction(arr []int, k int) []int {
	var data = make(map[float64][]int)
	//var answer []int
	var array []float64
	//var count int
	if len(arr) == 2 {
		return []int{arr[0], arr[1]}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			temp1 := float64(arr[i])
			temp2 := float64(arr[j])
			data[temp1/temp2] = []int{arr[i], arr[j]}
			array = append(array, temp1/temp2)
		}
	}
	// for i := 0; i < len(array); i++ {
	// 	for j := 0; j < len(array); j++ {
	// 		if array[i] < array[j] {
	// 			ele := array[i]
	// 			array[i] = array[j]
	// 			array[j] = ele
	// 		}
	// 	}
	// }
	sort.Float64s(array)
	return data[array[k-1]]
}
