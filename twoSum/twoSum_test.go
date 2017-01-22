package twoSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Slice  []int
	Target int
	Want   []int
}

func TestTwoSum(t *testing.T) {
	c := []Case{
		Case{
			Slice:  []int{-1, 0, 1, 2, 3},
			Target: -1,
			Want:   []int{0, 1},
		},
		Case{
			Slice:  []int{-1, 0, 1, 2, 3},
			Target: 1,
			Want:   []int{1, 2},
		},
		Case{
			Slice:  []int{-1, 0, 1, 2, 3},
			Target: 5,
			Want:   []int{3, 4},
		},
	}
	for _, i := range c {
		o := twoSum(i.Slice, i.Target)
		assert.Equal(t, o, i.Want)
	}
}
