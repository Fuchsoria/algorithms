package barn

import (
	"container/list"
)

type Stack[T any] struct {
	list *list.List
}

func (s *Stack[T]) Push(item T) {
	s.list.PushBack(item)
}

func (s *Stack[T]) emptyValue() T {
	var empty T

	return empty
}

func (s *Stack[T]) Pop() T {
	if s.Len() <= 0 {
		return s.emptyValue()
	}

	item := s.list.Back()
	value := item.Value.(T)

	s.list.Remove(item)

	return value
}

func (s *Stack[T]) Len() int {
	return s.list.Len()
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{
		list.New(),
	}
}
