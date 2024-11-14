package unittest

import (
	"alfredats/golang-poker/src/classes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DealerTestSuite struct {
  suite.Suite
  d classes.Dealer
}

///////////////////////////////////////////////////////////////////////

func (s *DealerTestSuite) SetupSuite() { }
func (s *DealerTestSuite) TearDownSuite() {}

func (s *DealerTestSuite) SetupTest() {
  s.d = classes.Dealer{}
  s.d.Init()
}

func (s *DealerTestSuite) TearDownTest() {}

///////////////////////////////////////////////////////////////////////


func (s *DealerTestSuite) Test_DealReturnsCard() {
  n := 1
  cards, err := s.d.Deal(n)
  s.Nil(err)
  s.Len(cards, n)
}

func (s *DealerTestSuite) Test_DealReturnsFullDeck() {
  n := 52
  cards, err := s.d.Deal(n)
  s.Nil(err)
  s.Len(cards, n)

  for _, c := range cards {
    s.True(c.Valid())
  }

  cards, err = s.d.Deal(1)
  s.Error(err)
}




 
func Test_Unit_Dealer(t *testing.T) {
  suite.Run(t, new(DealerTestSuite))
}


