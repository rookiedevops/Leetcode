package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length-2; i++ {
		target := -nums[i]

		o := find(nums[i+1:], target)
		if len(o) > 0 {
			for _, ints := range o {
				var ok bool
				slice := append(ints, nums[i])
				sort.Ints(slice)
				for _, k := range res {
					if k[0] == slice[0] && k[1] == slice[1] && k[2] == slice[2] {
						ok = true
						break
					}
				}
				if !ok {
					res = append(res, slice)
				}
			}
		}
	}
	return res
}

func binarySearch(ints []int, target int) int {
	low := 0
	high := len(ints) - 1
	for {
		if low > high {
			break
		}
		middle := (low + high) / 2
		if ints[middle] == target {
			return middle
		} else if ints[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

func find(ints []int, target int) [][]int {
	res := make([][]int, 0)

	low := 0
	high := len(ints) - 1
	for {
		if low >= high {
			break
		}
		sum := ints[low] + ints[high]
		if sum == target {
			o := []int{ints[low], ints[high]}
			res = append(res, o)
			low++
			high--
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return res
}
