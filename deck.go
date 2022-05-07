package deck

import (
	"errors"
	"math/rand"
)

var (
	ErrNoCardsRemaining        = errors.New("no cards remaining")
	ErrNotEnoughCardsRemaining = errors.New("not enough cards remaining")
)

// Deck holds a regular set of cards and manupulates them.
type Deck interface {
	Count() int
	RemainingCount() int
	Draw() (Card, error)
	DrawX(amount int) ([]Card, error)
	Discard() error
	DiscardX(amount int) error
	Shuffle(amount int)
	FullShuffle()
}

// NewDeck returns a normal deck of playing cards
// with the specified number of jokers.
func NewDeck(jokers int) Deck {
	d := &deck{}
	for _, suit := range Suits {
		for _, value := range Values {
			d.addCard(NewCard(value, suit))
		}
	}

	for i := 0; i < jokers; i++ {
		d.addCard(NewJoker())
	}

	return d
}

// deck implements Deck.
type deck struct {
	cards []Card
	drawn []Card
}

// addCard stores a card in the deck.
func (d *deck) addCard(c Card) {
	d.cards = append(d.cards, c)
}

// Count returns the total number of cards in the deck.
func (d *deck) Count() int {
	return len(d.cards) + len(d.drawn)
}

// RemainingCount returns the number of undrawn cards.
func (d *deck) RemainingCount() int {
	return len(d.cards)
}

// removeCard moves the top card to the drawn list.
func (d *deck) removeCard() Card {
	c := d.cards[0]
	d.drawn = append(d.drawn, c)
	d.cards = d.cards[1:]
	return c
}

// removeCards moves the top amount of cards to the drawn list.
func (d *deck) removeCards(amount int) []Card {
	c := d.cards[:amount]
	d.drawn = append(d.drawn, c...)
	d.cards = d.cards[amount:]
	return c
}

// Draw returns a single card. An error is returnd if
// there are no more undrawn cards.
func (d *deck) Draw() (Card, error) {
	if d.RemainingCount() == 0 {
		return nil, ErrNoCardsRemaining
	}
	return d.removeCard(), nil
}

// DrawX returns a number of cards equal to amount.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) DrawX(amount int) ([]Card, error) {
	if d.RemainingCount() < amount {
		return nil, ErrNotEnoughCardsRemaining
	}
	return d.removeCards(amount), nil
}

// Discard removes a card from the undrawn cards.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) Discard() error {
	if d.RemainingCount() == 0 {
		return ErrNoCardsRemaining
	}
	_ = d.removeCard()
	return nil
}

// DiscardX removes undrawn cards equal to amount.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) DiscardX(amount int) error {
	if d.RemainingCount() < amount {
		return ErrNotEnoughCardsRemaining
	}
	_ = d.removeCards(amount)
	return nil
}

// Shuffle randomizes the undrawn cards amount times.
// Does nothing when amount is less than or equal to zero.
func (d *deck) Shuffle(amount int) {
	if amount < 0 {
		amount = 0
	}

	for i := 0; i < amount; i++ {
		var drawn []Card

		for d.RemainingCount() != 0 {
			randomIndex := rand.Intn(d.RemainingCount())
			c := d.cards[randomIndex]
			drawn = append(drawn, c)
			d.cards = append(d.cards[0:randomIndex], d.cards[randomIndex+1:]...)
		}

		d.cards = drawn
	}
}

// FullShuffle combines the drawn and undrawn cards, then shuffles them.
func (d *deck) FullShuffle() {
	d.cards = append(d.cards, d.drawn...)
	d.drawn = nil

	d.Shuffle(10)
}
