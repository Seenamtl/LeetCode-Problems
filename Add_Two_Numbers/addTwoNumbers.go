package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented by linked lists.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy node to simplify list handling
	current := dummy
	carry := 0

	// Traverse both lists until both are nil and carry is 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, n := range nums {
		current.Next = &ListNode{Val: n}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to print a linked list
func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print(" -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	// Example 1: (2 -> 4 -> 3) + (5 -> 6 -> 4) = 7 -> 0 -> 8
	l1 := createList([]int{2, 4, 3})
	l2 := createList([]int{5, 6, 4})
	result := addTwoNumbers(l1, l2)
	printList(result) // Output: 7 -> 0 -> 8

	// Example 2: (0) + (0) = 0
	l3 := createList([]int{0})
	l4 := createList([]int{0})
	printList(addTwoNumbers(l3, l4)) // Output: 0

	// Example 3: (9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9) + (9 -> 9 -> 9 -> 9) = 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
	l5 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	l6 := createList([]int{9, 9, 9, 9})
	printList(addTwoNumbers(l5, l6)) // Output: 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1
}
