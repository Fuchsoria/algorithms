package tree

import (
	"math/rand"
)

type TreapTree struct {
	root *TreapNode
}

type TreapNode struct {
	Value       int
	Priority    int
	Left, Right *TreapNode
}

func (n *TreapNode) ForEach(fn func(v int)) {
	if n != nil {
		n.Left.ForEach(fn)
		fn(n.Value)
		n.Right.ForEach(fn)
	}
}

func (n *TreapNode) Merge(other *TreapNode, overwrite bool) *TreapNode {
	if n == nil {
		return other
	}

	if other == nil {
		return n
	}

	if n.Priority < other.Priority {
		other, n, overwrite = n, other, !overwrite
	}

	left, mid, right := other.Split(n.Value)
	value := n.Value
	if overwrite && mid != nil {
		value = mid.Value
	}

	left = n.Left.Merge(left, overwrite)
	right = n.Right.Merge(right, overwrite)

	return &TreapNode{value, n.Priority, left, right}
}

func (n *TreapNode) Split(v int) (left, mid, right *TreapNode) {
	leftp, rightp := &left, &right
	for {
		if n == nil {
			*leftp = nil
			*rightp = nil

			return left, nil, right
		}

		root := &TreapNode{n.Value, n.Priority, nil, nil}

		switch {
		case n.Value < v:
			*leftp = root
			root.Left = n.Left
			leftp = &root.Right
			n = n.Right
		case n.Value > v:
			*rightp = root
			root.Right = n.Right
			rightp = &root.Left
			n = n.Left
		default:
			*leftp = n.Left
			*rightp = n.Right

			return left, root, right
		}
	}
}

func (n *TreapNode) join(other *TreapNode) *TreapNode {
	var result *TreapNode
	resultp := &result

	for {
		if n == nil {
			*resultp = other
			return result
		}
		if other == nil {
			*resultp = n
			return result
		}

		if n.Priority <= other.Priority {
			root := &TreapNode{n.Value, n.Priority, n.Left, nil}
			*resultp = root
			resultp = &root.Right
			n = n.Right
		} else {
			root := &TreapNode{other.Value, other.Priority, nil, other.Right}
			*resultp = root
			resultp = &root.Left
			other = other.Left
		}
	}
}

func (n *TreapNode) search(v int) (*TreapNode, error) {
	for {
		if n == nil {
			return nil, ErrNodeNotFound
		}

		switch {
		case n.Value == v:
			return n, nil
		case n.Value < v:
			n = n.Right
		case n.Value > v:
			n = n.Left
		}
	}
}

func (n *TreapNode) remove(v int) *TreapNode {
	left, _, right := n.Split(v)

	return left.join(right)
}

func (t *TreapTree) Remove(v int) {
	t.root = t.root.remove(v)
}

func (t *TreapTree) Search(v int) (*TreapNode, error) {
	return t.root.search(v)
}

func (t *TreapTree) Insert(v int) {
	priority := rand.Intn(10000000)

	t.root = t.root.Merge(&TreapNode{v, priority, nil, nil}, false)
}

func (t *TreapTree) Inorder() []int {
	arr := []int{}

	if t.root == nil {
		return arr
	}

	t.root.ForEach(func(v int) {
		arr = append(arr, v)
	})

	return arr
}
