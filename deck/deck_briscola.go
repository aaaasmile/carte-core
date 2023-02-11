package deck

var _deck40Lbs = []string{"_Ab", "_2b", "_3b", "_4b", "_5b", "_6b", "_7b", "_Fb", "_Cb", "_Rb",
	"_Ac", "_2c", "_3c", "_4c", "_5c", "_6c", "_7c", "_Fc", "_Cc", "_Rc",
	"_Ad", "_2d", "_3d", "_4d", "_5d", "_6d", "_7d", "_Fd", "_Cd", "_Rd",
	"_As", "_2s", "_3s", "_4s", "_5s", "_6s", "_7s", "_Fs", "_Cs", "_Rs"}

var _briscPoints = []int{11, 0, 10, 0, 0, 0, 0, 2, 3, 4}

// rank = { sei: 6, cav: 9, qua: 4, re: 10, set: 7, due: 2, cin: 5, asso: 12, fan: 8, tre: 11 };
var _briscRank = []int{12, 2, 11, 4, 5, 6, 7, 8, 9, 10}

func (d *Deck) InitBriscola() {
	d.deckType = DTBriscola
	d.lblFn = getDeckLblForBriscola
	d.valueFn = getDeckValueForBriscola
	d.rankFn = getDeckRankForBriscola
	d.initialize(40)
}

func getDeckLblForBriscola(ix int) string {
	return _deck40Lbs[ix]
}

func getDeckValueForBriscola(ix int) int {
	return _briscPoints[ix%10]
}

func getDeckRankForBriscola(ix int) int {
	return _briscRank[ix%10]
}
