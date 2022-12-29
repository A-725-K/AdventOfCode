package types

type Pair[TF any, TS any] struct {
  first TF
  second TS
}

func NewPair[TF any, TS any](fst TF, snd TS) Pair[TF, TS] {
  return Pair[TF, TS]{first: fst, second: snd}
} 

func (p Pair[TF, TS]) Fst() TF {
  return p.first
}

func (p Pair[TF, TS]) Snd() TS {
  return p.second
}

