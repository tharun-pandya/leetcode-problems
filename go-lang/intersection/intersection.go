package main

import "fmt"

func main() {
	var array1, array2 []int
	var len1, len2, data int
	fmt.Println("Enter length of first array")
	fmt.Scan(&len1)
	for i := 0; i < len1; i++ {
		fmt.Println("Enter data for first array", i, "position")
		fmt.Scan(&data)
		array1 = append(array1, data)
	}
	fmt.Println("Enter length of second array")
	fmt.Scan(&len2)
	for i := 0; i < len2; i++ {
		fmt.Println("Enter data for second array", i, "position")
		fmt.Scan(&data)
		array2 = append(array2, data)
	}
	fmt.Println(array1, array2)
	fmt.Println(intersection(array1, array2))
}
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for k := 0; k < len(result); k++ {
					if nums1[i] == result[k] {
						count++
						break
					}
				}
				if count == 0 {
					result = append(result, nums1[i])
				}
				count = 0
			}
		}
	}
	return result
}
