func dailyTemperatures(temperatures []int) []int {
	monoStack := make([]int, 0)
	result := make([]int, len(temperatures))
	for i, v := range temperatures {
		for {
			n := len(monoStack) - 1
			if n == -1 {
				monoStack = append(monoStack, i)
				break
			}
			if v > temperatures[monoStack[n]] {
				result[monoStack[n]] = i - monoStack[n]
				monoStack = monoStack[:n]
			} else {
				monoStack = append(monoStack, i)
				break
			}
		}
	}

	return result
}