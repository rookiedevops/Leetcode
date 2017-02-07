package medianOfTwoSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)

	if (total & 1) == 1 {
		return findKtd(nums1, nums2, total/2+1)
	} else {
		return (findKtd(nums1, nums2, total/2) + findKtd(nums1, nums2, total/2+1)) / 2
	}
}

func findKtd(nums1, nums2 []int, k int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		return findKtd(nums2, nums1, k)
	}
	if m == 0 {
		return float64(nums2[k-1])
	}
	if k == 1 {
		return float64(min(nums1[0], nums2[0]))
	}
	pa := min(k/2, m)
	pb := k - pa

	if nums1[pa-1] < nums2[pb-1] {
		return findKtd(nums1[pa:], nums2, k-pa)
	} else if nums1[pa-1] > nums2[pb-1] {
		return findKtd(nums1, nums2[pb:], k-pb)
	}
	return float64(nums1[pa-1])
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
