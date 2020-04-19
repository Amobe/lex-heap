package heap

import "github.com/stretchr/testify/suite"

type MinHeapSuite struct {
	suite.Suite
}

func (s *MinHeapSuite) TestMinHeap() {
	h := NewBinaryMinHeap(nil)

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
	h := NewBinaryMinHeap(giving)

	s.Equal(1, h.Search(1))
	s.Equal(2, h.Search(2))
	s.Equal(3, h.Search(3))
	s.Equal(4, h.Search(5))
	s.Equal(5, h.Search(4))
	s.Equal(0, h.Search(6))
}

func (s *MinHeapSuite) TestRemove() {
	giving := []int{3, 5, 1, 4, 2}
	h := NewBinaryMinHeap(giving)

	h.Remove(3)
	h.Remove(4)
	s.Equal(1, h.Poll())
	s.Equal(2, h.Poll())
	s.Equal(5, h.Poll())
	s.Equal(0, h.Poll())
}
