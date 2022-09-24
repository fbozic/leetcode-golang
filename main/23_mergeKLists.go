package main

import (
	"math"
)

// func main() {
// 	head1 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 4,
// 			},
// 		},
// 	}
// 	head2 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 3,
// 			Next: &ListNode{
// 				Val: 4,
// 			},
// 		},
// 	}
// 	printList(mergeKLists([]*ListNode{head1, nil, nil, head2, nil, nil, nil}))
// }

func mergeKLists(lists []*ListNode) *ListNode {
	head, smallestIndex := findSmallestListNode(lists)
	if head == nil {
		return nil
	}
	current := head
	lists[smallestIndex] = head.Next

	for {
		smallest, smallestIndex := findSmallestListNode(lists)
		if smallest == nil {
			break
		}

		current.Next = smallest
		current = smallest
		lists[smallestIndex] = smallest.Next
	}
	return head
}

func findSmallestListNode(lists []*ListNode) (*ListNode, int) {
	smallestVal := math.MaxInt
	var smallest *ListNode
	var smallestIndex int
	for i, head := range lists {
		if head == nil {
			continue
		}
		if head.Val < smallestVal {
			smallest = head
			smallestIndex = i
			smallestVal = head.Val
		}
	}

	return smallest, smallestIndex
}
