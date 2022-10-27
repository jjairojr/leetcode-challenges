package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	res := 0

	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[string(s[r])]; ok {
			l = max(l, index)
		}
		res = max(res, r-l+1)
		m[string(s[r])] = r + 1
	}
	return res
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

}
