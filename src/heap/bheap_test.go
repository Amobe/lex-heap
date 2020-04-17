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
	givin := []int{1, 2, 3, 4, 5}
	h := NewBinaryHeap(givin)

	for _, v := range givin {
		s.Equal(v, h.Poll())
	}

	s.Equal(0, h.Poll())
}
