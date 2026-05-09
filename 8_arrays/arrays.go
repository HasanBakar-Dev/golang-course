package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums)
	fmt.Println(len(nums))

	fmt.Println("__________________option 2: 2d array______________________")

	// Correct ways to initialize and append to a 2D slice:

	// Method 1: Initialize with empty slices, then append a 2D slice
	nums2 := [][]int{{}, {}}
	// Correct way to append a 2D slice - wrap it in braces
	nums2 = append(nums2, []int{1, 2, 3})
	nums2 = append(nums2, []int{4, 5, 6})

	fmt.Println(nums2)

	// Method 2: Initialize directly with 2D slices
	nums3 := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums3)

	// Method 3: If you want to append multiple rows at once
	nums4 := [][]int{{}, {}}
	nums4 = append(nums4, []int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println(nums4)
}
