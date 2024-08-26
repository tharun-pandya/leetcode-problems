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
	fmt.Println(getCommon(array1, array2))
}
func getCommon(nums1 []int, nums2 []int) int {
	var i = 1
	if nums1[0] == nums2[0] {
		return nums1[0]
	}
	if nums1[len(nums1)-1] < nums2[0] || nums2[len(nums2)-1] < nums1[0] {
		return -1
	} else if nums1[0] < nums2[0] {
		for {
			if nums1[i] == nums2[0] {
				return nums2[0]
			}
			if nums1[i] > nums2[0] {
				break
			}
			i++
		}
		for j := i - 1; j < len(nums1); j++ {
			for k := 0; k < len(nums2); k++ {
				if nums1[j] == nums2[k] {
					return nums2[k]
				}
			}
		}
	} else if nums2[0] < nums1[0] {
		for {
			if nums2[i] == nums1[0] {
				return nums1[0]
			}
			if nums2[i] > nums1[0] {
				break
			}
			i++
		}
		for j := i - 1; j < len(nums2); j++ {
			for k := 0; k < len(nums1); k++ {
				if nums2[j] == nums1[k] {
					return nums1[k]
				}
			}
		}
	}
	fmt.Println("nothing is execute")
	return -1
}
