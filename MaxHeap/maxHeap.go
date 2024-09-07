package maxHeap

import "errors"

type values interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float64 | ~float32 |
		~string
}

type MaxHeap[T values] []T

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

func (h *MaxHeap[T]) Max() (T, bool) {
	if h == nil || len(*h) == 0 {
		var defValue T
		return defValue, false
	}
	return (*h)[0], true
}

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
