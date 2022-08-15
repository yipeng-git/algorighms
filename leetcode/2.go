package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	var tail, head *ListNode
//	var carry int
//	for l1 != nil || l2 != nil {
//		tmp := carry
//		if l1 != nil {
//			tmp += l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			tmp += l2.Val
//			l2 = l2.Next
//		}
//		carry = tmp / 10
//		if head == nil {
//			head = &ListNode{Val: tmp % 10}
//			tail = head
//		} else {
//			tail.Next = &ListNode{Val: tmp % 10}
//			tail = tail.Next
//		}
//	}
//	if carry > 0 {
//		tail.Next = &ListNode{
//			Val: 1,
//		}
//	}
//	return head
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		tmp := carry
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		if tail == nil {
			tail = &ListNode{Val: tmp % 10}
		} else {
			tail.Next = &ListNode{Val: tmp % 10}
			tail = tail.Next
		}
		carry = tmp / 10
		if head == nil {
			head = tail
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}
