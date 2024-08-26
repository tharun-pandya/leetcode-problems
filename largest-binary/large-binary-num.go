// to run .exe go file use "./file-name "
package main

import "fmt"

func main() {
	var num string
	fmt.Println("enter binary number")
	fmt.Scan(&num)
	fmt.Println(maximumOddBinaryNumber(num))
}
func maximumOddBinaryNumber(s string) string {
	count1, count0 := 0, 0
	var result string
	for _, i := range s {
		if i == '1' {
			count1++
		} else {
			count0++
		}
	}
	for i := 1; i <= count1; i++ {
		if count1 == i {
			for j := 0; j < count0; j++ {
				result += "0"
			}
			result += "1"
		} else {
			result += "1"
		}
	}
	return result
}
