package heap

import (
	"fmt"
	"strings"
)

const (
	_HeapSize = 256
	_RootIdx  = 1
)

type binaryHeap struct {
	tree []int
	len  int
}

// NewBinaryMinHeap creates a min heap with giving in values.
func NewBinaryMinHeap(in []int) Heap {
	h := newBinaryHeap(in)
	h.Print()
	return h
}

// Insert inserts a val into the heap.
func (h *binaryHeap) Insert(val int) {
	h.len++
	idx := h.len
	if h.idxOutOfRange(idx) {
		return
	}
	h.tree[idx] = val
	h.bubbleUp(idx)
}

// Poll polls the smallest value from the heap.
func (h *binaryHeap) Poll() int {
	if h.len == 0 {
		return 0
	}

	v := h.tree[_RootIdx]
	h.removeIdx(_RootIdx)
	return v
}

// Remove removes the element which is same as giving val.
func (h *binaryHeap) Remove(val int) {
	idx := h.Search(val)
	h.removeIdx(idx)
}

// Search searches the index of giving value, return 0 if val not found.
func (h *binaryHeap) Search(val int) int {
	for idx := _RootIdx; idx <= h.len; idx++ {
		if h.tree[idx] == val {
			return idx
		}
	}
	return 0
}

func (h *binaryHeap) Print() {
	fmt.Printf("%s\n", h)
}

func (h *binaryHeap) String() string {
	s := strings.Join(strings.Fields(fmt.Sprint(h.tree[_RootIdx:h.len+1])), " ")
	return s
}

func newBinaryHeap(in []int) *binaryHeap {
	h := &binaryHeap{
		tree: make([]int, _HeapSize),
	}
	for _, v := range in {
		h.Insert(v)
	}
	return h
}

func newBinaryHeapWithTree(tree []int) *binaryHeap {
	h := &binaryHeap{
		tree: make([]int, _HeapSize),
	}
	for _, v := range tree {
		h.len++
		h.tree[h.len] = v
	}
	h.MinHeapify()
	return h
}

// RemoveIdx removes the element at the place of giving idx.
func (h *binaryHeap) removeIdx(idx int) {
	if h.invalidNode(idx) {
		return
	}
	h.swapIdx(idx, h.len)
	h.tree[h.len] = 0
	h.len--
	h.bubbleDown(idx)
}

func (h *binaryHeap) bubbleUp(idx int) {
	if h.invalidNode(idx) {
		return
	}
	largest := h.largeValueIdx(idx, h.parentIdx(idx))
	if idx != largest {
		h.swapIdx(largest, idx)
		h.bubbleUp(largest)
	}
}

func (h *binaryHeap) bubbleDown(idx int) {
	if h.invalidNode(idx) {
		return
	}
	smallest := h.smallValueIdx(h.leftChildIdx(idx), h.rightChildIdx(idx))
	if idx != h.smallValueIdx(idx, smallest) {
		h.swapIdx(smallest, idx)
		h.bubbleDown(smallest)
	}
}

func (h *binaryHeap) swapIdx(idxA, idxB int) {
	if h.idxOutOfRange(idxA) || h.idxOutOfRange(idxB) {
		return
	}
	tmp := h.tree[idxA]
	h.tree[idxA] = h.tree[idxB]
	h.tree[idxB] = tmp
}

func (h *binaryHeap) compareValueIdex(idxA, idxB int, cmpFunc func(idxA, idxB int) bool) int {
	if cmpFunc == nil {
		return 0
	}
	if h.invalidNode(idxB) {
		return idxA
	}
	if h.invalidNode(idxA) {
		return idxB
	}
	if cmpFunc(idxA, idxB) {
		return idxA
	}
	return idxB
}

func (h *binaryHeap) smallValueIdx(idxA, idxB int) int {
	cmp := func(idxA, idxB int) bool {
		return h.tree[idxA] < h.tree[idxB]
	}
	return h.compareValueIdex(idxA, idxB, cmp)
}

func (h *binaryHeap) largeValueIdx(idxA, idxB int) int {
	cmp := func(idxA, idxB int) bool {
		return h.tree[idxA] > h.tree[idxB]
	}
	return h.compareValueIdex(idxA, idxB, cmp)
}

func (h *binaryHeap) minHeapify(root int) {
	if h.invalidNode(root) {
		return
	}
	left := h.leftChildIdx(root)
	right := h.rightChildIdx(root)
	if root != h.smallValueIdx(root, left) {
		h.swapIdx(root, left)
	}
	if root != h.smallValueIdx(root, right) {
		h.swapIdx(root, right)
	}
	h.minHeapify(left)
	h.minHeapify(right)
}

// MinHeapify makes entire tree became a valid min heap.
func (h *binaryHeap) MinHeapify() {
	rootIdx := 1
	h.minHeapify(rootIdx)
}

func (h *binaryHeap) parentIdx(i int) int {
	idx := i / 2
	if h.idxOutOfRange(idx) || h.idxOutOfRange(i) {
		return 0
	}
	return idx
}

func (h *binaryHeap) leftChildIdx(i int) int {
	idx := i * 2
	if h.idxOutOfRange(idx) || h.idxOutOfRange(i) {
		return 0
	}
	return idx
}

func (h *binaryHeap) rightChildIdx(i int) int {
	idx := i*2 + 1
	if h.idxOutOfRange(idx) || h.idxOutOfRange(i) {
		return 0
	}
	return idx
}

func (h *binaryHeap) emptyNode(i int) bool {
	return i > h.len
}

func (h *binaryHeap) invalidNode(i int) bool {
	return h.idxOutOfRange(i) || h.emptyNode(i)
}

func (h *binaryHeap) idxOutOfRange(i int) bool {
	return i >= _HeapSize || i <= 0
}
