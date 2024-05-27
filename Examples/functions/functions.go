package main

import "fmt"

func plus(a int, b int) int{
	return a + b
}

func intSeq () func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

func plusPlus(a, b, c int) int{
	return a + b + c
}
func plusFirstElement(a, b, c int) (int, int){
	return a + b, a + c
}

func sum(nums ...int){
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 1)
	fmt.Println(res)

	res = plusPlus(1, 1, 1)
	fmt.Println(res)

	first, second :=  plusFirstElement(1, 1, 1)
	fmt.Println(first)
	fmt.Println(second)

	sum(1,2,3,4,5,6,7,8)

	nextInt := intSeq()
	
	fmt.Println(nextInt())

}
