package main

//break
import "fmt"

func main() {
	var array1, array2 []int
	var len1, len2, data int
	fmt.Println("enter length of the array")
	fmt.Scan(&len1)
	for i := 0; i < len1; i++ {
		fmt.Println("enter value to the array")
		fmt.Scan(&data)
		array1 = append(array1, data)
	}
	for i := 0; i < len2; i++ {
		fmt.Println("enter value to the array")
		fmt.Scan(&data)
		array2 = append(array2, data)
	}
	fmt.Println(findMedianSortedArrays(array1, array2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	j := 0
	for i := 0; i < len(nums1); i++ {
		for j < len(nums2) {
			if j == (len(nums2)-1) && nums1[i] < nums2[j] {
				nums1 = append(nums1, nums2...)
			} else if nums1[i] < nums2[j] {
				break
			} else if nums1[i] == nums2[j] {

			}
		}
	}
	return result
}
