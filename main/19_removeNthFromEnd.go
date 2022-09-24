package main

// func main() {
// 	head := &ListNode{
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
// 		},
// 	}
//
// 	newHead := removeNthFromEnd(head, 2)
// 	fmt.Println(newHead)
// 	newHead2 := removeNthFromEnd(head2, 2)
// 	printList(newHead2)
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := getListSize(head)
	if size == 0 {
		return nil
	}
	// invalid state
	if n > size {
		return nil
	}
	// one element in list, one to remove -> return nil
	if size == 1 && n == 1 {
		return nil
	}
	// to remove first element
	if n == size {
		return head.Next
	}
	toRemoveIndex := size - n

	connectorHead := head
	for i := 0; i < toRemoveIndex-1; i++ {
		connectorHead = connectorHead.Next
	}
	connectorHead.Next = connectorHead.Next.Next

	return head
}

func getListSize(head *ListNode) int {
	size := 0
	for head != nil {
		size++
		head = head.Next
	}
	return size
}
