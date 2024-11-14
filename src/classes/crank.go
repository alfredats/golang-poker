package classes

type Crank int

func (c Crank) Valid() bool {
  return (int(c) >= 0) && (int(c) <= 12)
}

func (c Crank) Compare(z Crank) int {
  return int(c) - int(z)
}
