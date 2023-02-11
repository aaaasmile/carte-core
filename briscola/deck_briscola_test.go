package briscola

import (
	"testing"
)

func TestDoSomeBriscola(t *testing.T) {
	d := DeckBriscola{}
	d.Init()

	t.Logf("Briscola deck: %s", d.String())
}
