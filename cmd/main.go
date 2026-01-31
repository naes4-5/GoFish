package main

import (
	"fmt"
	"log"

	"github.com/naes4-5/GoFish/game"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	i := 2
	a = append(a[:i], a[i+i:]...)
	for _, v := range a {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("\nNow we're here")

	b := make([]int, 0, 52)
	for u, v := range b {
		fmt.Printf("%d @ %d\n", v, u)
	}
	fmt.Println("& now we're here")

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
