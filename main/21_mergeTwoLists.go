package main

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
// 	printList(mergeTwoLists(head1, head2))
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var head *ListNode
	var current *ListNode
	if list1.Val < list2.Val {
		head, current = list1, list1
		list1 = list1.Next
	} else {
		head, current = list2, list2
		list2 = list2.Next
	}

	for list1 != nil || list2 != nil {
		if list1 == nil {
			current.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			current.Next = list1
			list1 = list1.Next
		} else if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	return head
}
