package main

import (
	"fmt"
)

func main() {
	var array []string
	var data string
	var len int
	fmt.Println("enter how many data elements")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("enter word")
		fmt.Scan(&data)
		array = append(array, data)
	}
	for i := 0; i < len; i++ {
		fmt.Println(array[i])
		for _, j := range array[i] {
			fmt.Print(string(j))
		}
		fmt.Println()
	}
	fmt.Println(longestCommonPrefix(array))
}
func longestCommonPrefix(strs []string) string {
	var result, temp string
	for _, j := range strs[0] {
		for _, k := range strs[1] {
			if j == k {
				count := 0
				for _, i := range temp {
					if i == j {
						count++
					}
				}
				if count == 0 {
					temp = temp + string(j)
				}
			}
		}
	}
	if temp == "" {
		return ""
	} else {
		for _, i := range strs[1:] {
			for _, j := range i {
				for _, k := range temp {
					if k == j {
						var res string
						count := 0
						for _, l := range result {
							if l == j {
								count++
							}
						}
						if count == 0 {
							res = res + string(j)
						}
						result = result + res
					}
				}
			}
			if result == "" {
				return ""
			}
		}
		return result
	}
}
