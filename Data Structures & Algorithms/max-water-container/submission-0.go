func maxArea(heights []int) int {
	mx := -1

	left, right := 0, len(heights)-1

	for left < right {
		curMax := (right - left) * min(heights[left], heights[right])

		if mx < curMax {
			mx = curMax
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return mx
}