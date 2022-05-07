package deck

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
	return nil
}

// deck implements Deck.
type deck struct {
	cards []Card
	drawn []Card
}

// addCard stores a card in the deck.
func (d *deck) addCard(c Card) {

}

// Count returns the total number of cards in the deck.
func (d *deck) Count() int {
	return 0
}

// RemainingCount returns the number of undrawn cards.
func (d *deck) RemainingCount() int {
	return 0
}

// Draw returns a single card. An error is returnd if
// there are no more undrawn cards.
func (d *deck) Draw() (Card, error) {
	return nil, nil
}

// DrawX returns a number of cards equal to amount.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) DrawX(amount int) ([]Card, error) {
	return nil, nil
}

// Discard removes a card from the undrawn cards.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) Discard() error {
	return nil
}

// DiscardX removes undrawn cards equal to amount.
// An error is returned if there are less cards
// remaining than amount.
func (d *deck) DiscardX(amount int) error {
	return nil
}

// Shuffle randomizes the undrawn cards amount times.
// Does nothing when amount is less than or equal to zero.
func (d *deck) Shuffle(amount int) {

}

// FullShuffle combines the drawn and undrawn cards, then shuffles them.
func (d *deck) FullShuffle() {

}
