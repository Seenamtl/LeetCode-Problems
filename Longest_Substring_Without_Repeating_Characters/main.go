package main

import "fmt"

// lengthOfLongestSubstring returns the length of the longest substring
// without repeating characters.
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int) // stores last seen index of each character
	maxLen := 0
	start := 0 // start index of current window

	for i := 0; i < len(s); i++ {
		if lastIdx, found := charIndex[s[i]]; found && lastIdx >= start {
			// Move start right past the previous occurrence
			start = lastIdx + 1
		}
		charIndex[s[i]] = i
		// Update max length
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
