package briscola

import (
	"testing"

	"github.com/aaaasmile/carte-core/deck"
)

func TestDoSomeBriscola(t *testing.T) {
	d := DeckBriscola{}
	d.Init()

	t.Logf("Briscola deck: %s", d.String())

	card := d.GetCard("_3s")
	if card.Value() != 10 {
		t.Errorf("Tre di spade value is %d  instead of 10", card.Value())
	}

	card = d.GetCard("_2c")
	if card.Suit() != deck.Coppe {
		t.Errorf("Due di coppe has suit  %d instead of coppe", card.Suit())
	}
}
