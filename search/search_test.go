package search

import (
	"github.com/4179e1/algo/utils"
	"sort"
	"testing"
)

var unorderedCases = [][]int{
	[]int{54, 57, 23, 84, 98, 10, 11, 12, 68, 77, 16, 18, 29, 33, 48},
	[]int{32, 23, 1},
	[]int{23},
}

var orderedCases [][]int

func init() {
	for _, l := range unorderedCases {
		nl := utils.IntsCopy(l)
		sort.Ints(nl)
		orderedCases = append(orderedCases, nl)

	}
}

type searchFunc func(int, []int) int

func search(t *testing.T, sf searchFunc, key int, a []int, want int) {
	got := sf(key, a)
	if got != want {
		t.Errorf("Serach %v from %v, expect %v got %v", key, a, want, got)
	}
}

func iterSearch(key int, a []int) int {
	for i, v := range a {
		if key == v {
			return i
		}
	}

	return -1
}

func searchWrapper(t *testing.T, sf searchFunc, bin bool) {

	keys := []int{50, 23}

	cases := unorderedCases
	if bin {
		cases = orderedCases
	}

	for _, key := range keys {
		var wants []int
		for _, c := range cases {
			wants = append(wants, iterSearch(key, c))
		}
		for i, c := range cases {
			search(t, sf, key, c, wants[i])
		}
	}

	for key := range keys {
		search(t, sf, key, []int{}, -1)
	}

}

func unorderTest(t *testing.T, sf searchFunc) {
	searchWrapper(t, sf, false)
}

func orderedTest(t *testing.T, sf searchFunc) {
	searchWrapper(t, sf, true)
}

func TestBinarySearch(t *testing.T) {
	orderedTest(t, BinarySearch)
}
