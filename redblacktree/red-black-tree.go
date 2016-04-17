package redblacktree

const (
	BLACK = 0
	RED   = 1
)

type node struct {
	value, color        int
	parent, left, right *node
}

type RedBlackTree struct {
	root  *node
	count int
}

func getNode(n *node, v int) *node {
	if n == nil {
		return nil
	} else if v == n.value {
		return n
	} else if v < n.value {
		return getNode(n.left, v)
	} else {
		return getNode(n.right, v)
	}
}

func (t *RedBlackTree) Contains(v int) bool {
	if getNode(t.root, v) != nil {
		return true
	}
	return false
}

func rightRotation(n *node) {
	oldLeft := n.left
	n.left = oldLeft.right
	if n.left != nil {
		n.left.parent = n
	}
	oldLeft.right = n
	oldLeft.right.parent = oldLeft
}

func leftRotation(n *node) {
	oldRight := n.right
	n.right = oldRight.left
	if n.right != nil {
		n.right.parent = n
	}
	oldRight.left = n
	oldRight.left.parent = oldRight
}

func fixViolation(parent *node) {
	if parent.color == BLACK {
		// Nothing to fix.
		return
	}

	leftParent := parent.value < parent.parent.value
	leftNew := parent.left != nil && parent.left.color == RED

	if leftParent == true { // Right rotation.
		if parent.parent.right.color == RED {
			parent.parent.color = RED
			parent.parent.left.color = BLACK
			parent.parent.right.color = BLACK
		} else {
			if leftNew == true {
				rightRotation(parent.parent)
			} else {
				leftRotation(parent)
				rightRotation(parent.parent)
			}
		}
	} else { // Left rotation.
		if parent.parent.left != nil && parent.parent.left.color == RED {
			parent.parent.color = RED
			parent.parent.left.color = BLACK
			parent.parent.right.color = BLACK
		} else {
			if leftNew != true {
				leftRotation(parent.parent)
			} else {
				rightRotation(parent)
				leftRotation(parent.parent)
			}
		}
	}
}

func fixRedRedRight(parent *node) {
	if parent.right.right && parent.right.right == RED {

	}
}

func addNode(n *node, v int) {
	if v < n.value {
		if n.left == nil {
			n.left = &node{value: v, parent: n, color: RED}
		} else {
			addNode(n.left, v)
			if n.color == RED && n.left.color == RED {
				fixRedRedLeft(n)
			}
		}
	} else {
		if n.right == nil {
			n.right = &node{value: v, parent: n, color: RED}
		} else {
			addNode(n.right, v)
			if n.color == RED && n.right.color == RED {
				fixRedRedRight(n)
			}
		}
	}
}

func (t *RedBlackTree) Insert(v int) {
	if t.root == nil {
		t.root = &node{value: v, color: BLACK}
	} else {
		addNode(t.root, v)
		t.root.parent = nil
		t.root.color = BLACK
	}
	t.count++
}
