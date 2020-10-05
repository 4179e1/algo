package dp

import (
	"github.com/4179e1/algo/utils"
	"sort"
	"testing"
)

func TestCoinChange(t *testing.T) {
	type coinChg struct {
		coins  []int
		amount int
		want   []int
	}

	cases := []coinChg{
		//{
		//	[]int{1, 2, 5},
		//	3,
		//	[]int{1, 2},
		//},
		{
			[]int{1, 2, 5},
			11,
			[]int{1, 5, 5},
		},
	}

	for _, item := range cases {
		got := CoinChange(item.coins, item.amount)
		sort.Ints(item.want)
		sort.Ints(got)
		if !utils.IntsEqual(got, item.want) {
			t.Errorf("CoinExchang %v => %v, got %v, want %v\n", item.coins, item.amount, got, item.want)
		}

	}
}
