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
	tree      []int
	len       int
	isMaxHeap bool
}

func newBinaryMinHeap() *binaryHeap {
	h := &binaryHeap{
		tree:      make([]int, _HeapSize),
		isMaxHeap: false,
	}
	return h
}

func newBinaryMinHeapWithTree(tree []int) *binaryHeap {
	h := newBinaryMinHeap()
	for _, v := range tree {
		h.len++
		h.tree[h.len] = h.doNegative(v)
	}
	return h
}

func newBinaryMaxHeap() *binaryHeap {
	h := newBinaryMinHeap()
	h.isMaxHeap = true
	return h
}

func newBinaryMaxHeapWithTree(tree []int) *binaryHeap {
	h := newBinaryMaxHeap()
	for _, v := range tree {
		h.len++
		h.tree[h.len] = h.doNegative(v)
	}
	return h
}

// Insert inserts a val into the heap.
func (h *binaryHeap) Insert(val int) {
	idx := h.len + 1
	if h.idxOutOfRange(idx) { // bug
		return
	}
	h.len++
	h.tree[idx] = h.doNegative(val)
	h.bubbleUp(idx)
}

// Poll polls the smallest value from the heap.
func (h *binaryHeap) Poll() int {
	if h.len == 0 {
		return 0
	}

	val := h.tree[_RootIdx]
	h.removeIdx(_RootIdx)
	return h.doNegative(val)
}

// Remove removes the element which is same as giving val.
func (h *binaryHeap) Remove(val int) {
	idx := h.Search(h.doNegative(val))
	h.removeIdx(idx)
}

// Search searches the index of giving value, return 0 if val not found.
func (h *binaryHeap) Search(val int) int {
	for idx := _RootIdx; idx <= h.len; idx++ {
		if h.tree[idx] == h.doNegative(val) {
			return idx
		}
	}
	return 0
}

func (h *binaryHeap) IsMinHeap() bool {
	return !h.isMaxHeap
}

func (h *binaryHeap) ToMinHeap() {
	if h.isMaxHeap {
		h.switchHeap()
	}
}

func (h *binaryHeap) ToMaxHeap() {
	if !h.isMaxHeap {
		h.switchHeap()
	}
}

func (h *binaryHeap) Print() {
	fmt.Printf("%s\n", h)
}

func (h *binaryHeap) String() string {
	s := strings.Join(strings.Fields(fmt.Sprint(h.tree[_RootIdx:h.len+1])), " ")
	return s
}

func (h *binaryHeap) switchHeap() {
	for i := h.len; i > 0; i-- {
		h.tree[i] = h.tree[i] * -1
	}
	h.isMaxHeap = !h.isMaxHeap
	h.Heapify()
}

func (h *binaryHeap) doNegative(in int) int {
	if h.isMaxHeap {
		return in * -1
	}
	return in
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

func (h *binaryHeap) invalidHeap(idx int) bool {
	if h.invalidNode(idx) {
		return false
	}
	if idx != h.smallValueIdx(idx, h.leftChildIdx(idx)) ||
		idx != h.smallValueIdx(idx, h.rightChildIdx(idx)) {
		return true
	}
	return h.invalidHeap(h.leftChildIdx(idx)) || h.invalidHeap(h.rightChildIdx(idx))
}

func (h *binaryHeap) InvalidHeap() bool {
	rootIdx := 1
	return h.invalidHeap(rootIdx)
}

// Heapify makes entire tree became a valid heap.
// time complexity: O(n)
// https://stackoverflow.com/questions/9755721/how-can-building-a-heap-be-on-time-complexity
func (h *binaryHeap) Heapify() {
	nonLeafNode := h.len / 2
	for i := nonLeafNode; i > 0; i-- {
		h.bubbleDown(i)
	}
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
