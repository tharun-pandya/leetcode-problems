package main

// pending 3075 problem

import "fmt"

func main() {
	var array []int
	var len, k, data int
	fmt.Println("enter length of the array")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("enter value to the array")
		fmt.Scan(&data)
		array = append(array, data)
	}
	fmt.Println("enter number of childerns")
	fmt.Scan(&k)
	fmt.Println(maximumHappinessSum(array, k))
}
func maximumHappinessSum(happiness []int, k int) int64 {
	var result int64
	var counti, countj int
	for i := 0; i < len(happiness); i++ {
		for j := 0; j < len(happiness); j++ {
			if happiness[i] < happiness[j] {
				temp := happiness[i]
				happiness[i] = happiness[j]
				happiness[j] = temp
			}
		}
	}
	if k == 1 {
		return int64(happiness[len(happiness)-1])
	}
	for i := len(happiness) - 1; i >= 0; i-- {
		result += int64(happiness[i])
		for j := i; j >= 0; j-- {
			if happiness[j] == 0 {
				continue
			}
			happiness[j] = happiness[j] - 1
			countj++
			if countj == k {
				countj = 0
				break
			}
		}
		counti++
		if counti == k {
			counti = 0
			break
		}
	}
	return result
}
