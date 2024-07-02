package main

import (
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int

	nums3 = append(nums3, nums1...)
	nums3 = append(nums3, nums2...)

	slices.Sort(nums3)

	i := len(nums3)

	if i%2 == 0 {
		i = i / 2
		return float64((float64(nums3[i]) + float64(nums3[i-1])) / 2)
	} else {
		return float64(nums3[(i / 2)])
	}
}
