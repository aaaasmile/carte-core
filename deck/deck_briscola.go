package deck

var _deck40Lbs = []string{"_Ab", "_2b", "_3b", "_4b", "_5b", "_6b", "_7b", "_Fb", "_Cb", "_Rb",
	"_Ac", "_2c", "_3c", "_4c", "_5c", "_6c", "_7c", "_Fc", "_Cc", "_Rc",
	"_Ad", "_2d", "_3d", "_4d", "_5d", "_6d", "_7d", "_Fd", "_Cd", "_Rd",
	"_As", "_2s", "_3s", "_4s", "_5s", "_6s", "_7s", "_Fs", "_Cs", "_Rs"}

func (d *Deck) InitToForty() {
	d.deckType = DTBriscola
	d.LblFn = getDeckLblForBriscola
	d.Items = make([]DeckItem, 40)

}

func getDeckLblForBriscola(ix int) string {
	return _deck40Lbs[ix]
}
