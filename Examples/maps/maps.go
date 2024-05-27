package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("Map: ", m)

	age := m["age"]
	fmt.Println(age)

	fmt.Println(len(m))

	delete(m, "addressNumber")
	fmt.Println("Map: ", m)

	clear(m)

	n:= map[string]int{"for": 1, "bar": 2}
	fmt.Println(n)

}