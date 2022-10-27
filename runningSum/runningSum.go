package main

import "fmt"

func runningSum(nums []int) []int {
	previousSum := 0
	arrayAux := []int{}

	for i := 0; i < len(nums); i++ {
		previousSum += nums[i]
		arrayAux = append(arrayAux, previousSum)
	}
	return arrayAux
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
