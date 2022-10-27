package main

import "fmt"

var RomanNumerals = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0
	biggerNum := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := RomanNumerals[rune(s[i])]

		if num >= biggerNum {
			biggerNum = num
			result += num
			continue
		}
		result -= num
	}
	return result
}

func main() {
	fmt.Println(romanToInt("XVII"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("MCM"))
}
