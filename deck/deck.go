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
	Lbl   string
	Rank  int
	Value int
	Suit  SuitType
	Index int
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
	d.items = make([]DeckItem, 0)
	for ix := 0; ix < size; ix++ {
		item := DeckItem{
			Lbl:   spec.GetDeckLbl(ix),
			Value: spec.GetDeckValue(ix),
			Rank:  spec.GetDeckRank(ix),
			Suit:  spec.GetDeckSuit(ix),
			Index: ix,
		}
		d.items = append(d.items, item)
	}
}

func (d *Deck) GetCard(lbl string) DeckItem {
	ix := d.find_on_lbl(lbl)
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
		if v.Lbl == lbl {
			return ix
		}
	}
	return -1
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
