package utils

import "fmt"

func PrintArrayH(arr []int, label string) {
	fmt.Println(label)
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Print("\n\n")
}
