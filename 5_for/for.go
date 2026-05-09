package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop

	/*
		i := 1
		for i <= 6 {
			fmt.Println(i)
			i++
		}*/
	// infinite loop
	/*for {
		fmt.Println(1)
	}*/

	// classic for loop
	/*for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}*/

	// 1.22 version new feature: range loop
	for i := range 11 {
		fmt.Println(i)
	}

}
