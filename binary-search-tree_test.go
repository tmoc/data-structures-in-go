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
    t.Errorf("bst.Contains(18) should equal true")
  }

  bst.Insert(2)

  if bst.Contains(2) != true {
    t.Errorf("bst.Contains(2) should equal true")
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
    t.Errorf("bst.Contains(300) should equal true")
  }
  if bst.Contains(2) != true {
    t.Errorf("bst.Contains(2) should equal true")
  }

  bst.Remove(300)
  if bst.Contains(300) == true {
    t.Errorf("bst.Contains(300) should equal false")
  }
  bst.Remove(2)
  if bst.Contains(2) == true {
    t.Errorf("bst.Contains(2) should equal false")
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
    t.Errorf("bst.Max() should equal 300")
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
    t.Errorf("bst.Min() should equal 1")
  }
}

func TestPreOrder(t *testing.T) {
  bst := new(BST)
  items := []int{}
  bst.Insert(3)
  bst.Insert(1)
  bst.Insert(2)
  
  bst.PreOrder(func (x int) { items = append(items, x) })

  if items[0] != 3 {
    t.Errorf("items[0] should equal 3")
  }
  if items[1] != 1 {
    t.Errorf("items[1] should equal 1")
  }
  if items[2] != 2 {
    t.Errorf("items[2] should equal 2")
  }
}
