package main

import (
	"fmt"
)

func main() {
	var order, s string
	fmt.Println("Enter order")
	fmt.Scan(&order)
	fmt.Println("enter string")
	fmt.Scan(&s)
	fmt.Println(customSortString(order, s))
}
func customSortString(order string, s string) string {
	var result string
	for i := 0; i < len(order); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == order[i] {
				result = result + string(order[i])
			}
		}

	}
	temp := result
	for _, i := range s {
		count := 0
		for _, j := range temp {
			if i == j {
				count++
			}
		}
		if count == 0 {
			result = result + string(i)
		}
	}
	return result
}
