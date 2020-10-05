package dp

import (
	"github.com/4179e1/algo/utils"
)

func CoinChange(coins []int, amount int) []int {

	var dpf func(int) []int
	dpf = func(n int) []int {

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

		res := sential
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
			return nil
		}
		return res
	}

	return dpf(amount)
}
