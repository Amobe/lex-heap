package heap

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BHeapSuite struct {
	suite.Suite
}

func TestBHeapSuite(t *testing.T) {
	suite.Run(t, new(BHeapSuite))
}

func (s *BHeapSuite) TestNewBinaryMinHeap() {
	h := newBinaryMinHeap()
	s.NotNil(h)
	s.NotNil(h.tree)
	s.False(h.isMaxHeap)
}

func (s *BHeapSuite) TestNewBinaryMaxHeap() {
	h := newBinaryMaxHeap()
	s.NotNil(h)
	s.NotNil(h.tree)
	s.True(h.isMaxHeap)
}

func (s *BHeapSuite) TestMinHeapify() {
	giving := []int{3, 5, 1, 4, 2}
	h := newBinaryMinHeapWithTree(giving)
	h.Heapify()

	expect := []int{1, 2, 3, 4, 5}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestMaxHeapify() {
	giving := []int{3, 5, 1, 4, 2}
	h := newBinaryMaxHeapWithTree(giving)
	h.Heapify()

	expect := []int{5, 4, 3, 2, 1}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestInvalidMinHeapify() {
	giving := []int{3, 5, 8, 4, 10, 6, 9, 2, 7, 1}
	h := newBinaryMinHeapWithTree(giving)

	s.True(h.InvalidHeap())
	h.Heapify()
	s.False(h.InvalidHeap())
}

func (s *BHeapSuite) TestInvalidMaxHeapify() {
	giving := []int{3, 5, 8, 4, 10, 6, 9, 2, 7, 1}
	h := newBinaryMaxHeapWithTree(giving)

	s.True(h.InvalidHeap())
	h.Heapify()
	s.False(h.InvalidHeap())
}

func (s *BHeapSuite) TestToMinHeap() {
	giving := []int{3, 5, 8, 4, 10, 6, 9, 2, 7, 1}
	h := newBinaryMaxHeapWithTree(giving)
	h.Heapify()
	h.ToMinHeap()
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}
	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestToMaxHeap() {
	giving := []int{3, 5, 8, 4, 10, 6, 9, 2, 7, 1}
	h := newBinaryMinHeapWithTree(giving)
	h.Heapify()
	h.ToMaxHeap()
	expect := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}
	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestParentIdx() {
	h := binaryHeap{}
	s.Equal(0, h.parentIdx(0)) // out of range
	s.Equal(0, h.parentIdx(1)) // out of range
	s.Equal(1, h.parentIdx(3))
	s.Equal(62, h.parentIdx(125))
	s.Equal(63, h.parentIdx(126))
	s.Equal(63, h.parentIdx(127))
	s.Equal(64, h.parentIdx(128))
}

func (s *BHeapSuite) TestLeftChildIdx() {
	h := binaryHeap{}
	s.Equal(0, h.leftChildIdx(0)) // out of range
	s.Equal(2, h.leftChildIdx(1))
	s.Equal(6, h.leftChildIdx(3))
	s.Equal(250, h.leftChildIdx(125))
	s.Equal(252, h.leftChildIdx(126))
	s.Equal(254, h.leftChildIdx(127))
	s.Equal(0, h.leftChildIdx(128)) // out of range
}

func (s *BHeapSuite) TestRightChildIdx() {
	h := binaryHeap{}
	s.Equal(0, h.rightChildIdx(0)) // out of range
	s.Equal(3, h.rightChildIdx(1))
	s.Equal(7, h.rightChildIdx(3))
	s.Equal(251, h.rightChildIdx(125))
	s.Equal(253, h.rightChildIdx(126))
	s.Equal(255, h.rightChildIdx(127))
	s.Equal(0, h.rightChildIdx(128)) // out of range
}

func BenchmarkInvalidMinHeap(b *testing.B) {
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	h := newBinaryMinHeapWithTree(s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.InvalidHeap()
	}
}
