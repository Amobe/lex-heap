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

func newBinaryHeap(in []int) *binaryHeap {
	h := &binaryHeap{
		tree: make([]int, _HeapSize),
	}
	for _, v := range in {
		h.Insert(v)
	}
	return h
}

func NewBinaryHeap(in []int) Heap {
	h := newBinaryHeap(in)
	h.Print()
	return h
}

func NewMinBinaryHeap(in []int) Heap {
	h := newBinaryHeap(in)
	h.MinHeapify()
	h.Print()
	return h
}

func (h *binaryHeap) Insert(v int) {
	h.tree[h.len+1] = v
	h.len++
}

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

// RemoveIdx removes the element at the place of giving idx.
func (h *binaryHeap) removeIdx(idx int) {
	if h.idxOutOfRange(idx) || h.emptyNode(idx) {
		return
	}
	h.swapIdx(idx, h.len)
	h.len--
	h.bubbleDown(idx)
}

func (h *binaryHeap) bubbleDown(idx int) {
	if h.idxOutOfRange(idx) || h.emptyNode(idx) {
		return
	}
	smallest := h.smallValueIdx(h.leftChildIdx(idx), h.rightChildIdx(idx))
	if idx != h.smallValueIdx(idx, smallest) {
		h.swapIdx(smallest, idx)
		h.bubbleDown(smallest)
	}
}

func (h *binaryHeap) Print() {
	fmt.Printf("%s\n", h)
}

func (h *binaryHeap) String() string {
	s := strings.Join(strings.Fields(fmt.Sprint(h.tree[_RootIdx:h.len+1])), " ")
	return s
}

func (h *binaryHeap) swapIdx(idxA, idxB int) {
	if h.idxOutOfRange(idxA) || h.idxOutOfRange(idxB) {
		return
	}
	tmp := h.tree[idxA]
	h.tree[idxA] = h.tree[idxB]
	h.tree[idxB] = tmp
}

func (h *binaryHeap) smallValueIdx(idxA, idxB int) int {
	if h.idxOutOfRange(idxA) || h.emptyNode(idxA) {
		return idxB
	}
	if h.idxOutOfRange(idxB) || h.emptyNode(idxB) {
		return idxA
	}
	if h.tree[idxA] < h.tree[idxB] {
		return idxA
	}
	return idxB
}

func (h *binaryHeap) minHeapify(root int) {
	if h.idxOutOfRange(root) || h.emptyNode(root) {
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

func (h *binaryHeap) idxOutOfRange(i int) bool {
	return i >= _HeapSize || i <= 0
}

// bSearch searches the index of giving val in a binary tree start at giving idx.
func (h *binaryHeap) bSearch(idx, val int) int {
	if h.idxOutOfRange(idx) || h.emptyNode(idx) {
		return 0
	}
	if h.tree[idx] == val {
		return idx
	}
	if find := h.bSearch(h.leftChildIdx(idx), val); !h.idxOutOfRange(find) {
		return find
	}
	if find := h.bSearch(h.rightChildIdx(idx), val); !h.idxOutOfRange(find) {
		return find
	}
	return 0
}

// Search searches the index of giving value, return 0 if not found.
func (h *binaryHeap) Search(val int) int {
	return h.bSearch(_RootIdx, val)
}
