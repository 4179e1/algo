package dp

import (
	"sort"
	"testing"

	"github.com/4179e1/algo/utils"
)

type coinChangeFunc func([]int, int) []int

func coinChangeTest(t *testing.T, ccf coinChangeFunc) {
	type coinChg struct {
		coins  []int
		amount int
		want   []int
	}

	cases := []coinChg{
		{
			[]int{2, 5},
			0,
			[]int{},
		},
		{
			[]int{2, 5},
			3,
			nil,
		},
		{
			[]int{1, 2, 5},
			3,
			[]int{1, 2},
		},
		{
			[]int{1, 2, 5},
			11,
			[]int{1, 5, 5},
		},
	}

	for _, item := range cases {
		got := ccf(item.coins, item.amount)
		sort.Ints(item.want)
		sort.Ints(got)
		if !utils.IntsEqual(got, item.want) {
			t.Errorf("CoinExchang %v => %v, got %v, want %v\n", item.coins, item.amount, got, item.want)
		} else {
			//t.Logf("CoinExchang %v => %v, got %v, want %v\n", item.coins, item.amount, got, item.want)
		}

	}
}

func TestCoinChange(t *testing.T) {
	coinChangeTest(t, CoinChange)
}
func TestCoinChangeIter(t *testing.T) {
	coinChangeTest(t, CoinChangeIter)
}
