package dp

import (
	"github.com/4179e1/algo/utils"
)

/*
ZeroOneBackPack 01背包问题
costs: 物品消耗的空间
values: 物品的价值
volume： 背包的容量
*/
func ZeroOneBackPack(costs, values []int, volume int) int {
	if len(costs) != len(values) {
		panic("Length of item cost and values mismatch")
	}

	nums := len(costs)

	/*
		F[i, v] = max{F[i − 1, v], F[i − 1, v − Ci ] + Wi}
			表示当背包容量剩下v时，尝试放入物品i时的最优解（即此时背包中所有物品的价值和），但是有可能放不进去
			比如 F[5, 1] 表示容量剩下5,尝试放入物品1时的价值（可能放不进去）,它暗示着我们之前已经试过
				F[4, 1], F[3, 1], F[2, 1], F[1, 1]最后得到的最优解

		we need extra space, as index start from 1, [0, 0] has special meaning
	*/
	f := utils.NewIntMatrix(nums+1, volume+1)

	var i, v int
	for i = 1; i <= nums; i++ {
		// 数组从0开始，但是我们的循环从1开始
		cost := costs[i-1]
		value := values[i-1]
		for v = 1; v <= volume; v++ {
			if v >= cost {
				// 物品i能放进去
				// f[i-1][v] 表示同样容量下，放i之前哪件物品能取得的最优解
				// f[i-1][v-cost] 表示在减去当前物品cost的前提下，放i之前哪件物品的最优解
				f[i][v] = utils.IntMax(f[i-1][v], f[i-1][v-cost]+value)
			} else {
				// 放不进去 :(
				// 那么当前的值就是背包容量不变的情况下，上一轮最好取值
				// 比如物品5 f(5,v)放不进去，那么它就是前4件最好的取值f(4,v)，如此类推
				f[i][v] = f[i-1][v]
			}
			//fmt.Println(utils.IntMatrix(f))
		}
	}

	return f[nums][volume]
}

func ZeroOneBackPackSpaceOptimized(costs, values []int, volume int) int {
	if len(costs) != len(values) {
		panic("Length of item cost and values mismatch")
	}

	nums := len(costs)

	f := make([]int, volume+1)

	var i, v int
	for i = 1; i <= nums; i++ {
		cost := costs[i-1]
		value := values[i-1]
		for v = volume; v >= 1; v-- {
			if v >= cost {
				f[v] = utils.IntMax(f[v], f[v-cost]+value)
			}

			//fmt.Println(f)
		}
	}

	return f[volume]
}
