package bintree

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewFromSlice(t *testing.T) {
	is := []int{3, 9, 20, -1, -1, 15, 7}
	bt := NewFromSlice(is)

	spew.Dump(bt)
}
