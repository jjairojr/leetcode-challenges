package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	for i, value := range matrix {
		for j := range value {
			if i != 0 && j != 0 && matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
}
