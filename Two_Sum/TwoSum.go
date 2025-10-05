package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// Create a map to store the value -> index mapping
	m := make(map[int]int)

	// Loop through the array
	for i, num := range nums {
		complement := target - num
		// Check if the complement exists in the map
		if idx, found := m[complement]; found {
			return []int{idx, i}
		}
		// Otherwise, store the current number and its index
		m[num] = i
	}
	// Since the problem guarantees a solution, we don't need to handle "no result"
	return nil
}

func main() {
	// Example 1
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(nums1, target1)) // Output: [0, 1]

	// Example 2
	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println(twoSum(nums2, target2)) // Output: [1, 2]

	// Example 3
	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(nums3, target3)) // Output: [0, 1]
}
