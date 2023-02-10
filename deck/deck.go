package deck

type GetDeckLblFunc func(int) string
type GetDeckValRankFunc func(int) int

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
	ValueFn  GetDeckValRankFunc
	RankFn   GetDeckValRankFunc
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

func (d *Deck) Initialize() {
	for ix, item := range d.Items {
		item.Lbl = d.GetDeckLbl(ix)
		item.Value = d.GetDeckValue(ix)
		item.Rank = d.GetDeckRank(ix)
	}
}

func (d *Deck) GetDeckLbl(ix int) string {
	return d.LblFn(ix)
}

func (d *Deck) GetDeckValue(ix int) int {
	return d.ValueFn(ix)
}

func (d *Deck) GetDeckRank(ix int) int {
	return d.RankFn(ix)
}
