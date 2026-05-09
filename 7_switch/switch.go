package main

import "fmt"

func main() {
	// simple switch
	/*
		i := 9
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		case 6:
			fmt.Println("six")
		case 7:
			fmt.Println("seven")
		default:
			fmt.Println("others")

		}
	*/

	/*// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	default:
		fmt.Println("It's workday!")
	}*/

	/*// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer!")
		case string:
			fmt.Println("It's a string!")
		case bool:
			fmt.Println("It's a boolean!")
		case float32:
			fmt.Println("It's a float32")
		case float64:
			fmt.Println("It's a float64")
		default:
			fmt.Println("others", t)

		}
	}*/

	// type switch option
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer!")
		case string:
			fmt.Println("It's a string!")
		case bool:
			fmt.Println("It's a boolean!")
		case float32:
			fmt.Println("It's a float32")
		case float64:
			fmt.Println("It's a float64")
		default:
			fmt.Println("others")

		}
	}

	whoAmI("golang")
	whoAmI(4.9885676)
	whoAmI(5.8989489608908999)
	whoAmI(false)

}
