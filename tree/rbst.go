package tree

import "math/rand"

type rbsNode struct {
	Left  *rbsNode
	Right *rbsNode
	Value int
	Size  int
}

type rbsTree struct {
	root *rbsNode
}

func (b *rbsTree) Reset() {
	b.root = nil
}

func (b *rbsTree) rotateRight(p *rbsNode) *rbsNode {
	q := p.Left
	if q == nil {
		return p
	}

	p.Left = q.Right
	q.Right = p
	q.Size = p.Size

	b.fixSize(p)

	return q
}

func (b *rbsTree) rotateLeft(q *rbsNode) *rbsNode {
	p := q.Right
	if p == nil {
		return q
	}

	q.Right = p.Left
	p.Left = q
	p.Size = q.Size

	b.fixSize(q)

	return p
}

func (b *rbsTree) getSize(node *rbsNode) int {
	if node == nil {
		return 0
	}

	return node.Size
}

func (b *rbsTree) fixSize(node *rbsNode) {
	node.Size = b.getSize(node.Left) + b.getSize(node.Right) + 1
}

func (b *rbsTree) insertRoot(p *rbsNode, v int) *rbsNode {
	if p == nil {
		return &rbsNode{Value: v, Size: 1}
	}

	if v < p.Value {
		p.Left = b.insertRoot(p.Left, v)

		return b.rotateRight(p)
	} else {
		p.Right = b.insertRoot(p.Right, v)

		return b.rotateLeft(p)
	}
}

func (b *rbsTree) insert(p *rbsNode, v int) *rbsNode {
	if p == nil {
		return &rbsNode{Value: v, Size: 1}
	}

	if rand.Int()%(p.Size+1) == 0 {
		return b.insertRoot(p, v)
	}

	if p.Value > v {
		p.Left = b.insert(p.Left, v)
	} else {
		p.Right = b.insert(p.Right, v)
	}

	b.fixSize(p)

	return p
}

func (b *rbsTree) Insert(v int) {
	b.root = b.insert(b.root, v)
}

func (b *rbsTree) search(node *rbsNode, v int) *rbsNode {
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

func (b *rbsTree) Search(v int) (*rbsNode, error) {
	n := b.search(b.root, v)
	if n == nil {
		return nil, ErrNodeNotFound
	}

	return n, nil
}

func (b *rbsTree) inorder(node *rbsNode, fn func(v int)) {
	if node != nil {
		b.inorder(node.Left, fn)
		fn(node.Value)
		b.inorder(node.Right, fn)
	}
}

func (b *rbsTree) Inorder() []int {
	arr := []int{}

	b.inorder(b.root, func(v int) {
		arr = append(arr, v)
	})

	return arr
}

func (b *rbsTree) join(p, q *rbsNode) *rbsNode {
	if p == nil {
		return q
	}

	if q == nil {
		return p
	}

	if rand.Int()%(p.Size+q.Size) < p.Size {
		p.Right = b.join(p.Right, q)
		b.fixSize(p)
		return p
	} else {
		q.Left = b.join(p, q.Left)
		b.fixSize(q)
		return q
	}
}

func (b *rbsTree) remove(p *rbsNode, v int) *rbsNode {
	if p == nil {
		return p
	}

	if p.Value == v {
		q := b.join(p.Left, p.Right)

		return q
	} else if v < p.Value {
		p.Left = b.remove(p.Left, v)
	} else {
		p.Right = b.remove(p.Right, v)
	}

	return p
}

func (b *rbsTree) Remove(v int) {
	b.root = b.remove(b.root, v)
}

func (b *rbsTree) inOrderSuccessor(root *rbsNode) int {
	minimum := root.Value

	for root.Left != nil {
		minimum = root.Left.Value
		root = root.Left
	}

	return minimum
}
