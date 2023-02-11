package deck

import (
	"log"
	"testing"
)

func TestDoSomeBriscola(t *testing.T) {
	d := Deck{}
	d.InitBriscola()
	log.Println("Briscola deck: ", d.String())
}
