package utils

// source: https://stackoverflow.com/a/73039056
func CartesianProduct[T any](matrix [][]T) [][]T {
	nextIndex := func(ix []int, lens func(i int) int) {
		for j := len(ix) - 1; j >= 0; j-- {
			ix[j]++

			if j == 0 || ix[j] < lens(j) {
				return
			}

			ix[j] = 0
		}
	}

	lens := func(i int) int { return len(matrix[i]) }

	results := make([][]T, 0, len(matrix))
	for indexes := make([]int, len(matrix)); indexes[0] < lens(0); nextIndex(indexes, lens) {
		var temp []T

		for j, k := range indexes {
			temp = append(temp, matrix[j][k])
		}

		results = append(results, temp)
	}

	return results
}
