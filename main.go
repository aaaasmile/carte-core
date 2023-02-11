package main

import (
	"log"

	"github.com/aaaasmile/carte-core/deck"
)

func main() {
	d := deck.Deck{}
	d.InitBriscola()
	log.Println("Briscola deck: ", d.String())

	log.Println("That's all folks")
}
