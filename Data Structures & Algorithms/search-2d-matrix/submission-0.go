func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix[0])
	m := len(matrix)
	l, r := 0, n*m-1

	for l <= r {
		mid := (l + r) / 2

		row := mid / n
		col := mid % n

		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}