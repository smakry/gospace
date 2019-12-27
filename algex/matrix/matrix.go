package matrix

import (
	"fmt"
)

// purpose: 旋转二维数组n*n逆时针旋转45°后打印
// methods: 索引增长方向越界表示遍历行结束，使用dir表示当前的索引变化，绝对值为下一个起始（行或者列）

func printRotate(arr [][]int) {
	if arr == nil { return }

	len := len(arr)
	for x, y, dir := 0, len - 1, len - 1; x < len && y < len; {
		fmt.Print(arr[x][y], "\t")
		if (x + 1 >= len) || (y + 1 >= len) {
			fmt.Println()

			if (x == len -1) && (y == 0) {
				return
			}

			// need change dir
			if dir == 0 {
				x = 0
				y = 1
			}
			dir--

			// increase dir
			if dir >= 0 {
				x = 0
				y = dir
			} else {
				x = -dir
				y = 0
			}
		} else {
			x++
			y++
		}
	}
}
