package deck

type GetDeckLblFunc func(int) string

type DeckItem struct {
	Lbl   string
	Rank  int
	Value int
}

type DeckType int

type Deck struct {
	Items    []DeckItem
	deckType DeckType
	LblFn    GetDeckLblFunc
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

func (d *Deck) GetDeckLbl(ix int) string {
	return d.LblFn(ix)
}
