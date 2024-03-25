package stack

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	value *T
	prev  *node[T]
}

type Stack[T any] struct {
	length uint
	tail   *node[T]
}

func newStack[T any]() *Stack[T] {
	return &Stack[T]{
		length: 0,
		tail:   nil,
	}
}

func (s *Stack[T]) push(item T) {
	node := node[T]{
		value: &item,
		prev:  nil,
	}

	s.length++

	if s.tail != nil {
		prev := s.tail
		s.tail = &node
		node.prev = prev
	} else {
		s.tail = &node
	}

}

func (s *Stack[T]) pop() (*T, error) {
	if s.tail == nil {
		return nil, errors.New("stack is empty")
	}

	s.length--
	item := s.tail.value
	if s.length == 0 {
		s.tail = nil

		return item, nil
	} else {
		s.tail = s.tail.prev
		return item, nil
	}
}

func (s *Stack[T]) toSlice() []T {
	if s.length == 0 {
		return []T{}
	}

	slice := make([]T, s.length)

	for i := len(slice) - 1; i >= 0; i-- {
		item, err := s.pop()
		if err == nil {
			slice[i] = *item
		}
	}

	return slice
}

func (s *Stack[T]) peek() {
	fmt.Println(*s.tail.value)
}

func (s *Stack[T]) stackLength() uint {
	return s.length
}

func (s *Stack[T]) print() {
	fmt.Println("===> stack ", s.toSlice())

}

func UseStack() {
	var text = []string{"hello", "world", "from", "the", "computer"}
	s := newStack[string]()

	for _, v := range text {
		s.push(v)
	}

	s.peek()

	fmt.Println("stack length ", s.stackLength())

	fmt.Println(s)

	s.peek()
	s.peek()

	_, _ = s.pop()

	s.push("TV")

	s.peek()

	s.print()
}
