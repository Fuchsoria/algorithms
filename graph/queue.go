package graph

import (
	"container/list"
)

type Queue[T any | *any] struct {
	list *list.List
}

func (q *Queue[T]) Push(item T) {
	q.list.PushBack(item)
}

func (q *Queue[T]) emptyValue() T {
	var empty T

	return empty
}

func (q *Queue[T]) Shift() T {
	if q.Len() <= 0 {
		return q.emptyValue()
	}

	item := q.list.Front()
	value := item.Value.(T)

	q.list.Remove(item)

	return value
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		list.New(),
	}
}
