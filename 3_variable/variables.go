package main

import "fmt"

func main() {
	// variable declaration and values assign explicit type
	var name string = "Abu Bakar"
	fmt.Println(name)

	var price int = 300
	fmt.Println(price)

	var isActive bool = true
	fmt.Println(isActive)

	var height float32 = 5.9
	fmt.Println(height)

	// variable declare without value assign
	var college_name string

	// value assign later
	college_name = "Master Nazir Ahamed College."

	fmt.Println(college_name)

	// values short hand or inplicit ~ inference types
	studentId := "st34445"
	fmt.Println(studentId)

	isSaved := false
	fmt.Println(isSaved)
}
