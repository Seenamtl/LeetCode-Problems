package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		x /= 10

		// overflow check (optional, since LeetCode guarantees 32-bit input)
		if reversed > (1<<31-1-digit)/10 {
			return false
		}
		reversed = reversed*10 + digit
	}

	return original == reversed
}

func main() {
	tests := []int{121, -121, 10, 0, 12321, 1234321, 12345}
	for _, t := range tests {
		fmt.Printf("isPalindrome(%d) = %v\n", t, isPalindrome(t))
	}
}

/*func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return x == reversed || x == reversed/10
}
*/
