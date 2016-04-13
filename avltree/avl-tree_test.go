package avltree

import "testing"

func TestInsert(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)

	if avlTree.Contains(18) != true {
		t.Errorf("avlTree.Contains(18) should be true")
	}

	if avlTree.Contains(2) != true {
		t.Errorf("avlTree.Contains(2) should be true")
	}
}

// func TestRemove(t *testing.T) {
// 	avlTree := new(AVLTree)
//
// 	avlTree.Insert(1)
// 	avlTree.Insert(300)
// 	avlTree.Insert(18)
// 	avlTree.Insert(45)
// 	avlTree.Insert(2)
//
// 	if avlTree.Contains(300) != true {
// 		t.Errorf("avlTree.Contains(300) should be true")
// 	}
// 	if avlTree.Contains(2) != true {
// 		t.Errorf("avlTree.Contains(2) should be true")
// 	}
//
// 	avlTree.Remove(300)
// 	if avlTree.Contains(300) == true {
// 		t.Errorf("avlTree.Contains(300) should be false")
// 	}
// 	avlTree.Remove(2)
// 	if avlTree.Contains(2) == true {
// 		t.Errorf("avlTree.Contains(2) should be false")
// 	}
// }

func TestMax(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)
	avlTree.Insert(2)

	if avlTree.Max() != 300 {
		t.Errorf("avlTree.Max() should be 300")
	}
}

func TestMin(t *testing.T) {
	avlTree := new(AVLTree)

	avlTree.Insert(1)
	avlTree.Insert(300)
	avlTree.Insert(18)
	avlTree.Insert(45)
	avlTree.Insert(2)

	if avlTree.Min() != 1 {
		t.Errorf("avlTree.Min() should be 1")
	}
}

func TestPreOrder(t *testing.T) {
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.PreOrder(func(x int) { items = append(items, x) })

	if items[0] != 2 {
		t.Errorf("items[0] should be 2")
	}
	if items[1] != 1 {
		t.Errorf("items[1] should be 1")
	}
	if items[2] != 3 {
		t.Errorf("items[2] should be 3")
	}
}

func TestPostOrder(t *testing.T) {
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.PostOrder(func(x int) { items = append(items, x) })

	if items[0] != 1 {
		t.Errorf("items[0] should be 1")
	}
	if items[1] != 3 {
		t.Errorf("items[1] should be 3")
	}
	if items[2] != 2 {
		t.Errorf("items[2] should be 2")
	}
}

func TestInOrder(t *testing.T) {
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.InOrder(func(x int) { items = append(items, x) })

	if items[0] != 1 {
		t.Errorf("items[0] should be 1")
	}
	if items[1] != 2 {
		t.Errorf("items[1] should be 2")
	}
	if items[2] != 3 {
		t.Errorf("items[2] should be 3")
	}
}

func TestBreadthFirst(t *testing.T) {
	avlTree := new(AVLTree)
	items := []int{}
	avlTree.Insert(2)
	avlTree.Insert(1)
	avlTree.Insert(3)

	avlTree.BreadthFirst(func(x int) { items = append(items, x) })

	if items[0] != 2 {
		t.Errorf("items[0] should be 2")
	}
	if items[1] != 1 {
		t.Errorf("items[1] should be 1, got %d", items[1])
	}
	if items[2] != 3 {
		t.Errorf("items[2] should be 3 got %d", items[2])
	}
}
