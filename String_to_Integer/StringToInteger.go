package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	n := len(s)
	i := 0

	// 1) skip leading spaces
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	// 2) optional sign
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 3) digits
	res := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		d := int(s[i] - '0')

		// 4) overflow check BEFORE multiplying by 10
		if sign == 1 {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && d > 7) {
				return math.MaxInt32
			}
		} else {
			// for negative, last digit can be 8
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && d > 8) {
				return math.MinInt32
			}
		}

		res = res*10 + d
		i++
	}

	return sign * res
}

func main() {
	tests := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"+1",
		"00000-42a1234",
		"21474836460",
		"",
		"   ",
	}
	for _, t := range tests {
		fmt.Printf("myAtoi(%q) = %d\n", t, myAtoi(t))
	}
}
