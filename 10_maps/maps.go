package main

import (
	"fmt"
	"maps"
)

// associative data structure -> hash, object, dict
func main() {
	fmt.Println("_________ creating map _______________")

	m := make(map[string]string)

	fmt.Println("_________________________ setting an element __________________")

	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println("_________________ get an element________________")

	// important: key doesn't exit in map. it gives zero valued
	fmt.Println(m["name"], m["area"], m["phone"])
	fmt.Println(len(m))

	delete(m, "name")

	fmt.Println(m)

	clear(m)

	fmt.Println(m)

	m["nikename"] = "go"
	m["circle_area"] = "backend"

	n := map[string]int{"age": 27, "ID_no": 453, "phone": 4, "price": 45553}
	fmt.Println(n)

	v, age := n["age"]

	fmt.Println(v)
	if age {
		fmt.Println("It's exist!")
	} else {
		fmt.Println("It's not exist!")
	}

	n1 := map[string]int{"age": 27, "ID_no": 453, "phone": 4, "price": 45553}
	n2 := map[string]int{"age": 27, "ID_no": 453, "phone": 4, "price": 45553}

	fmt.Println(maps.Equal(n1, n2))
}
