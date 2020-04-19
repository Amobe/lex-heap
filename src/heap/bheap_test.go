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

func (s *BHeapSuite) TestNewBHeap() {
	giving := []int{1, 2, 3, 4, 5}
	h := NewBinaryHeap(giving)

	for _, v := range giving {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestBuildMinBHeap() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinBinaryHeap(giving)

	expect := []int{1, 2, 3, 4, 5}
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

func (s *BHeapSuite) TestBinarySearch() {
	giving := []int{3, 5, 1, 4, 2}
	h := newBinaryHeap(giving)
	rootIdx := 1

	s.Equal(1, h.bSearch(rootIdx, 3))
	s.Equal(2, h.bSearch(rootIdx, 5))
	s.Equal(3, h.bSearch(rootIdx, 1))
	s.Equal(4, h.bSearch(rootIdx, 4))
	s.Equal(5, h.bSearch(rootIdx, 2))
	s.Equal(0, h.bSearch(rootIdx, 6))

	h.MinHeapify()
	s.Equal(1, h.bSearch(rootIdx, 1))
	s.Equal(2, h.bSearch(rootIdx, 2))
	s.Equal(3, h.bSearch(rootIdx, 3))
	s.Equal(4, h.bSearch(rootIdx, 5))
	s.Equal(5, h.bSearch(rootIdx, 4))
	s.Equal(0, h.bSearch(rootIdx, 6))
}

func (s *BHeapSuite) TestMinHeapSearch() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinBinaryHeap(giving)

	s.Equal(1, h.Search(1))
	s.Equal(2, h.Search(2))
	s.Equal(3, h.Search(3))
	s.Equal(4, h.Search(5))
	s.Equal(5, h.Search(4))
	s.Equal(0, h.Search(6))
}

func (s *BHeapSuite) TestRemove() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinBinaryHeap(giving)

	h.Remove(3)
	h.Remove(4)
	s.Equal(1, h.Poll())
	s.Equal(2, h.Poll())
	s.Equal(5, h.Poll())
	s.Equal(0, h.Poll())
}
