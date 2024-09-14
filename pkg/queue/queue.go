// Package queue provides a simple queue implementation which can be used with any type.
//
// Queue is a linked list type data structure, which means that its elements are placed in the memory in an unrelated order.
// Every element contains a pointer to the next element. It makes iteration over the elements possible with cycles.
// But iteration in the middle of the queue is slow, and maybe you should search for a better structure (non-list) if you need
// to iterate over the elements often.
//
// Queue follows the FIFO principle. It means that the first element added to the queue will be the first one to be removed.
package queue

// node represents a single element of the queue. It contains only content variable and a pointer to the next element.
type node[T any] struct {
	content T
	next    *node[T]
}

// Queue is a structure containing a pointer to the first element of the queue and a pointer to the last element of the queue.
// It also contains the size of the queue.
type Queue[T any] struct {
	start *node[T]
	end   *node[T]
	size  int
}

// Push method adds a new element to the end of the queue.
// It takes an element as a parameter and adds it to the end of the queue.
// Push method returns nothing.
// The Complexity of this method is O(1).
func (q *Queue[T]) Push(n T) {
	var newNode node[T]
	newNode.content = n
	if q.start == nil {
		q.start = &newNode
		q.end = &newNode
	} else {
		q.end.next = &newNode
		q.end = &newNode
	}
	q.size++
}

// Pop method removes the first element of the queue and returns it.
// It takes nothing and returns the removed element and a boolean value indicating whether the operation was successful.
// If the queue is empty, it returns the default value of the element type and false ("Comma, ok" idiom).
// The Complexity of this method is O(1).
func (q *Queue[T]) Pop() (T, bool) {
	if q.start != nil {
		val := q.start.content
		q.start = q.start.next
		q.size--
		return val, true
	} else {
		var defValue T
		return defValue, false
	}
}

// Front method takes nothing and returns the first element of the queue and a
// boolean value indicating whether the operation was successful. If the queue is
// empty, it returns the default value of the element type and false ("Comma, ok"
// idiom).
// The Complexity of this method is O(1).
func (q *Queue[T]) Front() (T, bool) {
	if q.start != nil {
		return q.start.content, true
	} else {
		var defValue T
		return defValue, false
	}
}

// Clear method removes all elements from the queue.
// It takes nothing and returns nothing.
// The Complexity of this method is O(1).
func (q *Queue[T]) Clear() {
	q.start = nil
	q.end = nil
	q.size = 0
}

// Size method returns the size of the queue.
// It takes nothing and returns the size of the queue.
// The Complexity of this method is O(1).
func (q *Queue[T]) Size() int {
	return q.size
}
