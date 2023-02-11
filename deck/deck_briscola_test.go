package deck

import (
	"testing"
)

func TestDoSomeBriscola(t *testing.T) {
	d := Deck{}
	d.InitBriscola()

	t.Logf("Briscola deck: %s", d.String())
}
