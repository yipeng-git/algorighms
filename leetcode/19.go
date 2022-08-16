package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	reversed := reverse(head)
	return reverseAndPop(reversed, n)
}

func reverse(curr *ListNode) *ListNode {
	if curr == nil {
		return nil
	}
	head := curr
	var tail *ListNode
	for head != nil {
		next := head
		head = head.Next
		next.Next = tail
		tail = next
	}
	return tail
}

func reverseAndPop(curr *ListNode, n int) *ListNode {
	if curr == nil {
		return nil
	}
	head := curr
	cnt := 1
	var tail *ListNode
	for head != nil {
		next := head
		head = head.Next
		if cnt != n {
			next.Next = tail
			tail = next
		}
		cnt += 1
	}
	return tail
}
