package greedy

import (
//	"fmt"
)

// purpose: n个作业放到不同的系统（执行时间各不相同），求总耗时最低
// methods: 贪心算法计算所有的当前用时 + 新任务的用时最小的

func LeastExecuteTime(sys []int, jobs int) []int {
	if sys == nil || jobs <= 0 {
		return nil
	}

	do := len(sys)
	needTime := make([]int, do, do)
	for i := 1; i <= jobs; i++ {
		minTime := needTime[0] + sys[0]
		minIndex := 0
		
		for j := 1; j < do; j++ {
			if minTime > needTime[j] + sys[j] {
				minTime = needTime[j] + sys[j]
				minIndex = j
			}
		}

		needTime[minIndex] = minTime
	}

	return needTime
}

