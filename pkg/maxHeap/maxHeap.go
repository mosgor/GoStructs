// Package maxHeap contains an implementation of the Heap data structure
// which elements are sorted from largest to smallest and an implementation of methods
// that can be used to work with it.
//
// Heap is a tree-based data structure, which in fact is just a specifically ordered slice.
// If you try to iterate through it, you won't see that the elements are sorted.
// It is so because it is a tree with a specific order of elements.
//
// Also, it implements Heapify function that can convert any slice to MaxHeap.
package maxHeap

import "errors"

type values interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float64 | ~float32 |
		~string
}

// MaxHeap is a sorted in descending order implementation of the Heap data structure based on a slice.
// Because this structure is just a slice, you can use such functions as len() or cap() and even iterate through it.
// Meanwhile, it is possible to iterate through it, it is meaningless
// because the elements are ordered in a specific way.
// MaxHeap takes generic type T, that can be every int, uint, float or string type.
type MaxHeap[T values] []T

// Insert method adds a new element to the heap.
// To insert a new element, it is used so-called siftUp action.
// It returns an error if the receiver is nil.
// The Complexity of this method is O(log n).
func (h *MaxHeap[T]) Insert(k T) error {
	if h == nil {
		return errors.New("heap pointer is nil")
	}
	*h = append(*h, k)
	for i := len(*h) - 1; i != 0; {
		if (*h)[i] > (*h)[(i-1)/2] {
			(*h)[i], (*h)[(i-1)/2] = (*h)[(i-1)/2], (*h)[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
	return nil
}

// Extract method removes the root and concurrently the largest element from the heap and returns it.
// This method returns a type T default value if the heap is empty or if the receiver is nil.
// To see if the return value is an extracted value or a default value, it is used "comma ok" idiom.
// Extract method is implemented using so-called siftDown action.
// The Complexity of this method is O(log n).
func (h *MaxHeap[T]) Extract() (T, bool) {
	if h == nil || len(*h) == 0 {
		var defValue T
		return defValue, false
	}
	heapLen := len(*h) - 1
	res := (*h)[0]
	(*h)[0] = (*h)[heapLen]
	*h = (*h)[:heapLen]
	siftDown(*h, 0)
	return res, true
}

// Max method returns the root and concurrently the largest element from the heap.
// This method returns a type T default value if the heap is empty or if the receiver is nil.
// To see if the return value is really in the heap, or it is just the default value, it is used "comma ok" idiom.
// The Complexity of this method is O(1).
func (h *MaxHeap[T]) Max() (T, bool) {
	if h == nil || len(*h) == 0 {
		var defValue T
		return defValue, false
	}
	return (*h)[0], true
}

// Heapify function takes a slice of type T and returns a heap made from the slice.
// To make a heap from a slice, it is used so-called siftDown action.
// The Complexity of this method is O(n).
func Heapify[T values](arr []T) MaxHeap[T] {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		ind := i
		siftDown(arr, ind)
	}
	return arr
}

func siftDown[T values](arr []T, ind int) {
	heapLen := len(arr)
	for {
		if ind*2+1 >= heapLen {
			break
		}
		if ind*2+2 >= heapLen {
			if arr[ind] < arr[ind*2+1] {
				arr[ind], arr[ind*2+1] = arr[ind*2+1], arr[ind]
			}
			break
		}
		left, right := arr[ind*2+1], arr[ind*2+2]
		if arr[ind] > left && arr[ind] > right {
			break
		}
		if left > right {
			arr[ind], arr[ind*2+1] = arr[ind*2+1], arr[ind]
			ind = ind*2 + 1
		} else {
			arr[ind], arr[ind*2+2] = arr[ind*2+2], arr[ind]
			ind = ind*2 + 2
		}
	}
}
