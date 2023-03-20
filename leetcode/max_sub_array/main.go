package max_sub_array

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if sum < 0 {
			sum = 0
			// 负数情况，容易忽略
			if max < num {
				max = num
			}
		} else {
			if max < sum {
				max = sum
			}
		}
	}
	return max
}
