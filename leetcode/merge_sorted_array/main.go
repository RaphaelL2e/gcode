package merge_sorted_array

import "sort"

// 取末位大的先放
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if n == 0 {
			break
		}
		// 特殊情况 m == 0
		if m != 0 && nums1[(m-1)] > nums2[(n-1)] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}

// 先合并再排序
// 性能差
func merge2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
