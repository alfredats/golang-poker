package classes

type cuite int 

func (s cuite) valid() bool {
  return (int(s) > 0) && (int(s) < 5)
}

func (s cuite) compare(z cuite) int {
  return int(s) - int(z)
}
