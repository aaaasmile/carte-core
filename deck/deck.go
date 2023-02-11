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
	items    []DeckItem
	deckType DeckType
	lblFn    GetDeckLblFunc
	valueFn  GetDeckValRankFunc
	rankFn   GetDeckValRankFunc
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

func (d *Deck) initialize(size int) {
	d.items = make([]DeckItem, 0)
	for ix := 0; ix < size; ix++ {
		item := DeckItem{
			Lbl:   d.GetDeckLbl(ix),
			Value: d.GetDeckValue(ix),
			Rank:  d.GetDeckRank(ix),
		}
		d.items = append(d.items, item)
	}
}

func (d *Deck) GetDeckLbl(ix int) string {
	return d.lblFn(ix)
}

func (d *Deck) GetDeckValue(ix int) int {
	return d.valueFn(ix)
}

func (d *Deck) GetDeckRank(ix int) int {
	return d.rankFn(ix)
}

func (d *Deck) GetItem(ix int) DeckItem {
	return d.items[ix]
}

func (d *Deck) String() string {
	res := ""
	for ix, val := range d.items {
		if ix > 0 {
			res += "," + val.Lbl
		} else {
			res = val.Lbl
		}
	}
	return res
}
