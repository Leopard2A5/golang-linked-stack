package linked_stack

import "errors"

var ErrEmpty = errors.New("Stack is empty")

type StackItem[T any] struct {
	next *StackItem[T]
	value T
}

type LinkedStack[T any] struct {
	next *StackItem[T]
}

func NewStack[T any]() LinkedStack[T] {
	return LinkedStack[T]{}
}

func Len[T any](stack *LinkedStack[T])(len int64) {
	for item := stack.next; item != nil; item = item.next {
		len += 1
	}
	return
}

func Push[T any](stack *LinkedStack[T], value T) {
	newItem := StackItem[T]{ next: stack.next, value: value }
	stack.next = &newItem
}

func Pop[T any](stack *LinkedStack[T]) (T, error) {
	if stack.next == nil {
		var ret T
		return ret, ErrEmpty
	}

	var ret = stack.next.value
	stack.next = stack.next.next

	return ret, nil
}
