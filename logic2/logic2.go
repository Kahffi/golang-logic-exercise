package logic2

func generateIntGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)

	for idx := range grid {
		grid[idx] = make([]int, cols)
	}

	return grid

}

func Soal1(n int) [][]int {
	result := generateIntGrid(n, n)

	for _, row := range result {
		for colIdx := range row {
			row[colIdx] = colIdx*2 + 1
		}
	}

	return result
}

func Soal2(n int) [][]int {
	result := generateIntGrid(n, n)

	for _, row := range result {
		for colIdx := range row {
			row[colIdx] = (colIdx + 1) * 2
		}
	}

	return result

}

func Soal3(n int) [][]int {
	result := generateIntGrid(n, n)

	for idx, row := range result {
		for colIdx := range row {
			row[colIdx] = (idx*n+colIdx)*2 + 1
		}
	}

	return result

}

func Soal4(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, row := range result {
		for colIdx := range row {

			row[colIdx] = currentVal
			if rowIdx == 0 || colIdx == 0 || colIdx == n-1 {

				currentVal += 3
			} else {
				currentVal += 2
			}
		}
	}

	return result
}

func Soal5(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1
	direction := 1

	for _, row := range result {
		for colIdx := range row {
			if direction > 0 {
				row[colIdx] = currentVal
			} else {
				row[len(row)-colIdx-1] = currentVal
			}
			currentVal += 2
		}
		direction *= -1
	}

	return result

}

func Soal6a(n int) [][]int {
	result := generateIntGrid(n, n)

	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
				currentVal += 3
			} else {
				rows[n-1-colIdx] = currentVal
				currentVal += 2

			}
		}
	}

	return result
}

func Soal6b(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
				currentVal += 3
			} else {
				rows[n-1-colIdx] = currentVal
				currentVal += 2

			}
		}
	}

	return result
}
func Soal6c(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx%2 == 0 {
				rows[colIdx] = currentVal
				currentVal += 3
			} else {
				rows[colIdx] = currentVal
				currentVal -= 3

			}

		}

	}

	return result
}

func Soal7(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx == colIdx {
				rows[colIdx] = currentVal
				currentVal += 2
			}

		}

	}

	return result

}

func Soal8(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx == colIdx {
				result[len(result)-rowIdx-1][colIdx] = currentVal
				currentVal += 2
			}

		}

	}

	return result

}

func Soal9(n int) [][]int {
	result := generateIntGrid(n, n)
	currentVal := 1

	for rowIdx, rows := range result {
		for colIdx := range rows {
			if rowIdx == colIdx {
				rows[colIdx] = currentVal
				currentVal += 2
			}

		}

	}

	return result

}
