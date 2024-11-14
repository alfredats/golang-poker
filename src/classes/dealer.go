package classes

import (
	crand "crypto/rand"
	rand "math/rand"

	"encoding/binary"
  "errors"
)

const (
  crank_cnt = 13
  cuite_cnt = 4
  Deck_total = 52
)

type Dealer struct {
  deck []int
  dealt int
  rng *rand.Rand
  rng_src rand.Source

}

func (d *Dealer) Init() {
  d.deck = make([]int, Deck_total)
  for i := 0; i < Deck_total; i++ {
    d.deck[i] = i
  }

  var b [8]byte
  _, err := crand.Read(b[:])
  if err != nil {
    panic("cannot seed rng with secure rng")
  }

  d.rng_src= rand.NewSource(int64(binary.LittleEndian.Uint64(b[:])))
  d.rng = rand.New(d.rng_src)

  d.rng.Shuffle(len(d.deck), func(i, j int) {
    d.deck[i], d.deck[j] = d.deck[j], d.deck[i]
  })
}


func (d *Dealer) Deal(n int) ([]Card, error) {
  if d.dealt >=  Deck_total {
    return nil, errors.New("No more cards available")
  }
  end := d.dealt + n
  var err error = nil
  if end > Deck_total { 
    n = Deck_total - d.dealt
    end = Deck_total 
    err = errors.New("Deck is fully dealt")
  }

  cards := make([]Card, n)
  for i, c := range d.deck[d.dealt : end] {
    cards[i] = Card{Crank(c%crank_cnt), Cuite(c/crank_cnt)}
  }
  d.dealt = end
  return cards, err 
}

