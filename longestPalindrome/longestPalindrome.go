package main

import "fmt"

func longestPalindrome(s string) int {
	result := 0
	m := countLetters(s)
	for _, v := range m {
		result += v / 2 * 2
		if result%2 == 0 && v%2 == 1 {
			result += 1
		}
	}
	return result
}

func countLetters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, value := range s {
		m[value]++
	}
	return m
}

func main() {
	fmt.Println(longestPalindrome("a"))
}
