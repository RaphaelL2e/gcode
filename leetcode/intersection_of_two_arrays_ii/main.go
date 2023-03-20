package intersection_of_two_arrays_ii

// by myself
func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 10)
	for _, i := range nums1 {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 1
		}
	}
	n := make(map[int]int, 10)
	for _, j := range nums2 {
		if _, ok := n[j]; ok {
			n[j] += 1
		} else {
			n[j] = 1
		}
	}
	var res []int
	for k, v := range m {
		if v != 0 {
			if v > n[k] {
				v = n[k]
			}
			for i := 0; i < v; i++ {
				res = append(res, k)
			}
		}
	}
	return res
}

// 优化
func intersect2(nums1 []int, nums2 []int) []int {
	// 对较小的数组hash
	var tmp []int
	if len(nums1) > len(nums2) {
		tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}
	m := make(map[int]int)
	for _, i := range nums1 {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 1
		}
	}
	var res []int
	for _, j := range nums2 {
		// 特殊 m[j] > 0 才是有效计数
		if _, ok := m[j]; ok && m[j] > 0 {
			m[j] -= 1
			res = append(res, j)
		}
	}
	return res
}
