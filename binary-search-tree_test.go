package ds

import "testing"

func TestInsert(t *testing.T) {
  bst := new(BST)

  bst.Insert(1)
  bst.Insert(300)
  bst.Insert(18)
  bst.Insert(45)
  bst.Insert(2)

  if bst.Contains(18) != true {
    t.Errorf("bst.Contains(18) should be true")
  }

  bst.Insert(2)

  if bst.Contains(2) != true {
    t.Errorf("bst.Contains(2) should be true")
  }
}

func TestRemove(t *testing.T) {
  bst := new(BST)

  bst.Insert(1)
  bst.Insert(300)
  bst.Insert(18)
  bst.Insert(45)
  bst.Insert(2)

  if bst.Contains(300) != true {
    t.Errorf("bst.Contains(300) should be true")
  }
  if bst.Contains(2) != true {
    t.Errorf("bst.Contains(2) should be true")
  }

  bst.Remove(300)
  if bst.Contains(300) == true {
    t.Errorf("bst.Contains(300) should be false")
  }
  bst.Remove(2)
  if bst.Contains(2) == true {
    t.Errorf("bst.Contains(2) should be false")
  }
}

func TestMax(t *testing.T) {
  bst := new(BST)

  bst.Insert(1)
  bst.Insert(300)
  bst.Insert(18)
  bst.Insert(45)
  bst.Insert(2)

  if bst.Max() != 300 {
    t.Errorf("bst.Max() should be 300")
  }
}

func TestMin(t *testing.T) {
  bst := new(BST)

  bst.Insert(1)
  bst.Insert(300)
  bst.Insert(18)
  bst.Insert(45)
  bst.Insert(2)

  if bst.Min() != 1 {
    t.Errorf("bst.Min() should be 1")
  }
}

func TestPreOrder(t *testing.T) {
  bst := new(BST)
  items := []int{}
  bst.Insert(2)
  bst.Insert(1)
  bst.Insert(3)

  bst.PreOrder(func (x int) { items = append(items, x) })

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
  bst := new(BST)
  items := []int{}
  bst.Insert(2)
  bst.Insert(1)
  bst.Insert(3)

  bst.PostOrder(func (x int) { items = append(items, x) })

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
  bst := new(BST)
  items := []int{}
  bst.Insert(2)
  bst.Insert(1)
  bst.Insert(3)

  bst.InOrder(func (x int) { items = append(items, x) })

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
