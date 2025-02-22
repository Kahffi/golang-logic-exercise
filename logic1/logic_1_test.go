package logic1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Logic1Test struct {
	name     string
	request  int
	expected []int
}

func runTableTests(t *testing.T, tests []Logic1Test, funcToTest func(n int) []int) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := funcToTest(test.request)
			assert.Equal(t, test.expected, res)

		})
	}

}

func TestSoal1(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 1 n=8",
			request:  8,
			expected: []int{1, 3, 5, 7, 9, 11, 13, 15},
		},
		{
			name:     "Soal 1 n=5",
			request:  5,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "Soal 1 n=10",
			request:  10,
			expected: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		},
	}

	runTableTests(t, tests, Soal1)

}

func TestSoal2(t *testing.T) {

	tests := []Logic1Test{
		{
			name:     "Soal 2 n=5",
			request:  5,
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "Soal 2 n=3",
			request:  3,
			expected: []int{2, 4, 6},
		},
		{
			name:     "Soal 2 n=10",
			request:  10,
			expected: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
	}
	runTableTests(t, tests, Soal2)
}

func TestSoal3(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 3 n=10",
			request:  10,
			expected: []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
		},
		{
			name:     "Soal 3 n=9",
			request:  9,
			expected: []int{3, 6, 9, 12, 15, 18, 21, 24, 27},
		},
		{
			name:     "Soal 3 n=12",
			request:  12,
			expected: []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36},
		},
		{
			name:     "Soal 3 n=15",
			request:  15,
			expected: []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45},
		},
	}

	runTableTests(t, tests, Soal3)

}

func TestSoal4(t *testing.T) {
	tests := []Logic1Test{{
		name:     "Soal 4 n=1",
		request:  1,
		expected: []int{1},
	},
		{

			name:     "Soal 4 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 4 n=10",
			request:  10,
			expected: []int{19, 17, 15, 13, 11, 9, 7, 5, 3, 1}},
	}

	runTableTests(t, tests, Soal4)

}
func TestSoal5(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 5 n=1",
			request:  1,
			expected: []int{2},
		},
		{

			name:     "Soal 5 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 5 n=10",
			request:  10,
			expected: []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2},
		},
	}

	runTableTests(t, tests, Soal5)

}
func TestSoal6(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 6 n=1",
			request:  1,
			expected: []int{3},
		},
		{

			name:     "Soal 6 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 6 n=10",
			request:  10,
			expected: []int{30, 27, 24, 21, 18, 15, 12, 9, 6, 3},
		},
	}

	runTableTests(t, tests, Soal6)

}
func TestSoal7(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 7 n=1",
			request:  1,
			expected: []int{1},
		},
		{

			name:     "Soal 7 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 7 n=10",
			request:  10,
			expected: []int{1, 3, 5, 7, 9, 9, 7, 5, 3, 1},
		},
		{
			name:     "Soal 7 n=11",
			request:  11,
			expected: []int{1, 3, 5, 7, 9, 11, 9, 7, 5, 3, 1},
		},
	}

	runTableTests(t, tests, Soal7)

}

func TestSoal8(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 8 n=1",
			request:  1,
			expected: []int{2},
		},
		{
			name:     "Soal 8 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 8 n=10",
			request:  10,
			expected: []int{2, 4, 6, 8, 10, 10, 8, 6, 4, 2},
		},
		{
			name:     "Soal 8 n=11",
			request:  11,
			expected: []int{2, 4, 6, 8, 10, 12, 10, 8, 6, 4, 2},
		},
	}
	runTableTests(t, tests, Soal8)
}

func TestSoal9(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 9 n=1",
			request:  1,
			expected: []int{3},
		},
		{
			name:     "Soal 9 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 9 n=10",
			request:  10,
			expected: []int{3, 6, 9, 12, 15, 15, 12, 9, 6, 3},
		},
		{
			name:     "Soal 9 n=11",
			request:  11,
			expected: []int{3, 6, 9, 12, 15, 18, 15, 12, 9, 6, 3},
		},
	}
	runTableTests(t, tests, Soal9)
}

func TestSoal10(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 10 n=1",
			request:  1,
			expected: []int{2},
		},
		{
			name:     "Soal 10 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 10 n=10",
			request:  10,
			expected: []int{2, 0, 4, 0, 8, 0, 16, 0, 32, 0},
		},
	}
	runTableTests(t, tests, Soal10)

}
func TestSoal11(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 11 n=1",
			request:  1,
			expected: []int{0},
		},
		{
			name:     "Soal 11 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 11 n=10",
			request:  10,
			expected: []int{0, 1, 0, 3, 0, 6, 0, 12, 0, 24},
		},
	}
	runTableTests(t, tests, Soal11)
}
func TestSoal12(t *testing.T) {
	tests := []Logic1Test{
		{
			name:     "Soal 12 n=1",
			request:  1,
			expected: []int{1},
		},
		{
			name:     "Soal 12 n=0",
			request:  0,
			expected: []int{},
		},
		{
			name:     "Soal 12 n=12",
			request:  12,
			expected: []int{1, 3, 5, 7, 1, 3, 5, 7, 1, 3, 5, 7},
		},
		{
			name:     "Soal 12 n=11",
			request:  11,
			expected: []int{1, 3, 5, 7, 1, 3, 5, 7, 1, 3, 5},
		},
	}
	runTableTests(t, tests, Soal12)
}
