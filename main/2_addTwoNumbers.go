package main

// func main() {
// 	head1 := &ListNode{
// 		Val: 2,
// 		Next: &ListNode{
// 			Val: 4,
// 			Next: &ListNode{
// 				Val: 3,
// 			},
// 		},
// 	}
// 	head2 := &ListNode{
// 		Val: 5,
// 		Next: &ListNode{
// 			Val: 6,
// 			Next: &ListNode{
// 				Val: 4,
// 			},
// 		},
// 	}
// 	fmt.Println(addTwoNumbers(head1, head2))
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sumTot := carry
		if l1 != nil {
			sumTot += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sumTot += l2.Val
			l2 = l2.Next
		}

		sumMod := 0
		if sumTot >= 10 {
			sumMod = sumTot - 10
			carry = 1
		} else {
			sumMod = sumTot
			carry = 0
		}
		current.Next = &ListNode{Val: sumMod}
		current = current.Next
	}

	return dummyHead.Next
}
