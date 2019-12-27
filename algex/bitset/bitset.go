package bitset

import (
	"fmt"
)

// purpose: 求集合所有子集
// methods: 1.对应为二进制位数，但这种方式缺陷一个是重复值（可去重就是），再者就是二进制的位数有限
// 			2.使用迭代方式，当前结果 = 前面的结果 + 前面的结果分别‘追加’当前值 + 当前值

const maxBit = 64

func BitSonSet(st []int, rmdup func([]int) []int) {
	if st == nil { return }

	if rmdup != nil {
		st = append(st[0:0], rmdup(st)[:]...)
	}

	if len(st) > maxBit {
		panic("bit over")
	}

	var all int64
	all = 1 << len(st) - 1 // all son set has (2 ^ n - 1)
	for i := int64(1); i <= all; i++ {
		for j := 0; j < len(st); j++ {
			if (i & (1 << j)) == 1 << j {
				fmt.Print(st[j], " ")
			}
		}
		fmt.Println()
	}
}

func RemoveDuplicate(s []int) []int {
	if s == nil {
		return nil
	}

	res := make([]int, 0, 0)
	tmpSet := make(map[int]bool)
	for _, v := range s {
		if _, ok := tmpSet[v]; !ok {
			res = append(res, v)
			tmpSet[v] = true
		}
	}
	return res
}

// iteration methods: 前面的每个 + 前面每个分别‘追加当前’ + 当前
// eg:{a, b, c, d}
// step1: {a}
// step2: {a, ab, b} 
// step3: {a, ab, b, ac, abc, bc, c} 
// step4: {a, ab, b, ac, abc, bc, c, ad, abd, bd, acd, abcd, bcd, cd, d} 
func IterationSonSet(st []int) {
	if st == nil || len(st) == 0 { return }

	res := make([][]int, 0, 0)
	for _, v := range st {
		for _, v2 := range res {
			tmp := v2[:]
			res = append(res, append(tmp, v))
		}

		res = append(res, []int{v})
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

