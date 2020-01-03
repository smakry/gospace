package square

import (
	"testing"
	"fmt"
)

func TestSquareIsN(t *testing.T) {
	tdata := []int{
		0,
		1,
		2,
		15,
		25,
		99,
	}
	
	for _, v := range tdata {
		if ok, m := SquareIsN(v); ok {
			fmt.Println(m, v)
		}
	}
}

func TestIsOf2Square(t *testing.T) {
	tdata := []int{
		0,
		1,
		2,
		15,
		16,
		64,
	}
	
	for _, v := range tdata {
		if IsOf2Square(v) {
			fmt.Println(v)
		}
	}
}
