package avltree

type node struct {
	value  int
	height int
	left   *node
	right  *node
}

type AVLTree struct {
	root  *node
	count int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(n *node) int {
	if n != nil {
		return n.height
	}
	return -1
}

func leftRotation(n *node) *node {
	oldRight := n.right
	n.right = oldRight.left
	oldRight.left = n
	n.height = max(height(n.left), height(n.right)) + 1
	oldRight.height = max(height(oldRight.left), height(oldRight.right)) + 1
	return oldRight
}

func leftAndRightRotation(n *node) *node {
	n.left = leftRotation(n.left)
	n = rightRotation(n)
	return n
}

func rightRotation(n *node) *node {
	oldLeft := n.left
	n.left = oldLeft.right
	oldLeft.right = n
	n.height = max(height(n.left), height(n.right)) + 1
	oldLeft.height = max(height(oldLeft.left), height(oldLeft.right)) + 1
	return oldLeft
}

func rightAndLeftRotation(n *node) *node {
	n.right = rightRotation(n.right)
	n = leftRotation(n)
	return n
}

func checkBalance(n *node) *node {
	if height(n.left)-height(n.right) == 2 { // Left too high, must rebalance.
		if height(n.left.left)-height(n.left.right) > 0 { // Left balance-factor subtree.
			n = rightRotation(n)
		} else {
			n = leftAndRightRotation(n)
		}
	} else if height(n.left)-height(n.right) == -2 { // Right too high, must rebalance.
		if height(n.right.left)-height(n.right.right) < 0 { // Right balance-factor subtree.
			n = leftRotation(n)
		} else {
			n = rightAndLeftRotation(n)
		}
	} else {
		n.height = max(height(n.left), height(n.right)) + 1 // Commenting this out simply means no rebalancing happens. Heights are always 0.
	}
	return n
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
	return checkBalance(n)
}

func (t *AVLTree) Insert(value int) {
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

func (t *AVLTree) Contains(value int) bool {
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

func getParent(t *AVLTree, value int) *node {
	if value == t.root.value || t.count < 2 {
		return nil
	}
	return getParentNode(t.root, value)
}

func (t *AVLTree) Remove(value int) bool {
	// Stack of visited nodes.
	pathStack := []*node{}

	nodeToRemove := t.root

	for nodeToRemove != nil && value != nodeToRemove.value {
		pathStack = append(pathStack, nodeToRemove)

		if value < nodeToRemove.value {
			nodeToRemove = nodeToRemove.left
		} else {
			nodeToRemove = nodeToRemove.right
		}
	}

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
			largestOnLeft := t.root.left
			parentIsNodeToRemove := true

			for largestOnLeft.right != nil {
				largestOnLeft = largestOnLeft.right
				parentIsNodeToRemove = false
			}

			if parentIsNodeToRemove == true {
				t.root.value = t.root.left.value
				t.root.left = nil
			} else {
				largestOnLeftParent := getParent(t, largestOnLeft.value)
				t.root.value = largestOnLeft.value

				if largestOnLeft.left != nil {
					largestOnLeftParent.right = largestOnLeft.left
				} else {
					largestOnLeftParent.right = nil
				}
			}
		}
		t.root = checkBalance(t.root)
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
		largestOnLeft := nodeToRemove.left
		parentIsNodeToRemove := true

		for largestOnLeft.right != nil {
			largestOnLeft = largestOnLeft.right
			parentIsNodeToRemove = false
		}

		if parentIsNodeToRemove == true {
			nodeToRemove.value = nodeToRemove.left.value
			nodeToRemove.left = nil
		} else {
			largestOnLeftParent := getParentNode(nodeToRemove, largestOnLeft.value)
			nodeToRemove.value = largestOnLeft.value

			if largestOnLeft.left != nil {
				largestOnLeftParent.right = largestOnLeft.left
			} else {
				largestOnLeftParent.right = nil
			}
		}
	}

	// Rebalance nodes on path.
	for i := len(pathStack) - 1; i > -1; i-- {
		pathStack[i] = checkBalance(pathStack[i])
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

func (t *AVLTree) Max() int {
	return findMaxNode(t.root)
}

func findMinNode(n *node) int {
	if n.left == nil {
		return n.value
	} else {
		return findMinNode(n.left)
	}
}

func (t *AVLTree) Min() int {
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

func (t *AVLTree) PreOrder(fn process) {
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

func (t *AVLTree) PostOrder(fn process) {
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

func (t *AVLTree) InOrder(fn process) {
	if t.root != nil {
		inOrderTraversal(t.root, fn)
	}
}

func breadthFirstTraversal(n *node, fn process) {
	queue := []*node{}
	current := n

	for current != nil {
		fn(current.value)

		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}

		length := len(queue)
		if length != 0 {
			current = queue[0]
			queue = queue[1:]
		} else {
			current = nil
		}
	}
}

func (t *AVLTree) BreadthFirst(fn process) {
	if t.root != nil {
		breadthFirstTraversal(t.root, fn)
	}
}
