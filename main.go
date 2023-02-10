package main

import (
	"log"

	"github.com/aaaasmile/carte-core/deck"
)

func main() {
	d := deck.Deck{}
	d.InitToForty()

	log.Println("That's all folks")
}
