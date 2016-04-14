package avltree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)

	if v := avlTree.Contains(18); v != true {
		t.Errorf("avlTree.Contains(18) should be true, got %v", v)
	}

	if v := avlTree.Contains(2); v != true {
		t.Errorf("avlTree.Contains(2) should be true, got %v", v)
	}
}

func TestRemove(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)
	avlTree.Insert(2)

	if v := avlTree.Contains(300); v != true {
		t.Errorf("avlTree.Contains(300) should be true, got %v", v)
	}
	if v := avlTree.Contains(2); v != true {
		t.Errorf("avlTree.Contains(2) should be true, got %v", v)
	}

	avlTree.Remove(300)
	if v := avlTree.Contains(300); v != false {
		t.Errorf("avlTree.Contains(300) should be false, got %v", v)
	}
	avlTree.Remove(2)
	if v := avlTree.Contains(2); v != false {
		t.Errorf("avlTree.Contains(2) should be false, got %v", v)
	}
}

func TestMax(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)
	avlTree.Insert(2)

	if v := avlTree.Max(); v != 300 {
		t.Errorf("avlTree.Max() should be 300, got %v", v)
	}
}

func TestMin(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)
	avlTree.Insert(2)

	if v := avlTree.Min(); v != 1 {
		t.Errorf("avlTree.Min() should be 1, got %v", v)
	}
}

func TestPreOrder(t *testing.T) {
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.PreOrder(func(x int) { items = append(items, x) })

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
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.PostOrder(func(x int) { items = append(items, x) })

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
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.InOrder(func(x int) { items = append(items, x) })

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
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.BreadthFirst(func(x int) { items = append(items, x) })

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
