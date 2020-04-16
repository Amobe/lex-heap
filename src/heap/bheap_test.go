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
	givin := []int{5, 4, 1, 7, 3}
	h := NewBinaryHeap(givin)

	for _, v := range givin {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}
