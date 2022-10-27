package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	var i, j = 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i += 1
		}
		j += 1
	}
	return i == len(s)
}

func main() {
	fmt.Println(isSubsequence("acb", "aaaabchgdbcccc"))
}
