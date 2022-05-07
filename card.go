package deck

import "fmt"

type Suit int

const (
	Clubs Suit = iota + 1
	Diamonds
	Hearts
	Spades
	JokerSuit
)

func (s Suit) String() string {
	switch s {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	default:
		return "Joker"
	}
}

type Value int

const (
	Two Value = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	JokerValue
)

func (v Value) String() string {
	switch v {
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	case Ace:
		return "Ace"
	default:
		return "Joker"
	}
}

var (
	// Values is all the card values in a normal playing card deck.
	Values = []Value{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	// Suits is all the card suits in a normal playing card deck.
	Suits = []Suit{Clubs, Diamonds, Hearts, Spades}
)

// Card represents a card from a regular playing card deck.
type Card interface {
	String() string
	IsJoker() bool
}

type card struct {
	Value
	Suit
}

// NewCard returns a card from a normal deck with the
// specified value and suit.
func NewCard(v Value, s Suit) Card {
	return &card{Value: v, Suit: s}
}

// NewJoker returns a joker card.
func NewJoker() Card {
	return NewCard(JokerValue, JokerSuit)
}

// String returns the string value of this card.
func (c *card) String() string {
	if c.IsJoker() {
		return "Joker"
	}

	return fmt.Sprintf("%v of %v", c.Value, c.Suit)
}

// IsJoker checks whether this card is a joker.
func (c *card) IsJoker() bool {
	return c.Value == JokerValue || c.Suit == JokerSuit
}
