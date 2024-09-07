package maxHeap

type values interface {
	~int | ~float32 |
		~string | ~byte |
		~rune
}

type MaxHeap[T values] []T

func (h MaxHeap[T]) Insert(k T) MaxHeap[T] {
	h = append(h, k)
	for i := len(h) - 1; i != 0; {
		if h[i] > h[(i-1)/2] {
			h[i], h[(i-1)/2] = h[(i-1)/2], h[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
	return h
}

func (h MaxHeap[T]) ExtractMax() (MaxHeap[T], T) {
	heapLen := len(h) - 1
	if heapLen+1 == 0 {
		var defValue T
		return h, defValue
	}
	res := h[0]
	h[0] = h[heapLen]
	h = h[:heapLen]
	var ind int
	for {
		if ind*2+1 >= heapLen {
			break
		}
		if ind*2+2 >= heapLen {
			if h[ind] < h[ind*2+1] {
				h[ind], h[ind*2+1] = h[ind*2+1], h[ind]
			}
			break
		}
		left, right := h[ind*2+1], h[ind*2+2]
		if h[ind] > left && h[ind] > right {
			break
		}
		if left > right {
			h[ind], h[ind*2+1] = h[ind*2+1], h[ind]
			ind = ind*2 + 1
		} else {
			h[ind], h[ind*2+2] = h[ind*2+2], h[ind]
			ind = ind*2 + 2
		}
	}
	return h, res
}

func (h MaxHeap[T]) Max() T {
	if len(h) == 0 {
		var defValue T
		return defValue
	}
	return h[0]
}
