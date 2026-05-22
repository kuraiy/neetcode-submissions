func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	counter := make(map[int]int)
	mx := 1

	for _, v := range nums {
		if _, ok := counter[v]; !ok {
			counter[v] = 1
			if leftSum, ok := counter[v-1]; ok {
				counter[v] += leftSum
				if mx < counter[v] {
					mx = counter[v]
				}
			}

			i := v + 1
			for {
				if _, ok := counter[i]; !ok {
					if i-1 != v {
						counter[i-1] += counter[v]
						if mx < counter[i-1] {
							mx = counter[i-1]
						}
					}
					break
				}
				i++
			}
		}
	}
	return mx
}