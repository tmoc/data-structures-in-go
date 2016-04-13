package bstree

import "testing"

func TestInsert(t *testing.T) {
	bst := new(BinarySearchTree)

	bst.Insert(1)
	bst.Insert(300)
	bst.Insert(18)
	bst.Insert(45)
	bst.Insert(2)

	if v := bst.Contains(18); v != true {
		t.Errorf("bst.Contains(18) should be true, got %v", v)
	}

	bst.Insert(2)

	if v := bst.Contains(2); v != true {
		t.Errorf("bst.Contains(2) should be true, got %v", v)
	}
}

func TestRemove(t *testing.T) {
	bst := new(BinarySearchTree)

	bst.Insert(1)
	bst.Insert(300)
	bst.Insert(18)
	bst.Insert(45)
	bst.Insert(2)

	if v := bst.Contains(300); v != true {
		t.Errorf("bst.Contains(300) should be true, got %v", v)
	}
	if v := bst.Contains(2); v != true {
		t.Errorf("bst.Contains(2) should be true, got %v", v)
	}

	bst.Remove(300)

	if v := bst.Contains(300); v != false {
		t.Errorf("bst.Contains(300) should be false, got %v", v)
	}
	bst.Remove(2)

	if v := bst.Contains(2); v != false {
		t.Errorf("bst.Contains(2) should be false, got %v", v)
	}
}

func TestMax(t *testing.T) {
	bst := new(BinarySearchTree)

	bst.Insert(1)
	bst.Insert(300)
	bst.Insert(18)
	bst.Insert(45)
	bst.Insert(2)

	if v := bst.Max(); v != 300 {
		t.Errorf("bst.Max() should be 300, got %v", v)
	}
}

func TestMin(t *testing.T) {
	bst := new(BinarySearchTree)

	bst.Insert(1)
	bst.Insert(300)
	bst.Insert(18)
	bst.Insert(45)
	bst.Insert(2)

	if v := bst.Min(); v != 1 {
		t.Errorf("bst.Min() should be 1, got %v", v)
	}
}

func TestPreOrder(t *testing.T) {
	bst := new(BinarySearchTree)
	items := []int{}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)

	bst.PreOrder(func(x int) { items = append(items, x) })

	if v := items[0]; v != 2 {
		t.Errorf("items[0] should be 2, got %v", v)
	}
	if v := items[1]; v != 1 {
		t.Errorf("items[1] should be 1, got %v", v)
	}
	if v := items[2]; v != 3 {
		t.Errorf("items[2] should be 3, got %v", v)
	}
}

func TestPostOrder(t *testing.T) {
	bst := new(BinarySearchTree)
	items := []int{}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)

	bst.PostOrder(func(x int) { items = append(items, x) })

	if v := items[0]; v != 1 {
		t.Errorf("items[0] should be 1, got %v", v)
	}
	if v := items[1]; v != 3 {
		t.Errorf("items[1] should be 3, got %v", v)
	}
	if v := items[2]; v != 2 {
		t.Errorf("items[2] should be 2, got %v", v)
	}
}

func TestInOrder(t *testing.T) {
	bst := new(BinarySearchTree)
	items := []int{}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)

	bst.InOrder(func(x int) { items = append(items, x) })

	if v := items[0]; v != 1 {
		t.Errorf("items[0] should be 1, got %v", v)
	}
	if v := items[1]; v != 2 {
		t.Errorf("items[1] should be 2, got %v", v)
	}
	if v := items[2]; v != 3 {
		t.Errorf("items[2] should be 3, got %v", v)
	}
}

func TestBreadthFirst(t *testing.T) {
	bst := new(BinarySearchTree)
	items := []int{}
	bst.Insert(2)
	bst.Insert(1)
	bst.Insert(3)

	bst.BreadthFirst(func(x int) { items = append(items, x) })

	if v := items[0]; v != 2 {
		t.Errorf("items[0] should be 2, got %v", v)
	}
	if v := items[1]; v != 1 {
		t.Errorf("items[1] should be 1, got %v", v)
	}
	if v := items[2]; v != 3 {
		t.Errorf("items[2] should be 3 got %v", v)
	}
}
