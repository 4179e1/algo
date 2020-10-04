package backtrack

import (
	"fmt"
	"testing"
)

func TestNQueen(t *testing.T) {
	cases := map[int]int{
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		8: 92,
	}

	for n, want := range cases {
		cbs := NQueen(n)
		got := len(cbs)
		if want != len(cbs) {
			t.Errorf("%v Queens want %v got %v", n, want, got)
		}
	}
}

func TestNQueenVisulize(t *testing.T) {
	cbs := NQueen(4)
	fmt.Println(cbs)
}
