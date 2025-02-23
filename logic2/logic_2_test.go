package logic2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Logic2Test struct {
	name     string
	request  int
	expected [][]int
}

func tableTestRunner(t *testing.T, tests []Logic2Test, funcToRun func(n int) [][]int) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := funcToRun(test.request)
			assert.Equal(t, test.expected, result)
		})

	}

}

func TestSoal1(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 1 n=5",
			request: 5,
			expected: [][]int{
				{1, 3, 5, 7, 9},
				{1, 3, 5, 7, 9},
				{1, 3, 5, 7, 9},
				{1, 3, 5, 7, 9},
				{1, 3, 5, 7, 9}},
		},
		{
			name:     "Soal 1 n=1",
			request:  1,
			expected: [][]int{{1}},
		},
		{
			name:     "Soal 1 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 1 n=6",
			request: 6,
			expected: [][]int{
				{1, 3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9, 11},
				{1, 3, 5, 7, 9, 11},
			},
		},
	}
	tableTestRunner(t, tests, Soal1)
}
func TestSoal2(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 2 n=5",
			request: 5,
			expected: [][]int{
				{2, 4, 6, 8, 10},
				{2, 4, 6, 8, 10},
				{2, 4, 6, 8, 10},
				{2, 4, 6, 8, 10},
				{2, 4, 6, 8, 10},
			},
		},
		{
			name:     "Soal 2 n=1",
			request:  1,
			expected: [][]int{{2}},
		},
		{
			name:     "Soal 2 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 2 n=6",
			request: 6,
			expected: [][]int{
				{2, 4, 6, 8, 10, 12},
				{2, 4, 6, 8, 10, 12},
				{2, 4, 6, 8, 10, 12},
				{2, 4, 6, 8, 10, 12},
				{2, 4, 6, 8, 10, 12},
				{2, 4, 6, 8, 10, 12},
			},
		},
	}

	tableTestRunner(t, tests, Soal2)
}
func TestSoal3(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 3 n=5",
			request: 5,
			expected: [][]int{
				{1, 3, 5, 7, 9},
				{11, 13, 15, 17, 19},
				{21, 23, 25, 27, 29},
				{31, 33, 35, 37, 39},
				{41, 43, 45, 47, 49},
			},
		},
		{
			name:     "Soal 3 n=1",
			request:  1,
			expected: [][]int{{1}},
		},
		{
			name:     "Soal 3 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 3 n=4",
			request: 4,
			expected: [][]int{
				{1, 3, 5, 7},
				{9, 11, 13, 15},
				{17, 19, 21, 23},
				{25, 27, 29, 31},
			},
		},
	}
	tableTestRunner(t, tests, Soal3)
}
func TestSoal4(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 4 n=5",
			request: 5,
			expected: [][]int{
				{1, 4, 7, 10, 13},
				{16, 19, 22, 25, 28},
				{31, 34, 37, 40, 43},
				{46, 49, 52, 55, 58},
				{61, 64, 67, 70, 73},
			},
		},
	}
	tableTestRunner(t, tests, Soal4)
}
func TestSoal5(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 5 n=2",
			request: 2,
			expected: [][]int{
				{1, 3},
				{7, 5},
			},
		},
		{
			name:    "Soal 5 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 5 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 5 n=3",
			request: 3,
			expected: [][]int{
				{1, 3, 5},
				{11, 9, 7},
				{13, 15, 17},
			},
		},
	}
	tableTestRunner(t, tests, Soal5)
}
func TestSoal6(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 6 n=2",
			request: 2,
			expected: [][]int{
				{1, 4},
				{10, 7},
			},
		},
		{
			name:    "Soal 6 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 6 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 6 n=3",
			request: 3,
			expected: [][]int{
				{1, 4, 7},
				{16, 13, 10},
				{19, 22, 25},
			},
		},
	}
	tableTestRunner(t, tests, Soal6a)

}
func TestSoal7(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 7 n=2",
			request: 2,
			expected: [][]int{
				{1, 0},
				{0, 3},
			},
		},
		{
			name:    "Soal 7 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 7 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 7 n=3",
			request: 3,
			expected: [][]int{
				{1, 0, 0},
				{0, 3, 0},
				{0, 0, 5},
			},
		},
	}
	tableTestRunner(t, tests, Soal7)
}
func TestSoal8(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 8 n=2",
			request: 2,
			expected: [][]int{
				{0, 3},
				{1, 0},
			},
		},
		{
			name:    "Soal 8 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 8 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 8 n=3",
			request: 3,
			expected: [][]int{
				{0, 0, 5},
				{0, 3, 0},
				{1, 0, 0},
			},
		},
	}
	tableTestRunner(t, tests, Soal8)
}
func TestSoal9(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 9 n=3",
			request: 3,
			expected: [][]int{
				{1, 0, 5},
				{0, 3, 0},
				{1, 0, 5},
			},
		},
		{
			name:    "Soal 9 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 9 n=0",
			request:  0,
			expected: [][]int{},
		},
	}
	tableTestRunner(t, tests, Soal9)
}
func TestSoal10(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 10 n=3",
			request: 3,
			expected: [][]int{
				{1, 0, 0},
				{1, 3, 0},
				{1, 3, 5},
			},
		},
		{
			name:    "Soal 10 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 10 n=0",
			request:  0,
			expected: [][]int{},
		},
	}
	tableTestRunner(t, tests, Soal10)
}
func TestSoal11(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 11 n=3",
			request: 3,
			expected: [][]int{
				{1, 3, 5},
				{0, 3, 5},
				{0, 0, 5},
			},
		},
		{
			name:    "Soal 11 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 11 n=0",
			request:  0,
			expected: [][]int{},
		},
	}
	tableTestRunner(t, tests, Soal11)
}
func TestSoal12(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 12 n=3",
			request: 3,
			expected: [][]int{
				{1, 0, 5},
				{1, 3, 5},
				{1, 0, 5},
			},
		},
		{
			name:    "Soal 12 n=1",
			request: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "Soal 12 n=0",
			request:  0,
			expected: [][]int{},
		},
		{
			name:    "Soal 12 n=5",
			request: 5,
			expected: [][]int{
				{1, 0, 0, 0, 9},
				{1, 3, 0, 7, 9},
				{1, 3, 5, 7, 9},
				{1, 3, 0, 7, 9},
				{1, 0, 0, 0, 9},
			},
		},
	}
	tableTestRunner(t, tests, Soal12)
}
