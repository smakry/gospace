package group

import (
	"testing"
)

func TestCombinStr(t *testing.T) {
	str := []rune("干得好a")
	CombinStr(str, make([]rune, 0, 0))
}
