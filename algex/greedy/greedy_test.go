package greedy

import(
	"fmt"
	"testing"
)

func TestLeastExecuteTime(t *testing.T) {
	sys := []int{7, 10, 6, 8}
	res := LeastExecuteTime(sys, 6)
	fmt.Println(res)
}