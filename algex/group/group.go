package group

import (
	"fmt"
)

// purpose: 给定字符串的组合
// methods: 1.递归选择字符，每次选择一个，在剩余的下一轮迭代选择
// 			2.非递归方式，就是用选择的结果放到下一轮循环中选择

func CombinStr(str []rune, res []rune) {
	if str == nil {
		if len(str) == 0 {
			fmt.Println(string(res))
		}
		return
	}

	for i := 0; i < len(str); i++ {
		var tmp []rune
		tmp = append(tmp, str[:i]...)
		tmp = append(tmp, str[i + 1:len(str)]...)
		CombinStr(tmp, append(res, str[i]))
	}
}

