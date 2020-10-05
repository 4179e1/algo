package dp

import (
	"github.com/4179e1/algo/utils"
)

func CoinChange(coins []int, amount int) []int {

	var memo = make(map[int][]int)

	var dpf func(int) []int
	dpf = func(n int) []int {

		// 查找dptable，防止重复计算
		res, found := memo[n]
		if found {
			return res
		}

		// 结束条件
		if n == 0 {
			return []int{}
		}
		// 要换的硬币小于0,无解
		if n < 0 {
			return nil
		}

		// 哨兵
		var sential []int
		for i := 0; i < amount; i++ {
			sential = append(sential, 0)
		}

		res = sential
		for _, coin := range coins {
			subsolution := dpf(n - coin)
			if subsolution == nil {
				// 子问题无解
				continue
			}

			// 从所有结果里面取一个最短的
			cur := append(utils.IntsCopy(subsolution), coin)
			//fmt.Printf("Exchange %v best? %v <---> %v\n", n, res, cur)
			res = utils.IntSliceShorter(res, cur)
		}

		if utils.IntsEqual(res, sential) {
			memo[n] = nil
			return nil
		}
		memo[n] = res
		return res
	}

	return dpf(amount)
}

func CoinChangeIter(coins []int, amount int) []int {

	var sential []int
	for i := 0; i <= amount; i++ {
		sential = append(sential, 0)
	}

	var dp = make(map[int][]int)
	for i := 1; i <= amount; i++ {
		dp[i] = sential
	}
	dp[0] = []int{}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			// 无解
			if (i - coin) < 0 {
				continue
			}
			cur := append(utils.IntsCopy(dp[i-coin]), coin)
			//fmt.Printf("Exchange %v best? %v <---> %v\n", i, dp[i], cur)
			dp[i] = utils.IntSliceShorter(dp[i], cur)
		}
	}
	if utils.IntsEqual(dp[amount], sential) {
		return nil
	}
	return dp[amount]
}
