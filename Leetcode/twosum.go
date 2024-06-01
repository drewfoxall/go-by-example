package main
import "fmt"

func main() {
	fmt.Println(twoSum([]int { 2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int { 3, 2, 4}, 6))
	fmt.Println(twoSum([]int {3, 3}, 9))

}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i< n; i++ {
		for y := i + 1; y < n; y++ {
			if nums[y] + nums[i] == target{
				return []int{i, y}
			}

		}
	}




	return []int{}
}