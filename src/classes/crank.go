package classes

type crank int

func (c crank) valid() bool {
  return (int(c) > 0) && (int(c) < 15)
}

func (c crank) compare(z crank) int {
  return int(c) - int(z)
}
