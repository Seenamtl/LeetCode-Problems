package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	res := 0
	for x > 0 {
		pop := x % 10
		x /= 10

		// check for overflow before multiplying
		if res > (math.MaxInt32-pop)/10 {
			return 0
		}

		res = res*10 + pop
	}

	return res * sign
}

func main() {
	tests := []int{123, -123, 120, 0, 1534236469}
	for _, t := range tests {
		fmt.Printf("reverse(%d) = %d\n", t, reverse(t))
	}
}
