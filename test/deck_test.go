package game_test

import (
	"testing"

	"github.com/naes4-5/GoFish/game"
)

func TestCard_GetSuit(t *testing.T) {
	tests := []struct {
		name    string
		suit    game.Suit
		want    string
		wantErr bool
	}{
		{"spades", game.Spades, "Spades", false},
		{"hearts", game.Hearts, "Hearts", false},
		{"clubs", game.Clubs, "Clubs", false},
		{"diamonds", game.Diamonds, "Diamonds", false},
		{"invalid", game.Suit(99), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &game.Card{Suit: tt.suit}
			got, err := c.GetSuit()
			if (err != nil) != tt.wantErr {
				t.Errorf("Card.GetSuit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Card.GetSuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeDeck(t *testing.T) {
	d := game.MakeDeck()
	if len(d.Cards) != 52 {
		t.Errorf("MakeDeck() returned deck with %d cards, want 52", len(d.Cards))
	}

	// Verify suit distribution
	suits := make(map[game.Suit]int)
	for _, c := range d.Cards {
		suits[c.Suit]++
	}

	expectedSuits := []game.Suit{game.Spades, game.Hearts, game.Clubs, game.Diamonds}
	for _, s := range expectedSuits {
		if suits[s] != 13 {
			t.Errorf("MakeDeck() suit %v has %d cards, want 13", s, suits[s])
		}
	}
}

func TestDeck_DrawCard(t *testing.T) {
	t.Run("draw from empty deck", func(t *testing.T) {
		d := &game.Deck{Cards: []game.Card{}}
		_, err := d.DrawCard()
		if err == nil {
			t.Error("DrawCard() expected error for empty deck, got nil")
		}
	})

	t.Run("draw until empty", func(t *testing.T) {
		d := game.MakeDeck()
		initialCount := len(d.Cards)

		for i := range initialCount {
			_, err := d.DrawCard()
			if err != nil {
				t.Fatalf("DrawCard() failed at draw %d: %v", i+1, err)
			}
			if len(d.Cards) != initialCount-i-1 {
				t.Errorf("Deck length after %d draws = %d, want %d", i+1, len(d.Cards), initialCount-i-1)
			}
		}

		if d.HasCardsLeft() {
			t.Error("HasCardsLeft() returned true for empty deck")
		}

		_, err := d.DrawCard()
		if err == nil {
			t.Error("DrawCard() expected error after deck exhausted, got nil")
		}
	})
}

func TestDeck_HasCardsLeft(t *testing.T) {
	tests := []struct {
		name  string
		cards []game.Card
		want  bool
	}{
		{"empty", []game.Card{}, false},
		{"one card", []game.Card{{Rank: 1, Suit: game.Spades}}, true},
		{"full deck", game.MakeDeck().Cards, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &game.Deck{Cards: tt.cards}
			if got := d.HasCardsLeft(); got != tt.want {
				t.Errorf("Deck.HasCardsLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
