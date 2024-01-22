package structures_test

import (
	"structures"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericStack(t *testing.T) {
	// Test the stack with integers
	intStack := structures.NewStack[int]()
	intStack.Push(42)
	intStack.Push(17)
	poppedInt := intStack.Pop()

	assert.Equal(t, poppedInt, 17, "Popped integer should be 17")

	// Test the stack with strings
	stringStack := structures.NewStack[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")
	poppedString := stringStack.Pop()

	assert.Equal(t, poppedString, "World", "Popped string should be 'World'")

	// Test the stack with booleans
	boolStack := structures.NewStack[bool]()
	boolStack.Push(true)
	boolStack.Push(false)
	poppedBool := boolStack.Pop()
	poppedBool2 := boolStack.Pop()

	assert.Equal(t, poppedBool, false, "Popped bool should be false")
	assert.Equal(t, poppedBool2, true, "Popped bool should be true")

	// Test attempt pop empty stack should not be nil
	intStack = structures.NewStack[int]()
	poppedInt = intStack.Pop()

	assert.Equal(t, poppedInt, 0, "Popped integer should be zero value")
}

func Test_GenericQueue(t *testing.T) {
	//testing queued ints
	intQueue := structures.NewQueue[int]()
	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)
	dequeued := intQueue.Dequeue()
	dequeued2 := intQueue.Dequeue()

	assert.Equal(t, 1, dequeued, "dequeued should be 1")
	assert.Equal(t, 2, dequeued2, "dequeued should be 2")

	//testing queued strings
	stringQueue := structures.NewQueue[string]()
	stringQueue.Enqueue("hello")
	stringQueue.Enqueue("world")
	stringDequeued := stringQueue.Dequeue()
	stringDequeued2 := stringQueue.Dequeue()

	assert.Equal(t, "hello", stringDequeued, "dequeued should be hello")
	assert.Equal(t, "world", stringDequeued2, "dequeued should be world")

	boolQueue := structures.NewQueue[bool]()
	boolDequeued := boolQueue.Dequeue()

	assert.Equal(t, false, boolDequeued, "dequeued should be zero value")
}
