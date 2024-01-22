package structures_test

import (
	"structures" // import your structures package
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	list := structures.NewList[int]() // create a new list of int
	node1 := &structures.Node[int]{Items: []int{1}}
	node2 := &structures.Node[int]{Items: []int{2}}
	node3 := &structures.Node[int]{Items: []int{3}}

	list.Prepend(node3)
	list.Prepend(node2)
	list.Prepend(node1)

	// Verify Length and Head
	assert.Equal(t, 3, list.Length)
	assert.Equal(t, node1, list.Head)

	// Verify Order
	assert.Equal(t, []int{1}, list.Head.Items)
	assert.Equal(t, []int{2}, list.Head.NextNode.Items)
	assert.Equal(t, []int{3}, list.Head.NextNode.NextNode.Items)
}

func TestAppend(t *testing.T) {
	list := structures.NewList[int]() // create a new list of int
	node1 := &structures.Node[int]{Items: []int{1}}
	node2 := &structures.Node[int]{Items: []int{2}}
	node3 := &structures.Node[int]{Items: []int{3}}

	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	// Verify Length
	assert.Equal(t, 3, list.Length)

	// Verify Order
	assert.Equal(t, []int{1}, list.Head.Items)
	assert.Equal(t, []int{2}, list.Head.NextNode.Items)
	assert.Equal(t, []int{3}, list.Head.NextNode.NextNode.Items)

	// Verify Tail
	tail := list.Head.NextNode.NextNode
	assert.Equal(t, node3, tail)
	assert.Nil(t, tail.NextNode) // The last node should point to nil
}

func TestDelete(t *testing.T) {
	equals := func(a, b int) bool { return a == b }

	t.Run("Delete from an empty list", func(t *testing.T) {
		list := structures.NewList[int]()
		list.Delete(1, equals)
		assert.Nil(t, list.Head)
		assert.Equal(t, 0, list.Length)
	})

	t.Run("Delete a non-existent node", func(t *testing.T) {
		list := structures.NewList[int]()
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Delete(2, equals)
		assert.Equal(t, 1, list.Length)
	})

	t.Run("Delete the head node", func(t *testing.T) {
		list := structures.NewList[int]()
		node := &structures.Node[int]{Items: []int{1}}
		list.Append(node)
		list.Delete(1, equals)
		assert.Nil(t, list.Head)
		assert.Equal(t, 0, list.Length)
	})

	t.Run("Delete a middle node", func(t *testing.T) {
		list := structures.NewList[int]()
		node1 := &structures.Node[int]{Items: []int{1}}
		node2 := &structures.Node[int]{Items: []int{2}}
		list.Append(node1)
		list.Append(node2)
		list.Append(&structures.Node[int]{Items: []int{3}})
		list.Delete(2, equals)
		assert.Equal(t, 2, list.Length)
		assert.Equal(t, []int{1}, list.Head.Items)
		assert.Equal(t, []int{3}, list.Head.NextNode.Items)
	})

	t.Run("Delete the last node", func(t *testing.T) {
		list := structures.NewList[int]()
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Append(&structures.Node[int]{Items: []int{2}})
		list.Delete(2, equals)
		assert.Equal(t, 1, list.Length)
		assert.Nil(t, list.Head.NextNode)
	})

	t.Run("Delete multiple nodes", func(t *testing.T) {
		list := structures.NewList[int]()
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Append(&structures.Node[int]{Items: []int{2}})
		list.Delete(1, equals)
		assert.Equal(t, 1, list.Length)
		assert.Equal(t, []int{2}, list.Head.Items)
	})

	t.Run("Delete subsecuent nodes", func(t *testing.T) {
		list := structures.NewList[int]()
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Append(&structures.Node[int]{Items: []int{1}})
		list.Append(&structures.Node[int]{Items: []int{2}})
		list.Append(&structures.Node[int]{Items: []int{3}})

		// Delete the nodes with value 1, which are the first two nodes (consecutive head nodes)
		list.Delete(1, equals)

		assert.Equal(t, 2, list.Length)
		assert.NotNil(t, list.Head)
		assert.Equal(t, []int{2}, list.Head.Items) // The new head should now have the value 2
		assert.Equal(t, []int{3}, list.Head.NextNode.Items)
	})
}

func TestSearch(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	list := structures.NewList[int]()
	node1 := &structures.Node[int]{Items: []int{1, 2}}
	node2 := &structures.Node[int]{Items: []int{3, 4}}
	node3 := &structures.Node[int]{Items: []int{5, 6}}

	list.Append(node1)
	list.Append(node2)
	list.Append(node3)

	// Test for finding an existing item
	result := list.Search(3, equals)
	assert.Equal(t, node2, result)

	// Test for an item that doesn't exist
	result = list.Search(7, equals)
	assert.Nil(t, result)

	// Test for an item in the first node
	result = list.Search(1, equals)
	assert.Equal(t, node1, result)

	// Test for an item in the last node
	result = list.Search(6, equals)
	assert.Equal(t, node3, result)
}
