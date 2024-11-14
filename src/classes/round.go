package classes


type Round struct {
  acts []PlayAct
}

func (r *Round) valid(p PlayAct) bool {
  if p == nil { 
    return false
  }
  return false
}
