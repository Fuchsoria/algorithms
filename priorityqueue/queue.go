package priorityqueue

type Node struct {
	prev  *Node
	next  *Node
	Value int64
}

func (n *Node) hasPrev() bool {
	if n.prev != nil {
		return true
	}

	return false
}

func (n *Node) hasNext() bool {
	if n.next != nil {
		return true
	}

	return false
}

type Queue struct {
	first *Node
	last  *Node
}

func (q *Queue) hasFirst() bool {
	if q.first != nil {
		return true
	}

	return false
}

func (q *Queue) hasLast() bool {
	if q.last != nil {
		return true
	}

	return false
}

func (q *Queue) Add(value int64) {
	if q.hasLast() {
		newItem := &Node{
			prev:  q.last,
			Value: value,
		}

		q.last.next = newItem
		q.last = newItem
	} else {
		newItem := &Node{
			Value: value,
		}

		q.first = newItem
		q.last = newItem
	}
}

func (q *Queue) Take() *Node {
	if q.hasFirst() {
		value := q.first

		if q.first.hasNext() {
			q.first = q.first.next
		} else {
			q.first = nil
			q.last = nil
		}

		return value
	}

	return nil
}

func (q *Queue) GetValues() []int64 {
	values := []int64{}
	node := q.first

	for node != nil {
		values = append(values, node.Value)

		node = node.next
	}

	return values
}
