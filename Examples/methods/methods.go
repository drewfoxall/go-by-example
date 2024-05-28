package main

import "fmt"

type receiver struct{
	width,height int
}

func (r *receiver) area() int{
	return r.width * r.height
}

func main() {
	r := receiver{width: 10, height: 5}

	fmt.Println("Area: ", r.area())

}