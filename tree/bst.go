package tree

import (
	"errors"
)

var ErrNodeNotFound = errors.New("couldn't find node")

type bsNode struct {
	Left  *bsNode
	Right *bsNode
	Value int
}

type bsTree struct {
	root *bsNode
	arr  []int
}

func (b *bsTree) Reset() {
	b.root = nil
}

func (b *bsTree) insert(node *bsNode, v int) *bsNode {
	if b.root == nil {
		b.root = &bsNode{Value: v}

		return b.root
	}

	if node == nil {
		return &bsNode{Value: v}
	}

	// UNIQUE VALUE TREE
	if v == node.Value {
		return node
	}

	if v <= node.Value {
		node.Left = b.insert(node.Left, v)
	}

	if v >= node.Value {
		node.Right = b.insert(node.Right, v)
	}

	return node
}

func (b *bsTree) Insert(v int) {
	b.insert(b.root, v)
}

func (b *bsTree) search(node *bsNode, v int) *bsNode {
	if node == nil {
		return nil
	}

	if node.Value == v {
		return node
	}

	if v < node.Value {
		return b.search(node.Left, v)
	}

	return b.search(node.Right, v)
}

func (b *bsTree) Search(v int) (*bsNode, error) {
	n := b.search(b.root, v)
	if n == nil {
		return nil, ErrNodeNotFound
	}

	return n, nil
}

func (b *bsTree) inorder(node *bsNode) {
	if node != nil {
		b.inorder(node.Left)
		b.arr = append(b.arr, node.Value)
		b.inorder(node.Right)
	}
}

func (b *bsTree) Inorder() []int {
	b.arr = []int{}

	b.inorder(b.root)

	return b.arr
}

func (b *bsTree) remove(root *bsNode, value int) *bsNode {
	if root == nil {
		return root
	}

	if value < root.Value {
		root.Left = b.remove(root.Left, value)
	} else if value > root.Value {
		root.Right = b.remove(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Value = b.inOrderSuccessor(root.Right)
		root.Right = b.remove(root.Right, root.Value)
	}

	return root
}

func (b *bsTree) Remove(v int) {
	b.remove(b.root, v)
}

func (b *bsTree) inOrderSuccessor(root *bsNode) int {
	minimum := root.Value

	for root.Left != nil {
		minimum = root.Left.Value
		root = root.Left
	}

	return minimum
}
