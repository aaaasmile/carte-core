package deck

type SuitType int

type GetDeckLblFunc func(int) string
type GetDeckValRankFunc func(int) int
type GetDeckSuitFunc func(int) SuitType

type DeckItem struct {
	Lbl   string
	Rank  int
	Value int
	Suit  SuitType
	Index int
}

type DeckType int

type Deck struct {
	items    []DeckItem
	DeckType DeckType
	LblFn    GetDeckLblFunc
	ValueFn  GetDeckValRankFunc
	RankFn   GetDeckValRankFunc
	SuitFn   GetDeckSuitFunc
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

const (
	Bastoni SuitType = iota
	Coppe
	Denari
	Spade
	Quadri
	Fiori
	Cuori
	Picche
	Jolly
)

func (d *Deck) Initialize(size int) {
	d.items = make([]DeckItem, 0)
	for ix := 0; ix < size; ix++ {
		item := DeckItem{
			Lbl:   d.GetDeckLbl(ix),
			Value: d.GetDeckValue(ix),
			Rank:  d.GetDeckRank(ix),
			Suit:  d.GetDeckSuit(ix),
			Index: ix,
		}
		d.items = append(d.items, item)
	}
}

func (d *Deck) GetRank(lbl string) int {
	ix := d.find_on_lbl(lbl)
	return d.RankFn(ix)
}

func (d *Deck) GetValue(lbl string) int {
	ix := d.find_on_lbl(lbl)
	return d.ValueFn(ix)
}

func (d *Deck) GetSuit(lbl string) SuitType {
	ix := d.find_on_lbl(lbl)
	return d.SuitFn(ix)
}

func (d *Deck) find_on_lbl(lbl string) int {
	for ix, v := range d.items {
		if v.Lbl == lbl {
			return ix
		}
	}
	return -1
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

func (d *Deck) GetDeckSuit(ix int) SuitType {
	return d.SuitFn(ix)
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
