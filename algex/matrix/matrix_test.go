package matrix

import (
	"testing"
	"fmt"
)


func TestRotate(t *testing.T) {
	arry := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	printRotate(arry)

	for _, v1 := range arry {
		for _, v2 := range v1 {
			fmt.Print(v2, "\t")
		}
		fmt.Println()
	}
}