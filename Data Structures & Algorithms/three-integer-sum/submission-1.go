func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			if nums[left]+nums[right] == -nums[i] {
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}

			if nums[left]+nums[right] < -nums[i] {
				left++
				for {
					if left < right && nums[left-1] == nums[left] {
						left++
					} else {
						break
					}
				}
			} else {
				right--
				for {
					if left < right && nums[right+1] == nums[right] {
						right--
					} else {
						break
					}
				}
			}
		}
	}

	return result
}
