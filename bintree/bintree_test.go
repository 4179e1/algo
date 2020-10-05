package bintree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestNewFromSliceEmpty(t *testing.T) {
	is := []interface{}{}
	bt := NewFromSlice(is)

	if bt != nil {
		t.Errorf("tree shoule be nil\n")
	}
}
func TestNewFromSlice(t *testing.T) {
	is := []interface{}{3, 9, 20, nil, nil, 15, 7}
	bt := NewFromSlice(is)

	spew.Dump(bt)
}
