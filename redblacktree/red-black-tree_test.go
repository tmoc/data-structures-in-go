package redblacktree

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestInsert(t *testing.T) {
	rbTree := new(RedBlackTree)

	rbTree.Insert(1)
	rbTree.Insert(2)
	rbTree.Insert(3)
	spew.Dump(rbTree)

	if v := rbTree.Contains(3); v != true {
		t.Errorf("rbTree.Contains(3) should be true, got %v", v)
	}
}
