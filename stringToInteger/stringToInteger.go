package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	start := false
	flag := 1
	res := 0
	for _, v := range s {
		if v == ' ' && !start {
			continue
		}
		if v == '-' && !start {
			flag = -1
			start = true
			continue
		}
		if v == '+' && !start {
			start = true
			continue
		}

		if v >= 48 && v <= 57 {
			start = true
			digit := int(v-'0') * flag
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > 7) {
				return math.MaxInt32
			}
			if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < -8) {
				return math.MinInt32
			}
			res = res*10 + digit
		} else {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(("42"))
	fmt.Println(myAtoi("        -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("abc 4193 with words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("-+-+1"))
	fmt.Println(myAtoi("21474836460"))
	fmt.Println(myAtoi("0000-42a1234"))
	fmt.Println(myAtoi("3.14159"))
}
