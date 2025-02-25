package logic3

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
				{1, 0, 0, 0, 0},
				{5, 3, 0, 0, 0},
				{7, 9, 11, 0, 0},
				{19, 17, 15, 13, 0},
				{21, 23, 25, 27, 29},
			},
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
	}
	tableTestRunner(t, tests, Soal1)
}
func TestSoal2(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 1 n=5",
			request: 5,
			expected: [][]int{
				{1, 3, 5, 7, 9},
				{0, 17, 15, 13, 11},
				{0, 0, 19, 21, 23},
				{0, 0, 0, 27, 25},
				{0, 0, 0, 0, 29},
			},
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
	}
	tableTestRunner(t, tests, Soal2)
}
func TestSoal3(t *testing.T) {
	tests := []Logic2Test{
		{
			name:    "Soal 1 n=5",
			request: 5,
			expected: [][]int{
				{2, 4, 6, 8, 10},
				{0, 13, 16, 19, 22},
				{0, 0, 28, 26, 24},
				{0, 0, 0, 31, 34},
				{0, 0, 0, 0, 36},
			},
		},
		{
			name:     "Soal 1 n=1",
			request:  1,
			expected: [][]int{{2}},
		},
		{
			name:     "Soal 1 n=0",
			request:  0,
			expected: [][]int{},
		},
	}
	tableTestRunner(t, tests, Soal3)
}
