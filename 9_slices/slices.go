package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	// var names = make([]int, 5, 10)
	// initial we set size zero
	var names = make([]int, 0, 8)
	fmt.Println(names)
	fmt.Println(names == nil)
	names = append(names, 3, 4)
	fmt.Println(names)

	fmt.Println("____________________ option 3: how to size spread in go of slice_______________________")
	names = append(names, 3, 4, 4, 4, 8, 9, 3, 9, 8, 0, 4, 8, 4, 2, 4)
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))

	// _--- special slice declaration
	fmt.Println("____________________ option 4 _______________________")
	friends := []string{}
	fmt.Println(len(friends))
	fmt.Println(cap(friends))
	fmt.Println(friends)
	friends = append(friends, "fahim", "Karim", "eminy")
	fmt.Println(friends)
	fmt.Println("____________________ option 5: copy _______________________")

	friends2 := make([]string, len(friends))

	copy(friends2, friends)

	fmt.Println(friends2)

	fmt.Println("____________________ option 6: slice operator _______________________")

	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(nums2[0:5])
	fmt.Println(nums2[:4])
	fmt.Println(nums2[3:])

	fmt.Println("____________________ option 7: slices package _______________________")

	fmt.Println(slices.Equal(nums, nums2))

}
