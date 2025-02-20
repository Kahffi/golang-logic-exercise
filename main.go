package main

import (
	slice "github.com/Kahffi/go-print-slice"
	"github.com/Kahffi/golang-logic-exercise/logic1"
	"github.com/Kahffi/golang-logic-exercise/logic2"
)

func main() {
	slice.PrintSlice(logic1.Soal1(10), "logic1 - soal 1 (n=10)")
	slice.PrintSlice2D(logic2.Soal1(9), "logic2 - soal 1 (n=9)")
	slice.PrintSlice(logic1.Soal11(10), "logic1 - soal 1 (n=10)")
	slice.PrintSlice(logic1.Soal9(11), "logic1 - soal 9 (n=11)")
	slice.PrintSlice(logic1.Soal12(12), "logic1 - soal 12 (n=12)")
	slice.PrintSlice2D(logic2.Soal2(9), "logic2 - soal 2 (n=9)")
	slice.PrintSlice2D(logic2.Soal3(9), "logic2 - soal 3 (n=9)")
	slice.PrintSlice2D(logic2.Soal4(9), "logic2 - soal 4 (n=9)")
	slice.PrintSlice2D(logic2.Soal5(9), "logic2 - soal 5 (n=9)")
	slice.PrintSlice2D(logic2.Soal6a(9), "logic2 - soal 6 (n=9)")
	slice.PrintSlice2D(logic2.Soal7(9), "logic2 - soal 7 (n=9)")
	slice.PrintSlice2D(logic2.Soal8(9), "logic2 - soal 8 (n=9)")
}
