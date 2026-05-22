func topKFrequent(nums []int, k int) []int {
	type Counter struct {
		Count int
		Value int
	}

	numIndexes := make(map[int]int)
	result := make([]Counter, 0)

	for _, v := range nums {
		if i, ok := numIndexes[v]; ok {
			result[i].Count++
		} else {
			result = append(result, Counter{1, v})
			numIndexes[v] = len(result) - 1
		}

	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	ans := make([]int, 0)
	for _, v := range result[:k] {
		ans = append(ans, v.Value)
	}

	return ans
}
