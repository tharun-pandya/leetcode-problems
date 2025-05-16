package main

import "fmt"

func main() {
	var matrix [][]int
	var array []int
	var len, n, data int
	fmt.Println("Enter number of arrays")
	fmt.Scan(&len)
	fmt.Println("enter number of elements into arrays")
	fmt.Scan(&n)
	for i := 0; i < len; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("enter data to the array[", i, "][", j, "]")
			fmt.Scan(&data)
			array = append(array, data)
		}
		matrix = append(matrix, array)
	}
	rotate(matrix)
}
func rotate(matrix [][]int) {
	count := 0
	j := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		temp := matrix[count][j]
		matrix[count][j] = matrix[i][count]
		matrix[i][count] = temp

		if i == 0 {
			if count == len(matrix)-1 {
				break
			}
			count++
			j = 0
			i = len(matrix) - 1
		} else {
			j++
		}
	}
}
