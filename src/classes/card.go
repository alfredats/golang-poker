package classes

type Card struct {
  rank Crank 
  suite Cuite
}

func (c Card) Valid() bool {
  return c.rank.Valid() && c.suite.Valid()
}

func (c Card) Compare(z Card) int {
  var crank_cmp int = c.rank.Compare(z.rank)
  var cuite_cmp int = c.suite.Compare(z.suite)

  if crank_cmp != 0 {
    return crank_cmp
  } else if cuite_cmp != 0 {
    return cuite_cmp
  }
  return 0
}
