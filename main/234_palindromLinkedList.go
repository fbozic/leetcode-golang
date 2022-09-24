package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	head := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 2,
// 				Next: &ListNode{
// 					Val:  1,
// 					Next: nil,
// 				},
// 			},
// 		},
// 	}
// 	head2 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 1,
// 			},
// 		},
// 	}
// 	headNo1 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 		},
// 	}
//
// 	fmt.Println(isPalindrome(head))
// 	fmt.Println(isPalindrome(head2))
// 	fmt.Println(isPalindrome(headNo1))
// }

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	revertedHead, _ := revertLinkList(head)
	return compareList(head, revertedHead)
}

func compareList(head1 *ListNode, head2 *ListNode) bool {
	if head1 == nil && head2 == nil {
		return true
	}
	if head1 == nil || head2 == nil {
		return false
	}
	current := head1.Val == head2.Val

	next := compareList(head1.Next, head2.Next)
	return current && next
}

// return head and current
func revertLinkList(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		revertedHead := &ListNode{
			Val: head.Val,
		}
		return revertedHead, revertedHead
	}

	nextVal := head.Val

	revertedHead, current := revertLinkList(head.Next)
	next := &ListNode{
		Val: nextVal,
	}
	current.Next = next

	return revertedHead, next
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printList(head.Next)
}
