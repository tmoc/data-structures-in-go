package ds

type node struct {
  value int
  left *node
  right *node
}

type BST struct {
  root *node
  count int
}

func addNode(n *node, value int) *node {
  if n == nil {
    n = new(node)
    n.value = value
    return n
  } else if value < n.value {
    n.left = addNode(n.left, value)
  } else {
    n.right = addNode(n.right, value)
  }
  return n
}

func (t *BST) Insert(value int) {
  t.root = addNode(t.root, value)
  t.count++
}

func getNode(n *node, value int) *node {
  if n == nil {
    return nil
  } else if value == n.value {
    return n
  } else if value < n.value {
    return getNode(n.left, value)
  } else {
    return getNode(n.right, value)
  }
}

func (t *BST) Contains(value int) bool {
  if getNode(t.root, value) == nil {
    return false
  }
  return true
}

func getParentNode(n *node, value int) *node {
  if value < n.value {
    if value == n.left.value {
      return n
    } else {
      return getParentNode(n.left, value)
    }
  } else {
    if value == n.right.value {
      return n
    } else {
      return getParentNode(n.right, value)
    }
  }
}

func getParent(t *BST, value int) *node {
  if value == t.root.value || t.count < 2 {
    return nil
  }
  return getParentNode(t.root, value)
}

func (t *BST) Remove(value int) bool {
  nodeToRemove := getNode(t.root, value)

  if nodeToRemove == nil {
    return false
  }

  if t.count == 1 {
    t.root = nil
    t.count = 0
    return true
  }

  parentNode := getParent(t, value)

  // Covers edge case; removing the root and there is more than 1 node.
  if parentNode == nil {
    if t.root.left == nil && t.root.right != nil {
      t.root = t.root.right
    } else if t.root.left != nil && t.root.right == nil {
      t.root = t.root.left
    } else {
      largestValueOnLeft := t.root.left
      for largestValueOnLeft.right != nil {
        largestValueOnLeft = largestValueOnLeft.right
      }
      largetValueOnLeftParent := getParent(t, largestValueOnLeft.value)
      t.root.value = largestValueOnLeft.value
      largetValueOnLeftParent.right = nil
    }
    t.count--
    return true
  }

  if nodeToRemove.left == nil && nodeToRemove.right == nil {
    if nodeToRemove.value < parentNode.value {
      parentNode.left = nil
    } else {
      parentNode.right = nil
    }
  } else if nodeToRemove.left == nil && nodeToRemove.right != nil {
    if nodeToRemove.value < parentNode.value {
      parentNode.left = nodeToRemove.right
    } else {
      parentNode.right = nodeToRemove.right
    }
  } else if nodeToRemove.left != nil && nodeToRemove.right == nil {
    if nodeToRemove.value < parentNode.value {
      parentNode.left = nodeToRemove.left
    } else {
      parentNode.right = nodeToRemove.left
    }
  } else {
    largestValueNode := nodeToRemove.left

    for largestValueNode.right != nil {
      largestValueNode = largestValueNode.right
    }
    largestValueParent := getParentNode(nodeToRemove, largestValueNode.value)
    nodeToRemove.value = largestValueNode.value
    largestValueParent.right = nil
  }
  t.count--
  return true
}

func findMaxNode(n *node) int {
  if n.right == nil {
    return n.value
  } else {
    return findMaxNode(n.right)
  }
}

func (t *BST) Max() int {
  return findMaxNode(t.root)
}

func findMinNode(n *node) int {
  if n.left == nil {
    return n.value
  } else {
    return findMinNode(n.left)
  }
}

func (t *BST) Min() int {
  return findMinNode(t.root)
}

// Use to process data.
type process func(int)

func preOrderNodeTraversal(n *node, fn process) {
  fn(n.value)

  if n.left != nil {
    preOrderNodeTraversal(n.left, fn)
  }
  if n.right != nil {
    preOrderNodeTraversal(n.right, fn)
  }
}

func (t *BST) PreOrder(fn process) {
  if t.root != nil {
    preOrderNodeTraversal(t.root, fn)
  }
}

func postOrderTraversal(n *node, fn process) {
  if n.left != nil {
    postOrderTraversal(n.left, fn)
  }
  if n.right != nil {
    postOrderTraversal(n.right, fn)
  }
  fn(n.value)
}

func (t *BST) PostOrder(fn process) {
  if t.root != nil {
    postOrderTraversal(t.root, fn)
  }
}

func inOrderTraversal(n *node, fn process) {
  if n.left != nil {
    inOrderTraversal(n.left, fn)
  }
  fn(n.value)
  if n.right != nil {
    inOrderTraversal(n.right, fn)
  }
}

func (t *BST) InOrder(fn process) {
  if t.root != nil {
    inOrderTraversal(t.root, fn)
  }
}
