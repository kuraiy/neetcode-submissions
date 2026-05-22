func twoSum(nums []int, target int) []int {
    mapNums := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if v, ok := mapNums[target-nums[i]]; ok {
			return []int{v, i}
		}

		if _, ok := mapNums[nums[i]]; !ok {
			mapNums[nums[i]] = i
		}
	}
	return []int{}
}
