package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter divident")
	fmt.Scan(&num1)
	fmt.Println("Enter divisior")
	fmt.Scan(&num2)
	fmt.Println(divide(num1, num2))
}

// func divide(dividend int, divisor int) int {
// 	var result int = divisor
// 	var count int
// 	var negetive bool
// 	// testcase := 2
// 	// for i := 1; i < 31; i++ {
// 	// 	testcase *= 2
// 	// }
// 	if dividend < 0 && divisor < 0 {
// 		divisor = -divisor
// 		result = -result
// 		dividend = -dividend
// 	} else if dividend < 0 {
// 		dividend = -dividend
// 		negetive = true
// 	} else if divisor < 0 {
// 		divisor = -divisor
// 		result = -result
// 		negetive = true
// 	}
// 	for {
// 		// if testcase == count {
// 		// 	if negetive == true {
// 		// 		count = testcase
// 		// 	} else {
// 		// 		count = testcase - 1
// 		// 	}
// 		// 	break
// 		// }
// 		if dividend < result || count == 2147483648 {
// 			break
// 		}
// 		result += divisor
// 		count++
// 	}
// 	if count == 2147483648 {
// 		if negetive == true {
// 			return -count
// 		} else {
// 			return count - 1
// 		}
// 	}
// 	if negetive == true {
// 		return -count
// 	}
// 	return count
// }
func divide(dividend int, divisor int) int {
	var result int
	var count int
	if dividend < 0 && divisor < 0 {
		for {
			if result <= dividend || count == 2147483647 {
				return count
			}
			result += divisor
			count++
		}
	} else if dividend > 0 && divisor > 0 {
		//result = divisor
		for {
			if result >= dividend || count == 2147483647 {
				return count
			}
			result += divisor
			count++
		}
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		result = divisor
		for {
			if result >= dividend || count == 2147483648 {
				return -count
			}
			result += divisor
			count++
		}
	} else if dividend > 0 && divisor < 0 {
		divisor = -divisor
		result = divisor
		for {
			if result >= dividend || count == 2147483648 {
				return -count
			}
			result += divisor
			count++
		}
	}
	return 0
}
