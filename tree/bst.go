package tree

import (
	"errors"
)

var ErrNodeNotFound = errors.New("couldn't find node")

type bstNode struct {
	Left  *bstNode
	Right *bstNode
	Value int
}

type bstTree struct {
	root *bstNode
	arr  []int
}

func (b *bstTree) Reset() {
	b.root = nil
}

func (b *bstTree) insert(node *bstNode, v int) *bstNode {
	if b.root == nil {
		b.root = &bstNode{Value: v}

		return b.root
	}

	if node == nil {
		return &bstNode{Value: v}
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

func (b *bstTree) Insert(v int) {
	b.insert(b.root, v)
}

func (b *bstTree) search(node *bstNode, v int) *bstNode {
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

func (b *bstTree) Search(v int) (*bstNode, error) {
	n := b.search(b.root, v)
	if n == nil {
		return nil, ErrNodeNotFound
	}

	return n, nil
}

func (b *bstTree) inorder(node *bstNode) {
	if node != nil {
		b.inorder(node.Left)
		b.arr = append(b.arr, node.Value)
		b.inorder(node.Right)
	}
}

func (b *bstTree) Inorder() []int {
	b.arr = []int{}

	b.inorder(b.root)

	return b.arr
}

func (b *bstTree) remove(root *bstNode, value int) *bstNode {
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

func (b *bstTree) Remove(v int) {
	b.remove(b.root, v)
}

func (b *bstTree) inOrderSuccessor(root *bstNode) int {
	minimum := root.Value

	for root.Left != nil {
		minimum = root.Left.Value
		root = root.Left
	}

	return minimum
}
