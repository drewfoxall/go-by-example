package main 

import "fmt"

func main() {
	
	i := 1
	for i <= 10 {
		i = i + 1
		if i % 2 == 0{
			fmt.Println(i , "Is Even")
		} else {
			fmt.Println(i , "Is Odd")
		}
	}

}