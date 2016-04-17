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

func (t *RedBlackTree) rightRotation(n *node) {
	oldLeft := n.left
	n.left = oldLeft.right
	if oldLeft.right != nil {
		oldLeft.right.parent = n
	}
	oldLeft.parent = n.parent
	if n.parent == nil {
		t.root = n
	} else if n == n.parent.left {
		n.parent.left = oldLeft
	} else {
		n.parent.right = oldLeft
	}
	oldLeft.right = n
	n.parent = oldLeft
}

func (t *RedBlackTree) leftRotation(n *node) {
	oldRight := n.right
	n.right = oldRight.left
	if oldRight.left != nil {
		oldRight.left.parent = n
	}
	oldRight.parent = n.parent
	if n.parent == nil {
		t.root = oldRight
	} else if n == n.parent.left {
		n.parent.left = oldRight
	} else {
		n.parent.right = oldRight
	}
	oldRight.left = n
	n.parent = oldRight
}

func (t *RedBlackTree) fixRedRed(n *node) {
	for n.parent != nil && n.parent.color == RED {
		if n.parent == n.parent.parent.left {
			uncle := n.parent.parent.right
			if uncle != nil && uncle.color == RED {
				uncle.color = BLACK
				n.parent.color = BLACK
				n.parent.parent.color = RED
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					t.leftRotation(n)
				}
				n.parent.color = BLACK
				n.parent.parent.color = RED
				t.rightRotation(n.parent.parent)
			}
		} else {
			uncle := n.parent.parent.left
			if uncle != nil && uncle.color == RED {
				uncle.color = BLACK
				n.parent.color = BLACK
				n.parent.parent.color = RED
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rightRotation(n)
				}
				n.parent.color = BLACK
				n.parent.parent.color = RED
				t.leftRotation(n.parent.parent)
			}
		}
	}
}

func (t *RedBlackTree) Insert(v int) {
	if t.root == nil {
		t.root = &node{value: v}
		t.count = 1
		return
	}

	var parent *node
	willBeNil := t.root

	for willBeNil != nil {
		parent = willBeNil
		if v < willBeNil.value {
			willBeNil = willBeNil.left
		} else {
			willBeNil = willBeNil.right
		}
	}

	child := &node{value: v, parent: parent, color: RED}

	if v < parent.value {
		parent.left = child
	} else {
		parent.right = child
	}

	if parent.color == RED {
		t.fixRedRed(child)
		t.root.color = BLACK
	}

	t.count++
}
