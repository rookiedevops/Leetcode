package addTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	First  []int
	Second []int
	Want   []int
}

func TestAddTwoNumbers(t *testing.T) {
	c := []Case{
		Case{
			First:  []int{2, 4, 3},
			Second: []int{5, 6, 4},
			Want:   []int{7, 0, 8},
		},
		Case{
			First:  []int{2, 4, 3},
			Second: []int{8, 5, 6},
			Want:   []int{0, 0, 0, 1},
		},
	}
	for _, i := range c {
		node := addTwoNumbers(produceListNode(i.First), produceListNode(i.Second))
		actual := produceNumbers(node)
		assert.Equal(t, i.Want, actual)
	}
}

func produceListNode(numbers []int) *ListNode {
	tail := &ListNode{}
	pre := tail

	for _, i := range numbers {
		node := &ListNode{
			Val: i,
		}
		pre.Next = node
		pre = node
	}
	return tail.Next
}

func produceNumbers(n *ListNode) []int {
	numbers := make([]int, 0)
	for {
		if n == nil {
			break
		}
		numbers = append(numbers, n.Val)
		n = n.Next
	}
	return numbers
}
