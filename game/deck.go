package game

import (
	"errors"
	"math/rand/v2"
)

type Suit int

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

var suitName = map[Suit]string{
	Spades:   "Spades",
	Hearts:   "Hearts",
	Clubs:    "Clubs",
	Diamonds: "Diamonds",
}

type Card struct {
	Rank int8
	Suit Suit
}

// returns the suit of the current card, or an error if the suit is invalid
func (c *Card) GetSuit() (string, error) {
	suit, ok := suitName[c.Suit]
	if !ok {
		return "", errors.New("invalid suit for card")
	}
	return suit, nil
}

type Deck struct {
	Cards []Card
}

// returns a fresh deck
func MakeDeck() *Deck {
	d := Deck{
		Cards: make([]Card, 0, 52),
	}
	for i := range int8(13) {
		for suit := range suitName {
			d.Cards = append(d.Cards, Card{
				Rank: i,
				Suit: suit,
			})
		}
	}
	return &d
}

// returns true if there are cards remaining in the deck
func (d *Deck) HasCardsLeft() bool {
	return len(d.Cards) > 0
}

// returns a random card from the deck, or an error if there are no more cards
// to draw
func (d *Deck) DrawCard() (Card, error) {
	if !d.HasCardsLeft() {
		return Card{}, errors.New("no cards left to draw")
	}
	i := rand.IntN(len(d.Cards))
	c := d.Cards[i]
	if i == len(d.Cards)-1 {
		d.Cards = d.Cards[:i]
	} else {
		d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
	}
	return c, nil
}
