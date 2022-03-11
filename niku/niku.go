package niku

type Niku struct {
	Neme string
	Size int
}

func (a *Niku) Taberu() string {
	return a.Neme
}
