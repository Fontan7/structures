package structures

//"slices"

type LinkedList[T any] struct {
	Head   *Node[T]
	Length int
}

type Node[T any] struct {
	Items    []T
	NextNode *Node[T]
}

func NewList[T any]() *LinkedList[T] {
	return &LinkedList[T]{Head: nil, Length: 0}
}

func (l *LinkedList[T]) Prepend(newHead *Node[T]) {
	newHead.NextNode = l.Head
	l.Head = newHead
	l.Length++
}

func (l *LinkedList[T]) Append(newNode *Node[T]) {
	if l.Head == nil { //If list empty, add value at head
		l.Prepend(newNode)
		return
	}

	current := l.Head
	for current.NextNode != nil { //While still have nodes keep cycling
		current = current.NextNode
	}
	current.NextNode = newNode //When end reached add new node
	l.Length++
}

//This function will delete the whole node when it finds a match, not the value found, will delete all
//subsecuent matches too.
//It needs to receive as parameter the type of equality you want to find
//You can do so by creating a var and passing it such as: var equals = func(a, b int) bool { return a == b }
//A more complex type of equality (assuming you have a struct with these fields) can be:
//func equalsProduct(a, b Product) bool {return a.ID == b.ID && a.Price == b.Price}
func (l *LinkedList[T]) Delete(toDelete T, equals func(T, T) bool) {
	if l.Head == nil { // Check if the list is empty
		return
	}

	// Handling the case where the head needs to be deleted
	for l.Head != nil && contains(l.Head.Items, toDelete, equals) {
		l.Head = l.Head.NextNode
		l.Length--
	}

	current := l.Head
	var previous *Node[T]

	for current != nil {
		if contains(current.Items, toDelete, equals) {

			previous.NextNode = current.NextNode // Bypass the current node

			l.Length--
		} else {
			previous = current
		}
		current = current.NextNode
	}
}

//It needs to receive as parameter the type of equality you want to find
//You can do so by creating a var and passing it such as: var equals = func(a, b int) bool { return a == b }
//A more complex type of equality (assuming you have a struct with these fields) can be:
//func equalsProduct(a, b Product) bool {return a.ID == b.ID && a.Price == b.Price}
func (l *LinkedList[T]) Search(toSearch T, equals func(T, T) bool) *Node[T] {
	current := l.Head

	for current != nil {
		if contains(current.Items, toSearch, equals) {
			return current
		}
		current = current.NextNode
	}

	return nil // Return nil if the item is not found
}

// contains checks if a slice of T contains a value of T
func contains[T any](slice []T, value T, equals func(T, T) bool) bool {
	for _, item := range slice {
		if equals(item, value) {
			return true
		}
	}
	return false
}
