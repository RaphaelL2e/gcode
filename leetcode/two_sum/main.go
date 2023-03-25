package main

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i := range nums {
        if j,ok :=m[target-nums[i]];ok {
            return []int{j,i}
        }else{
            m[nums[i]] = i
        }
    }
    return nil
}

func twoSum2(nums []int, target int) []int {
	for i:=0;i<len(nums);i++ {
		for j :=i+1;j<len(nums);j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return nil
}
