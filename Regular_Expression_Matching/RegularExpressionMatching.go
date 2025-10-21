package main

import "fmt"

// isMatch returns true if pattern p matches the entire string s.
// Pattern supports '.' (any single char) and '*' (zero or more of prev char).
func isMatch(s, p string) bool {
	memo := make(map[[2]int]bool)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if v, ok := memo[[2]int{i, j}]; ok {
			return v
		}

		// If pattern consumed, match only if string also consumed
		if j == len(p) {
			return i == len(s)
		}

		// First character matches?
		first := i < len(s) && (p[j] == s[i] || p[j] == '.')

		var ans bool
		// If next pattern char is '*', we can:
		// 1) skip "x*" (zero occurrence), or
		// 2) if first matches, consume one char from s and stay on same p
		if j+1 < len(p) && p[j+1] == '*' {
			ans = dp(i, j+2) || (first && dp(i+1, j))
		} else {
			ans = first && dp(i+1, j+1)
		}

		memo[[2]int{i, j}] = ans
		return ans
	}

	return dp(0, 0)
}

func main() {
	fmt.Println(isMatch("aa", "a"))                   // false
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("ab", ".*"))                  // true
	fmt.Println(isMatch("aab", "c*a*b"))              // true
	fmt.Println(isMatch("mississippi", ".*s?"))       // false (pattern unsupported '?', just example)
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
}
