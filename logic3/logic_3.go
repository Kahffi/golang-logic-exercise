package logic3

func generateIntGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)

	for idx := range grid {
		grid[idx] = make([]int, cols)
	}

	return grid

}

func Soal1(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if colIdx > rowIdx {
				continue
			}
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
			} else {
				rows[(colIdx-rowIdx)*-1] = currentVal
			}
			currentVal += 2
		}
	}

	return result
}
func Soal2(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if colIdx < rowIdx {
				continue
			}
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
			} else {
				rows[n-colIdx-1+rowIdx] = currentVal
			}
			currentVal += 2
		}
	}

	return result
}
func Soal3(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := -1
	// addition := 3

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if colIdx > n-1-rowIdx {
				continue
			}

			if rowIdx%2 == 0 {
				currentVal += 2
				rows[colIdx] = currentVal
			} else {
				currentVal += 3
				rows[n-1-colIdx-rowIdx] = currentVal
			}
		}
	}

	return result
}

func Soal4(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if n-1-rowIdx < colIdx {
				continue
			}
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
			} else {
				rows[n-colIdx-rowIdx-1] = currentVal
			}
			currentVal += 2
		}
	}

	return result
}

func Soal5(n int) [][]int {
	result := generateIntGrid(n, n)
	val := 1
	mid := (n - 1) / 2

	for rowIdx := 0; rowIdx <= mid; rowIdx++ {
		for colIdx := 0; colIdx <= mid; colIdx++ {
			if !(rowIdx >= colIdx && rowIdx+colIdx <= n-1) {
				continue
			}
			if rowIdx%2 == 0 {
				result[rowIdx][colIdx] = val
				result[rowIdx][n-1-colIdx] = val
				result[n-1-rowIdx][n-1-colIdx] = val
				result[n-1-rowIdx][colIdx] = val
			} else {
				result[rowIdx][rowIdx-colIdx] = val
				result[rowIdx][(n - 1 + colIdx - rowIdx)] = val
				result[n-1-rowIdx][(n - 1 + colIdx - rowIdx)] = val
				result[n-1-rowIdx][rowIdx-colIdx] = val
			}
			val += 2
		}
	}

	return result
}

func Soal7(n int) [][]int {
	result := generateIntGrid(n, n)
	val := 1

	for rowIdx, row := range result {
		if rowIdx > (n-1)/2 {
			break
		}
		for colIdx := range row {
			if !(colIdx+rowIdx >= (n-1)/2 && colIdx-rowIdx <= (n-1)/2) {
				continue
			}
			if rowIdx%2 != 0 {
				result[rowIdx][colIdx] = val
				result[n-1-rowIdx][colIdx] = val
			} else {
				result[rowIdx][n-1-colIdx] = val
				result[n-1-rowIdx][n-1-colIdx] = val

			}
			val += 2
		}
	}

	return result
}
func Soal8(n int) [][]int {
	result := generateIntGrid(n, n)
	val := 1
	mid := (n - 1) / 2

	for rowIdx, row := range result {
		if rowIdx > mid {
			break
		}
		for colIdx := range row {
			if !(colIdx+rowIdx >= mid && colIdx-rowIdx <= mid) {
				continue
			}
			if rowIdx%2 != 0 {
				result[colIdx][rowIdx] = val
				result[colIdx][n-1-rowIdx] = val
			} else {
				result[n-1-colIdx][rowIdx] = val
				result[n-1-colIdx][n-1-rowIdx] = val

			}
			val += 2
		}
	}

	return result
}

func Soal10(n int) [][]int {
	result := generateIntGrid(n, n)
	mid := (n - 1) / 2

	for rowIdx, row := range result {
		if rowIdx > mid {
			break
		}
		val := 2*rowIdx + 1
		for colIdx := range row {
			if colIdx > mid {
				break
			}
			if colIdx <= mid-rowIdx {
				result[mid-rowIdx][mid-colIdx] = val
				result[mid+rowIdx][mid+colIdx] = val
				result[mid+rowIdx][mid-colIdx] = val
				result[mid-rowIdx][mid+colIdx] = val
				val += 2
			}
		}

	}
	return result

}
