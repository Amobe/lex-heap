package heap

type Heap interface {
	Insert(int)
	Poll() int
	Remove(int)
	Search(int) int
	Print()
	IsMinHeap() bool
	ToMinHeap()
	ToMaxHeap()
}

// NewMinHeap creates a min heap with giving in values.
func NewMinHeap(in []int) Heap {
	h := newBinaryMinHeap()
	for _, v := range in {
		h.Insert(v)
	}
	h.Print()
	return h
}

// NewMaxHeap creates a max heap with giving in values.
func NewMaxHeap(in []int) Heap {
	h := newBinaryMaxHeap()
	for _, v := range in {
		h.Insert(v)
	}
	h.Print()
	return h
}
