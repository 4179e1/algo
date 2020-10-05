package bintree

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestNewFromSlice(t *testing.T) {
	is := []interface{}{3, 9, 20, nil, nil, 15, 7}
	bt := NewFromSlice(is)

	spew.Dump(bt)
}
