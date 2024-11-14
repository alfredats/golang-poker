package classes

type Cuite int 

func (s Cuite) Valid() bool {
  return (int(s) >= 0) && (int(s) <= 3)
}

func (s Cuite) Compare(z Cuite) int {
  return int(s) - int(z)
}
