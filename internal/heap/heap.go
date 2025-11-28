// Package heap implements heap-related data structures and algorithms.
package heap

type Heap[T any] struct {
	data  []T
	cmpFn func(a, b T) bool
}

func NewHeap[T any](cmpFn func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:  make([]T, 0),
		cmpFn: cmpFn,
	}
}

func (h *Heap[T]) Items() []T {
	return h.data
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Pop() (item T, ok bool) {
	if h.Len() == 0 {
		var nilItem T
		return nilItem, false
	}

	result := h.data[0]

	lastIdx := h.Len() - 1

	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]

	h.heapifyDown()

	return result, true
}

func (h *Heap[T]) Push(item T) {
	h.data = append(h.data, item)
	h.heapifyUp(h.Len() - 1)
}

func (h *Heap[T]) heapifyUp(idx int) {
	for {
		parentIdx := h.parentIdx(idx)

		if parentIdx == idx || !h.cmpFn(h.data[idx], h.data[parentIdx]) {
			break
		}

		h.swap(idx, parentIdx)
		idx = parentIdx
	}
}

func (h *Heap[T]) heapifyDown() {
	idx := 0
	lastIdx := h.Len() - 1

	leftIdx, rightIdx := h.leftChildIdx(idx), h.rightChildIdx(idx)
	cmpIdx := 0

	for leftIdx <= lastIdx {
		if leftIdx == lastIdx {
			cmpIdx = leftIdx
		} else if h.cmpFn(h.data[leftIdx], h.data[rightIdx]) {
			cmpIdx = leftIdx
		} else {
			cmpIdx = rightIdx
		}

		if !h.cmpFn(h.data[cmpIdx], h.data[idx]) {
			break
		}

		h.swap(idx, cmpIdx)
		idx = cmpIdx
		leftIdx, rightIdx = h.leftChildIdx(idx), h.rightChildIdx(idx)
	}
}

func (h *Heap[T]) swap(aIdx, bIdx int) {
	h.data[aIdx], h.data[bIdx] = h.data[bIdx], h.data[aIdx]
}

func (h *Heap[T]) parentIdx(idx int) int {
	return (idx - 1) / 2
}

func (h *Heap[T]) leftChildIdx(idx int) int {
	return (idx * 2) + 1
}

func (h *Heap[T]) rightChildIdx(idx int) int {
	return (idx * 2) + 2
}
