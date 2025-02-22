package logic1

func Soal1(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = i*2 + 1
	}

	return result
}

func Soal2(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = i*2 + 2
	}

	return result
}

func Soal3(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = (i + 1) * 3
	}

	return result
}

func Soal4(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = (n-i-1)*2 + 1
	}

	return result
}

func Soal5(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = (n-i-1)*2 + 2
	}

	return result
}

func Soal6(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = ((n - i - 1) + 1) * 3
	}

	return result
}

func Soal7(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		if i != 0 && i >= n/2 {
			result[i] = (n-i-1)*2 + 1
		} else {

			result[i] = i*2 + 1
		}
	}

	return result
}

func Soal8(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		if i != 0 && i >= n/2 {
			result[i] = (n-i-1)*2 + 2
		} else {

			result[i] = i*2 + 2
		}
	}

	return result
}

func Soal9(n int) []int {
	var result = make([]int, n)

	for i := 0; i < n; i++ {
		if i != 0 && i >= n/2 {
			result[i] = ((n - i - 1) + 1) * 3
		} else {

			result[i] = (i + 1) * 3
		}
	}

	return result
}

func Soal10(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			result[i] = 0
		} else {
			if i == 0 {
				result[i] = 2
			} else {
				result[i] = result[i-2] * 2
			}
		}
	}

	return result
}

func Soal11(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = 0
		} else {
			if i <= 3 {
				result[i] = i
			} else {
				result[i] = result[i-2] * 2
			}
		}
	}

	return result
}

func Soal12(n int) []int {
	result := make([]int, n)

	for i := range result {
		result[i] = (i%4)*2 + 1
	}

	return result
}
