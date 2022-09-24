package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func main() {
// 	head1 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  5,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	head2 := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val: 3,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val: 5,
// 						Next: &ListNode{
// 							Val:  6,
// 							Next: nil,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	fmt.Println(middleNode(head1).Val)
// 	fmt.Println(middleNode(head2).Val)
// }

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	fast := head
	slow := head
	counter := 1

	for {
		if fast.Next == nil {
			break
		}
		if fast.Next.Next == nil {
			counter += 1
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
		counter += 2
	}

	if counter%2 == 0 {
		return slow.Next
	}
	return slow
}
