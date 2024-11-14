package classes

import (
	crand "crypto/rand"
	"encoding/binary"
	rand "math/rand"
)

const (
  crank_cnt = 13
  cuite_cnt = 4
)

type Dealer struct {
  dealt [][]int
  rng *rand.Rand
  rng_src rand.Source

}

func (d *Dealer) Init() {
  d.dealt = make([][]int, crank_cnt)
  for i := range d.dealt {
    d.dealt[i] = make([]int, cuite_cnt)
  }

  var b [8]byte
  _, err := crand.Read(b[:])
  if err != nil {
    panic("cannot seed rng with secure rng")
  }

  d.rng_src= rand.NewSource(int64(binary.LittleEndian.Uint64(b[:])))
  d.rng = rand.New(d.rng_src)
}

// internal method for dealing a single card
func (d *Dealer) deal() card {
  sut := d.rng.Intn(cuite_cnt)
  rnk := d.rng.Intn(crank_cnt)

  if (d.dealt[rnk][sut] != 0) {
    return d.deal()
  }

  d.dealt[rnk][sut] = 1
  return card{crank(rnk), cuite(sut)}
}

func (d *Dealer) Deal(n uint) []card {
  cards := make([]card, n)
  for i := range cards {
    cards[i] = d.deal()
  }
  return cards
}

