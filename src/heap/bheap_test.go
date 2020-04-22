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

func (s *BHeapSuite) TestNewBinaryHeap() {
	giving := []int{3, 5, 1, -1, 4, 0, 2}
	h := newBinaryHeap(giving)

	expect := []int{-1, 0, 1, 2, 3, 4, 5}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestNewBinaryMinHeap() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewBinaryMinHeap(giving)

	expect := []int{1, 2, 3, 4, 5}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestMinHeapify() {
	giving := []int{3, 5, 1, 4, 2}
	h := newBinaryHeapWithTree(giving)
	h.MinHeapify()

	expect := []int{1, 2, 3, 4, 5}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestInvalidMinHeapify() {
	giving := []int{3, 5, 1, 4, 2}
	h := newBinaryHeapWithTree(giving)

	s.True(h.InvalidMinHeap())

	h.MinHeapify()

	s.False(h.InvalidMinHeap())
}

func (s *BHeapSuite) TestMinHeapifyBig() {
	giving := []int{3, 5, 8, 4, 10, 6, 9, 2, 7, 1}
	h := newBinaryHeapWithTree(giving)
	h.MinHeapify()

	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *BHeapSuite) TestMinHeapOperation() {
	h := NewBinaryMinHeap(nil)

	h.Insert(1)
	h.Insert(3)
	h.Insert(2)
	s.Equal(1, h.Search(1))
	s.Equal(2, h.Search(3))
	s.Equal(3, h.Search(2))
	h.Remove(3)
	s.Equal(1, h.Search(1))
	s.Equal(2, h.Search(2))
	s.Equal(1, h.Poll())
	s.Equal(1, h.Search(2))
	s.Equal(2, h.Poll())
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

func (s *BHeapSuite) TestGeneralUseCase() {
	testCases := []struct {
		giving   []int
		expected []int
	}{
		{
			giving:   []int{5, 1, 4, 2, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			giving:   []int{4, 4, 3, 1, 9, 2},
			expected: []int{1, 2, 3, 4, 4, 9},
		},
		{
			giving:   []int{100},
			expected: []int{100},
		},
		{
			giving:   []int{3, 3, 3, 3, 3},
			expected: []int{3, 3, 3, 3, 3},
		},
		{
			giving:   []int{81, 87, -47, 59, -81, 18, 25, 40, 56, 0, 94, -11, 62, 89, 28, 74, 11, -45, -37, 6},
			expected: []int{-81, -47, -45, -37, -11, 0, 6, 11, 18, 25, 28, 40, 56, 59, 62, 74, 81, 87, 89, 94},
		},
	}

	for _, tc := range testCases {
		h := NewBinaryMinHeap(tc.giving)
		for _, v := range tc.expected {
			s.Equal(v, h.Poll())
		}
	}
}

func BenchmarkInvalidMinHeap(b *testing.B) {
	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	h := newBinaryHeapWithTree(s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.InvalidMinHeap()
	}
}
