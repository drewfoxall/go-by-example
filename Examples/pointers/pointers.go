package main

import "fmt"

func zeroVal(ival int){
	ival = 0
}

func zeroPtr(iptr *int){
	*iptr = 0
}


func main() {
	i := 1
	fmt.Println("Initial: " , i)
	zeroVal(i)
	fmt.Println("zeroVal: ", i)
	zeroPtr(&i)
	fmt.Println("Pointerval: ", i)
	fmt.Println("Pointerval: ", &i)
}