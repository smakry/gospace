package bitset

import (
	"testing"
)

func TestBitSonSet(t *testing.T) {
	arr := []int{11, 2, 456, 97, 11, 2}
	BitSonSet(arr, RemoveDuplicate)
}

func TestIterationSonSet(t *testing.T) {
	arr := []int{11, 2, 456, 97}
	IterationSonSet(arr)
}