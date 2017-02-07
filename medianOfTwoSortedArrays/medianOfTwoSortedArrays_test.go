package medianOfTwoSortedArrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	Nums1 []int
	Nums2 []int
	Want  float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	c := []Case{
		Case{
			Nums1: []int{1, 3},
			Nums2: []int{2},
			Want:  float64(2),
		},
		Case{
			Nums1: []int{1, 3},
			Nums2: []int{2, 4},
			Want:  (float64(2) + float64(3)) / 2,
		},
		Case{
			Nums1: []int{6},
			Nums2: []int{1, 2, 3, 4, 5},
			Want:  (float64(3) + float64(4)) / 2,
		},
	}
	for _, i := range c {
		w := findMedianSortedArrays(i.Nums1, i.Nums2)
		assert.Equal(t, i.Want, w)
	}
}
