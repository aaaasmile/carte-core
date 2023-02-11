package deck

type SuitType int

type DeckSpecializer interface {
	GetDeckLbl(int) string
	GetDeckValue(int) int
	GetDeckRank(int) int
	GetDeckSuit(int) SuitType
	GetDeckType() DeckType
}

type DeckItem struct {
	lbl   string
	rank  int
	value int
	suit  SuitType
	index int
}

// DeckItem is readonly
func (di *DeckItem) Value() int {
	return di.value
}

func (di *DeckItem) Rank() int {
	return di.rank
}

func (di *DeckItem) Suit() SuitType {
	return di.suit
}

func (di *DeckItem) Lbl() string {
	return di.lbl
}

type DeckType int

type Deck struct {
	items []DeckItem
	spec  DeckSpecializer
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

func (d *Deck) Initialize(size int, spec DeckSpecializer) {
	d.items = make([]DeckItem, 0, size)
	for ix := 0; ix < size; ix++ {
		item := DeckItem{
			lbl:   spec.GetDeckLbl(ix),
			value: spec.GetDeckValue(ix),
			rank:  spec.GetDeckRank(ix),
			suit:  spec.GetDeckSuit(ix),
			index: ix,
		}
		d.items = append(d.items, item)
	}
}

func (d *Deck) GetCard(lbl string) DeckItem {
	ix := d.find_on_lbl(lbl)
	return d.items[ix]
}

func (d *Deck) GetItem(ix int) DeckItem {
	return d.items[ix]
}

func (d *Deck) GetRank(lbl string) int {
	ix := d.find_on_lbl(lbl)
	return d.spec.GetDeckRank(ix)
}

func (d *Deck) GetValue(lbl string) int {
	ix := d.find_on_lbl(lbl)
	return d.spec.GetDeckValue(ix)
}

func (d *Deck) GetSuit(lbl string) SuitType {
	ix := d.find_on_lbl(lbl)
	return d.spec.GetDeckSuit(ix)
}

func (d *Deck) find_on_lbl(lbl string) int {
	for ix, v := range d.items {
		if v.lbl == lbl {
			return ix
		}
	}
	return -1
}

func (d *Deck) String() string {
	res := ""
	for ix, val := range d.items {
		if ix > 0 {
			res += "," + val.lbl
		} else {
			res = val.lbl
		}
	}
	return res
}
