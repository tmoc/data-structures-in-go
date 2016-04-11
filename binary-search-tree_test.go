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
