package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l3 := &ListNode{}

	l1Next := l1
	l2Next := l2
	l3Next := l3

	carry := 0

	for {
		var val1, val2 int
		if l1Next == nil && l2Next == nil {
			break
		}
		if l1Next != nil {
			val1 = l1Next.Val
			l1Next = l1Next.Next
		}
		if l2Next != nil {
			val2 = l2Next.Val
			l2Next = l2Next.Next
		}

		sum := carry + val2 + val1
		node := &ListNode{
			Val: (sum % 10),
		}
		l3Next.Next = node
		l3Next = node

		carry = sum / 10
	}

	if carry != 0 {
		l3Next.Next = &ListNode{
			Val: carry,
		}
	}
	return l3.Next
}
