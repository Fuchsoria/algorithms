package priorityqueue

import (
	"errors"
	"fmt"
	"sort"
)

var errorEmpty = errors.New("error empty queue")

type PrioritizedValue struct {
	Priority int64
	Values   []int64
}

type PriorityNode struct {
	prev     *PriorityNode
	next     *PriorityNode
	Priority int64
	Queue    Queue
}

func (n *PriorityNode) hasPrev() bool {
	if n.prev != nil {
		return true
	}

	return false
}

func (n *PriorityNode) hasNext() bool {
	if n.next != nil {
		return true
	}

	return false
}

type PriorityQueue struct {
	low    *PriorityNode
	high   *PriorityNode
	length int64
}

func (q *PriorityQueue) hasLow() bool {
	if q.low != nil {
		return true
	}

	return false
}

func (q *PriorityQueue) hasHigh() bool {
	if q.high != nil {
		return true
	}

	return false
}

func (q *PriorityQueue) increaseLen() {
	q.length++
}

func (q *PriorityQueue) decreaseLen() {
	if q.length > 0 {
		q.length--
	}
}

func (q *PriorityQueue) createQueueSorted(priority int64) *PriorityNode {
	defer q.increaseLen()

	// Edge case if only one priority
	if q.length == 1 {
		node := q.low

		// Check node priority is equal or greater than arg priority
		if node.Priority >= priority {
			// Create node with point to next node
			newNode := &PriorityNode{Priority: priority, next: node}

			// Rewrite points
			node.prev = newNode
			q.low = newNode
			q.high = node

			return newNode
		}

		// Create node with point to prev node
		newNode := &PriorityNode{Priority: priority, prev: node}

		// Rewrite points
		node.next = newNode
		q.low = node
		q.high = newNode

		return newNode
	}

	if q.hasLow() {
		node := q.low

		// Loop through nodes
		for node != nil {
			if !node.hasNext() {
				// 10, 20, 30, 40, 50 => 10, 20, 30, 40, 50, >60<
				newNode := &PriorityNode{Priority: priority, prev: node}

				node.next = newNode
				q.high = newNode

				return newNode
			}

			if node.Priority >= priority {
				if !node.hasPrev() {
					// 10, 20, 30, 40, 50 => >5<, 10, 20, 30, 40, 50
					newNode := &PriorityNode{Priority: priority, next: node}

					node.prev = newNode
					q.low = newNode

					return newNode
				}

				// 10, 20, 30, 40, 50 => 10, 20, 30, >35<, 40, 50
				newNode := &PriorityNode{Priority: priority, prev: node.prev, next: node}

				node.prev.next = newNode

				return newNode
			}

			node = node.next
		}
	}

	// empty => >10<
	newNode := &PriorityNode{Priority: priority}

	q.low = newNode
	q.high = newNode

	return newNode
}

func (q *PriorityQueue) getQueueByPriority(priority int64) *PriorityNode {
	if q.hasLow() {
		node := q.low

		for node != nil {
			if node.Priority == priority {
				return node
			}

			node = node.next
		}
	}

	// Fallback to create priority
	return q.createQueueSorted(priority)
}

func (q *PriorityQueue) getHighPriorityQueue() *Node {
	if q.hasHigh() {
		node := q.high

		// Loop through nodes
		for node != nil {
			// Check if queue has items
			if node.Queue.hasFirst() {
				// Take item from queue
				value := q.high.Queue.Take()

				// Check queue to remove it if it's empty
				if !node.Queue.hasFirst() {
					// Edge case for last item
					if !node.hasPrev() {
						q.high = nil
						q.low = nil
					} else {
						node.prev.next = nil
						q.high = node.prev
					}

					q.decreaseLen()
				}

				return value
			}

			node = node.prev
		}
	}

	return nil
}

func (q *PriorityQueue) Len() int64 {
	return q.length
}

func (q *PriorityQueue) Enqueue(priority int64, value int64) {
	queue := q.getQueueByPriority(priority)

	queue.Queue.Add(value)
}

func (q *PriorityQueue) Dequeue() (int64, error) {
	node := q.getHighPriorityQueue()
	if node != nil {
		return node.Value, nil
	}

	return 0, errorEmpty
}

func (q *PriorityQueue) GetPriorities() []int64 {
	values := []int64{}
	node := q.low

	for node != nil {
		values = append(values, node.Priority)

		node = node.next
	}

	return values
}

func (q *PriorityQueue) GetPrioritizedValues() []PrioritizedValue {
	values := []PrioritizedValue{}

	node := q.low

	for node != nil {
		values = append(values, PrioritizedValue{Priority: node.Priority, Values: node.Queue.GetValues()})

		node = node.next
	}

	return values
}

type PriorityQueueMapped struct {
	high int64
	data map[int64]*Queue
}

func (q *PriorityQueueMapped) removeAndReCalculateHigh() int64 {
	var newHigh int64 = 0

	for k, v := range q.data {
		if !v.hasFirst() {
			delete(q.data, k)
		} else {
			if k > newHigh {
				newHigh = k
			}
		}
	}

	q.high = newHigh

	return newHigh
}

func (q *PriorityQueueMapped) isHighestEmpty() bool {
	queue, ok := q.data[q.high]
	if !ok {
		return true
	}

	if !queue.hasFirst() {
		return true
	}

	return false
}

func (q *PriorityQueueMapped) validateEmpty() {
	if q.isHighestEmpty() {
		q.removeAndReCalculateHigh()
	}
}

func (q *PriorityQueueMapped) Enqueue(priority int64, value int64) {
	if q.data == nil {
		q.data = make(map[int64]*Queue)
	}

	if queue, ok := q.data[priority]; ok {
		queue.Add(value)
	} else {
		queue := Queue{}
		queue.Add(value)

		q.data[priority] = &queue
	}

	if priority > q.high {
		q.high = priority
	}
}

func (q *PriorityQueueMapped) Dequeue() (int64, error) {
	defer func() {
		q.validateEmpty()
	}()

	if q.data == nil {
		return 0, fmt.Errorf("data == nil: %w", errorEmpty)
	}

	if queue, ok := q.data[q.high]; ok {
		if queue.hasFirst() {
			node := queue.Take()

			return node.Value, nil
		} else {
			newHigh := q.removeAndReCalculateHigh()

			newHighQueue, newHighOK := q.data[newHigh]
			if newHighOK && newHighQueue.hasFirst() {
				node := queue.Take()

				return node.Value, nil
			}
		}
	}

	return 0, fmt.Errorf("zero condition: %w", errorEmpty)
}

func (q *PriorityQueueMapped) GetPrioritizedValues() []PrioritizedValue {
	values := []PrioritizedValue{}

	for k, v := range q.data {
		values = append(values, PrioritizedValue{Priority: k, Values: v.GetValues()})
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].Priority < values[j].Priority
	})

	return values
}
