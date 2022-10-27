package main

import "fmt"

func pivotIndex(nums []int) int {
	arraySum := sumArray(nums)
	leftSum := 0

	for i := 0; i < len(nums); i++ {
		if leftSum == (arraySum - leftSum - nums[i]) {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

func sumArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
