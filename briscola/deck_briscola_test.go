package briscola

import (
	"testing"
)

func TestDoSomeBriscola(t *testing.T) {
	d := DeckBriscola{}
	d.Init()

	t.Logf("Briscola deck: %s", d.String())

	card := d.GetCard("_3s")
	if card.Value != 10 {
		t.Errorf("Tre di spade value is 10  instead of %d", card.Value)
	}
}
