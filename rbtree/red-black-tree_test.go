package redblacktree

import "testing"

func TestInsert(t *testing.T) {
	rbTree := new(RedBlackTree)

	rbTree.Insert(1)
	rbTree.Insert(2)
	rbTree.Insert(3)
	rbTree.Insert(300)
	rbTree.Insert(18)
	rbTree.Insert(45)

	if v := rbTree.Contains(18); v != true {
		t.Errorf("rbTree.Contains(18) should be true, got %v", v)
	}

	if v := rbTree.Contains(2); v != true {
		t.Errorf("rbTree.Contains(2) should be true, got %v", v)
	}
}
