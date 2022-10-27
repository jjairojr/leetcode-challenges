package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for pos1, value := range nums {
		for pos2, sv := range nums {
			if value+sv == target && pos2 != pos1 {
				result = []int{pos2, pos1}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{1, 4, 8, 0, 20, 1, 7}, 27))
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}
