package heap

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MinHeapSuite struct {
	suite.Suite
}

func TestMinHeapSuite(t *testing.T) {
	suite.Run(t, new(MinHeapSuite))
}

func (s *MinHeapSuite) TestNewMinHeap() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinHeap(giving)

	expect := []int{1, 2, 3, 4, 5}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *MinHeapSuite) TestNewMaxHeap() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMaxHeap(giving)

	expect := []int{5, 4, 3, 2, 1}
	for _, v := range expect {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}

func (s *MinHeapSuite) TestMinHeapOperation() {
	h := NewMinHeap(nil)

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

func (s *MinHeapSuite) TestGeneralUseCase() {
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
		h := NewMinHeap(tc.giving)
		for _, v := range tc.expected {
			s.Equal(v, h.Poll())
		}
	}
}

func (s *MinHeapSuite) TestInsert() {
	h := NewMinHeap(nil)

	h.Insert(1)
	h.Insert(6)
	h.Insert(3)
	h.Insert(2)
	h.Insert(4)
	h.Insert(7)
	h.Insert(5)
	h.Insert(5)
	h.Insert(8)

	s.Equal(1, h.Poll())
	s.Equal(2, h.Poll())
	s.Equal(3, h.Poll())
	s.Equal(4, h.Poll())
	s.Equal(5, h.Poll())
	s.Equal(5, h.Poll())
	s.Equal(6, h.Poll())
	s.Equal(7, h.Poll())
	s.Equal(8, h.Poll())
}

func (s *MinHeapSuite) TestSearch() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinHeap(giving)

	s.Equal(1, h.Search(1))
	s.Equal(2, h.Search(2))
	s.Equal(3, h.Search(3))
	s.Equal(4, h.Search(5))
	s.Equal(5, h.Search(4))
	s.Equal(0, h.Search(6))
}

func (s *MinHeapSuite) TestRemove() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewMinHeap(giving)

	h.Remove(3)
	h.Remove(4)
	s.Equal(1, h.Poll())
	s.Equal(2, h.Poll())
	s.Equal(5, h.Poll())
	s.Equal(0, h.Poll())
}
