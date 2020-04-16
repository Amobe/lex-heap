package heap

import (
	"fmt"
	"strings"
)

type binaryHeap struct {
	tree []int
	len  int
}

func NewBinaryHeap(in []int) Heap {
	h := &binaryHeap{
		tree: make([]int, 256),
	}
	for _, v := range in {
		h.Insert(v)
	}
	return h
}

func (h *binaryHeap) Insert(v int) {
	h.tree[h.len] = v
	h.len++
	fmt.Printf("%s\n", h)
}

func (h *binaryHeap) Poll() int {
	if h.len == 0 {
		return 0
	}

	v := h.tree[0]
	for i := 1; i < h.len; i++ {
		h.swap(&h.tree[i-1], &h.tree[i])
	}
	h.len--
	return v
}

func (h *binaryHeap) Remove(v int) {

}

func (h *binaryHeap) String() string {
	s := strings.Join(strings.Fields(fmt.Sprint(h.tree[:h.len])), " ")
	return s
}

func (h *binaryHeap) swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
