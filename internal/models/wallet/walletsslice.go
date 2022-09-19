package wallet

type WalletsSlice []*Wallet

func (w WalletsSlice) Len() int           { return len(w) }
func (w WalletsSlice) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w WalletsSlice) Less(i, j int) bool { return w[i].Balance.Uint64() < w[j].Balance.Uint64() }
