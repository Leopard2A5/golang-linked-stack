package linked_stack

import "testing"
import "errors"

func TestLen(t *testing.T) {
	stack := NewStack[int64]()

	testLen(&stack, 0, t)

	Push(&stack, 5)
	testLen(&stack, 1, t)

	Push(&stack, 6)
	testLen(&stack, 2, t)
}

func testLen[T any](stack *LinkedStack[T], expectation int64, t *testing.T) {
	if result := Len(stack); result != expectation {
		t.Fatalf(`expected length %v, got %v`, expectation, result)
	}
} 

func TestPop(t *testing.T) {
	stack := NewStack[int]()

	testPop(&stack, 0, ErrEmpty, t)

	Push(&stack, 0)
	testPop(&stack, 0, nil, t)

	Push(&stack, 1)
	testPop(&stack, 1, nil, t)
}

func testPop[T comparable](
	stack *LinkedStack[T],
	expectedValue T,
	expectedError error,
	t *testing.T,
) {
	if result, err := Pop(stack); result != expectedValue || !errors.Is(err, expectedError) {
		t.Fatalf(`expected %v, %v but got %v, %v`, expectedValue, expectedError, result, err)
	}
}
