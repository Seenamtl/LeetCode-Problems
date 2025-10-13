package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array to binary-search it
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	lo, hi := 0, m
	half := (m + n + 1) / 2

	for lo <= hi {
		i := (lo + hi) / 2 // cut in nums1
		j := half - i      // cut in nums2

		l1 := math.Inf(-1)
		if i > 0 {
			l1 = float64(nums1[i-1])
		}
		r1 := math.Inf(1)
		if i < m {
			r1 = float64(nums1[i])
		}

		l2 := math.Inf(-1)
		if j > 0 {
			l2 = float64(nums2[j-1])
		}
		r2 := math.Inf(1)
		if j < n {
			r2 = float64(nums2[j])
		}

		// Found correct partition
		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 0 {
				leftMax := math.Max(l1, l2)
				rightMin := math.Min(r1, r2)
				return (leftMax + rightMin) / 2.0
			}
			return math.Max(l1, l2)
		}

		// Move partition
		if l1 > r2 {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}

	// Should never reach here if inputs are valid
	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{}))
}
