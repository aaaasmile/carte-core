package deck

type SuiteType int

type GetDeckLblFunc func(int) string
type GetDeckValRankFunc func(int) int
type GetDeckSuiteFunc func(int) SuiteType

type DeckItem struct {
	Lbl   string
	Rank  int
	Value int
	Suite SuiteType
	Index int
}

type DeckType int

type Deck struct {
	items    []DeckItem
	deckType DeckType
	lblFn    GetDeckLblFunc
	valueFn  GetDeckValRankFunc
	rankFn   GetDeckValRankFunc
	suiteFn  GetDeckSuiteFunc
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

const (
	Bastoni SuiteType = iota
	Coppe
	Denari
	Spade
)

func (d *Deck) initialize(size int) {
	d.items = make([]DeckItem, 0)
	for ix := 0; ix < size; ix++ {
		item := DeckItem{
			Lbl:   d.GetDeckLbl(ix),
			Value: d.GetDeckValue(ix),
			Rank:  d.GetDeckRank(ix),
			Suite: d.GetDeckSuite(ix),
			Index: ix,
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

func (d *Deck) GetDeckSuite(ix int) SuiteType {
	return d.suiteFn(ix)
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
