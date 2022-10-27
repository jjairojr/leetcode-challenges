package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = t[i]
		}
		if _, ok := m2[t[i]]; !ok {
			m2[t[i]] = s[i]
		}

		fmt.Println(m2[t[i]], s[i])

		if m[s[i]] != t[i] || m2[t[i]] != s[i] {
			return false
		}
	}

	return true
}

func main() {
	// fmt.Println(isIsomorphic("paper", "title")) // truer
	// fmt.Println(isIsomorphic("egg", "add"))     // true
	// fmt.Println(isIsomorphic("foo", "bar")) // false
	// fmt.Println(isIsomorphic("abc", "abc"))     // true
	// fmt.Println(isIsomorphic("badc", "baba"))   // false
	fmt.Println(isIsomorphic("egcd", "adfd")) // false
}
