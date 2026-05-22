func productExceptSelf(nums []int) []int {
	prefixMultiply := make([]int64, len(nums))
	suffixMultiply := make([]int64, len(nums))

	prefixMultiply[0] = int64(nums[0])
	suffixMultiply[len(nums)-1] = int64(nums[len(nums)-1])

	for i := 1; i < len(nums); i++ {
		prefixMultiply[i] = prefixMultiply[i-1] * int64(nums[i])
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffixMultiply[i] = suffixMultiply[i+1] * int64(nums[i])
	}

	result := make([]int, len(nums))

	result[0], result[len(nums)-1] = int(suffixMultiply[1]), int(prefixMultiply[len(nums)-2])

	for i := 1; i < len(nums)-1; i++ {
		result[i] = int(prefixMultiply[i-1] * suffixMultiply[i+1])
	}

	return result
}
