package main

import (
	"gospace/lru"
	"fmt"
)

func main() {
	pu := lru.NewPageUse()

	pu.View(1)
	pu.View(2)
	pu.View(3)
	pu.View(4)
	pu.View(5)
	pu.View(6)
	pu.View(1)

	for ptr := pu.UseQu.FrontNode; ptr != nil; ptr = ptr.NextNode {
		fmt.Println(ptr.Data)
	}
}

