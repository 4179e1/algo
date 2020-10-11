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

type backPackFunc func([]int, []int, int) int

func backPackTest(t *testing.T, bpf backPackFunc) {
	type bpEntry struct {
		costs  []int
		values []int
		volume int
		ans    int
	}

	cases := []bpEntry{
		{
			costs:  []int{1, 1, 1, 1, 2},
			values: []int{1, 2, 3, 4, 5},
			volume: 3,
			ans:    9,
		},
		{
			costs:  []int{1, 1, 1, 1, 1},
			values: []int{1, 2, 3, 4, 5},
			volume: 3,
			ans:    12,
		},
	}

	for _, item := range cases {

		ans := bpf(item.costs, item.values, item.volume)
		if ans != item.ans {
			t.Errorf("BackPack Size %d Item costs %v values %v, expected %d got %v",
				item.volume, item.costs, item.values, item.ans, ans)
		}
	}

}

func TestZeroOneBackPack(t *testing.T) {
	backPackTest(t, ZeroOneBackPack)
}

func TestZeroOneBackPackSO(t *testing.T) {
	backPackTest(t, ZeroOneBackPackSpaceOptimized)
}
