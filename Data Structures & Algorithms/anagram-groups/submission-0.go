func groupAnagrams(strs []string) [][]string {
	type StringWithIndex struct {
		Str string
		Ind int
	}
	sortedStrs := make([]StringWithIndex, 0)

	for i, v := range strs {
		runes := []rune(v)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedStrs = append(sortedStrs, StringWithIndex{string(runes), i})
	}

	sort.Slice(sortedStrs, func(i, j int) bool {
		return sortedStrs[i].Str < sortedStrs[j].Str
	})

	result := make([][]string, 0)
	ind := 0
	result = append(result, []string{strs[sortedStrs[ind].Ind]})

	for i := 1; i < len(sortedStrs); i++ {
		if sortedStrs[i].Str == sortedStrs[i-1].Str {
			result[ind] = append(result[ind], strs[sortedStrs[i].Ind])
		} else {
			ind++
			result = append(result, []string{strs[sortedStrs[i].Ind]})
		}
	}
	return result
}
