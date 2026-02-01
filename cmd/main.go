package main

import (
	"fmt"
	"log"

	"github.com/naes4-5/GoFish/game"
)

func main() {
	b := make([]int, 0, 52)
	for u, v := range b {
		fmt.Printf("%d @ %d\n", v, u)
	}
	fmt.Println()

	d := game.MakeDeck()
	for u, v := range d.Cards {
		s, err := v.GetSuit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %d of %s\n", u, v.Rank, s)
	}
	for d.HasCardsLeft() {
		c, err := d.DrawCard()
		if err != nil {
			log.Fatal(err)
		}
		s, err := c.GetSuit()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d of %s\n", c.Rank, s)
	}

}
