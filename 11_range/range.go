package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{4, 5, 6, 7, 8}

	// __________________option 1___________
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// __________________option 2___________

	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// }

	// m := map[string]int{"price": 45, "age": 27, "phones": 3}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// __________________option 3___________

	// for v := range m {
	// 	fmt.Println(v)
	// }

	// __________________option 4___________ index and charater unicode code point rune
	// 1. (i is below for loop) starting byte of rune

	for i, c := range "I am from Bangladesh" {
		fmt.Println(i, "-----", c, "______", string(c))
	}
}
