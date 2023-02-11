package ramino

import "github.com/aaaasmile/carte-core/deck"

// d = diamonds (quadri), c = clubs (fiori), h = hearts (cuori), s = spade (picche), Jo = Jolly
var _deckLbs = []string{"_Ad", "_2d", "_3d", "_4d", "_5d", "_6d", "_7d", "_8d", "_9d", "_10d", "_Jd", "_Qd", "_Kd",
	"_Ac", "_2c", "_3c", "_4c", "_5c", "_6c", "_7c", "_8c", "_9c", "_10c", "_Jc", "_Qc", "_Kc",
	"_Ah", "_2h", "_3h", "_4h", "_5h", "_6h", "_7h", "_8h", "_9h", "_10h", "_Jh", "_Qh", "_Kh",
	"_As", "_2s", "_3s", "_4s", "_5s", "_6s", "_7s", "_8s", "_9s", "_10s", "_Js", "_Qs", "_Ks",
	"_Jo", "_Jo"}

var _points = []int{11, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10}

var _rank = []int{14, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type DeckRamino struct {
	deck.Deck
}

func (d *DeckRamino) Init() {

	d.Initialize(108, d)
}

func (d *DeckRamino) GetDeckType() deck.DeckType {
	return deck.DTRamino
}

func (d *DeckRamino) GetDeckLbl(ixf int) string {
	ix := ixf % 54
	return _deckLbs[ix]
}

func (d *DeckRamino) GetDeckValue(ixf int) int {
	ix := ixf % 54
	if ix >= 52 {
		return 25
	}
	return _points[ix%13]
}

func (d *DeckRamino) GetDeckRank(ixf int) int {
	ix := ixf % 54
	if ix >= 52 {
		return 0
	}
	return _rank[ix%13]
}

func (d *DeckRamino) GetDeckSuit(ixf int) deck.SuitType {
	ix := ixf % 54
	if ix < 13 {
		return deck.Quadri
	}
	if ix < 26 {
		return deck.Fiori
	}
	if ix < 39 {
		return deck.Cuori
	}
	if ix < 52 {
		return deck.Picche
	}
	return deck.Jolly
}
