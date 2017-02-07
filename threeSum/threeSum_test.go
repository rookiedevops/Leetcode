package threeSum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Slice []int
	Want  [][]int
}

func TestThreeSum(t *testing.T) {
	c := []Case{
		Case{
			Slice: []int{-1, 0, 1, 2, -1, -4},
			Want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}
	for _, i := range c {
		o := threeSum(i.Slice)
		assert.Equal(t, o, i.Want)
	}
}
