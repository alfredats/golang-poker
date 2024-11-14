package classes

type card struct {
  rank crank 
  suite cuite
}


func (c card) valid() bool {
  return c.rank.valid() && c.suite.valid()
}

func (c card) compare(z card) int {
  var crank_cmp int = c.rank.compare(z.rank)
  var cuite_cmp int = c.suite.compare(z.suite)

  if crank_cmp != 0 {
    return crank_cmp
  } else if cuite_cmp != 0 {
    return cuite_cmp
  }
  return 0
}
