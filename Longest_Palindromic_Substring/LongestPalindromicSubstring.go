package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0

	expand := func(l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		// when loop exits, palindrome is s[l+1:r]
		return l + 1, r - 1
	}

	for i := 0; i < len(s); i++ {
		// Odd-length center at i
		l1, r1 := expand(i, i)
		// Even-length center between i and i+1
		l2, r2 := expand(i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
}
