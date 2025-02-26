package logic4

func sumOfNum(n int) int {
	val := 0
	for i := 1; i <= n; i++ {
		val += i
	}

	return val

}

func generateIntGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)

	for idx := range grid {
		grid[idx] = make([]int, cols)
	}

	return grid

}

func Soal1(n int) [][]int {
	len := sumOfNum(n)
	result := generateIntGrid(len, len)
	start := 0
	// end := 0

	for box := 1; box <= n; box++ {
		val := 1
		for i := 0; i < box; i++ {
			for j := 0; j < box; j++ {
				if i%2 == 0 {
					result[start+i][start+j] = val
				} else {
					result[start+i][start+box-j-1] = val
				}
				val += 2
			}
		}
		start += box

	}
	return result

}
