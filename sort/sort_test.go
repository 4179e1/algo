package sort

import (
	"sort"
	"testing"
)

var input = [][]int{
	[]int{1, 2, 3, 4, 5},
	[]int{5, 4, 3, 2, 1},
	[]int{2, 3, 5, 4, 1},
	[]int{1, 3, 5, 1, 3},
	[]int{3, 5, 1},
	[]int{2, 1},
	[]int{1},
	[]int{},
}

var expected = [][]int{}

func init() {
	for _, i := range input {
		nl := make([]int, len(i))
		copy(nl, i)
		sort.Ints(nl)
		expected = append(expected, nl)
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}

func sortFuncTest(t *testing.T, sf func([]int) []int) {
	for x := range input {
		ci := make([]int, len(input[x]))
		copy(ci, input[x])
		want := expected[x]
		got := sf(ci)

		if !Equal(got, want) {
			t.Errorf("Sort %v want %v got %v\n", ci, want, got)
		}
	}
}

func TestInsertSort(t *testing.T) {
	sortFuncTest(t, InsertSort)
}

func TestBubbleSort(t *testing.T) {
	sortFuncTest(t, BubbleSort)
}
