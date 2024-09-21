package sumSegmentTree

import "errors"

type values interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float64 | ~float32
}

type SumSegmentTree[T values] struct {
	elements []T
	arrLen   int
}

func MakeSegmentTree[T values](arr *[]T) SumSegmentTree[T] {
	var st SumSegmentTree[T]
	st.arrLen = len(*arr)
	st.elements = make([]T, len(*arr)*4)
	st.build(0, 0, st.arrLen, arr) // O(n)
	return st
}

func (st *SumSegmentTree[T]) build(index, lBorder, rBorder int, arr *[]T) {
	if rBorder-lBorder == 1 {
		st.elements[index] = (*arr)[lBorder]
	} else {
		mid := (lBorder + rBorder) / 2
		st.build(index*2+1, lBorder, mid, arr)
		st.build(index*2+2, mid, rBorder, arr)
		st.elements[index] = st.elements[index*2+1] + st.elements[index*2+2]
	}
}

func (st *SumSegmentTree[T]) update(pos, l, r, index int, value T) {
	if l > index || r <= index {
		return
	}
	if r-l == 1 {
		st.elements[pos] = value
		return
	}
	if mid := (l + r) / 2; index < mid {
		st.update(pos*2+1, l, mid, index, value)
	} else {
		st.update(pos*2+2, mid, r, index, value)
	}
	st.elements[pos] = st.elements[pos*2+1] + st.elements[pos*2+2]
}

func (st *SumSegmentTree[T]) Update(index int, value T) error {
	if index < 0 || index > st.arrLen {
		return errors.New("index out of range")
	}
	st.update(0, 0, st.arrLen, index, value) // O(log n)
	return nil
}

func (st *SumSegmentTree[T]) sumRange(pos, lBorder, rBorder, left, right int) T {
	if lBorder >= left && rBorder <= right {
		return st.elements[pos]
	}
	if rBorder <= left || lBorder >= right {
		return 0
	}
	mid := (lBorder + rBorder) / 2
	return st.sumRange(pos*2+1, lBorder, mid, left, right) + st.sumRange(pos*2+2, mid, rBorder, left, right)
}

func (st *SumSegmentTree[T]) SumRange(left, right int) (T, bool) {
	if left > right || left < 0 || right > st.arrLen {
		var def T
		return def, false
	}
	return st.sumRange(0, 0, st.arrLen, left, right), true // O(log n)
}
