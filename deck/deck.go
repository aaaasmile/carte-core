package deck

var _deck40Lbs = []string{"_Ab", "_2b", "_3b", "_4b", "_5b", "_6b", "_7b", "_Fb", "_Cb", "_Rb",
	"_Ac", "_2c", "_3c", "_4c", "_5c", "_6c", "_7c", "_Fc", "_Cc", "_Rc",
	"_Ad", "_2d", "_3d", "_4d", "_5d", "_6d", "_7d", "_Fd", "_Cd", "_Rd",
	"_As", "_2s", "_3s", "_4s", "_5s", "_6s", "_7s", "_Fs", "_Cs", "_Rs"}

var _deck52Lbs = []string{"_Ab", "_2b", "_3b", "_4b", "_5b", "_6b", "_7b", "_8b", "_9b", "_10b", "_Fb", "_Cb", "_Rb",
	"_Ac", "_2c", "_3c", "_4c", "_5c", "_6c", "_7c", "_8c", "_9c", "_10c", "_Fc", "_Cc", "_Rc",
	"_Ad", "_2d", "_3d", "_4d", "_5d", "_6d", "_7d", "_8d", "_9d", "_10d", "_Fd", "_Cd", "_Rd",
	"_As", "_2s", "_3s", "_4s", "_5s", "_6s", "_7s", "_8s", "_9s", "_10s", "_Fs", "_Cs", "_Rs"}

type DeckItem struct {
	Lbl   string
	Rank  int
	Value int
}

type DeckType int

type Deck struct {
	Items    []DeckItem
	deckType DeckType
}

const (
	DTBriscola DeckType = iota
	DTRamino
)

func (d *Deck) InitToForty() {
	d.deckType = DTBriscola
	d.Items = make([]DeckItem, 40)
	for ix, item := range d.Items {
		item.Lbl = d.GetDeckLbl(ix)
		item.Value = ix
		item.Rank = 0
	}
}

func (d *Deck) GetDeckLbl(ix int) string {
	switch d.deckType {
	case DTBriscola:
		return _deck40Lbs[ix]
	case DTRamino:
		return _deck52Lbs[ix]
	}
	return ""
}
