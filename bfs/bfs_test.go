package bfs

import (
	"testing"

	"github.com/4179e1/algo/bintree"
)

func TestMinBinTreeDepth(t *testing.T) {
	type Case struct {
		treeSlice []interface{}
		want      int
	}

	cases := []Case{
		{
			[]interface{}{},
			0,
		},
		{
			[]interface{}{1},
			1,
		},
		{
			[]interface{}{3, 9, nil},
			2,
		},
		{
			[]interface{}{3, nil, 20},
			2,
		},
		{
			[]interface{}{3, 9, 20, nil, 21, 15, 7},
			3,
		},
	}

	for _, item := range cases {
		tree := bintree.NewFromSlice(item.treeSlice)
		//spew.Dump(tree)
		got := MinBinTreeDepth(tree)
		if got != item.want {
			t.Errorf("MinDepth for %v want %v got %v\n", item.treeSlice, item.want, got)
		}
	}
}
