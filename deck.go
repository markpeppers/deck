package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//go:generate stringer -type=Symbol
type Symbol int

const (
	None Symbol = iota
	Spades
	Diamonds
	Clubs
	Hearts
)

//go:generate stringer -type=Rank
type Rank int

const (
	Joker Rank = iota
	Ace
	Two
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
)

type Card struct {
	Value Rank
	Suit  Symbol
}

func New(options ...func(*[]Card)) []Card {
	cards := make([]Card, 0, 52)
	for suit := Spades; suit <= Hearts; suit++ {
		for rank := 1; rank <= 13; rank++ {
			cards = append(cards, Card{Value: Rank(rank), Suit: suit})
		}
	}

	for _, option := range options {
		option(&cards)
	}
	return cards
}

func MultiDeck(n int) func(*[]Card) {
	return func(deck *[]Card) {
		for count := 1; count < n; count++ {
			*deck = append(*deck, New()...)
		}
	}
}

func Shuffle(deck *[]Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*deck), func(i, j int) { (*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i] })
}

func AddJokers(n int) func(*[]Card) {
	return func(deck *[]Card) {
		for i := 0; i < n; i++ {
			*deck = append(*deck, Card{Value: Joker, Suit: None})
		}
	}
}

func RemoveRank(v Rank) func(*[]Card) {
	return func(deck *[]Card) {
		for i := 0; i < len(*deck); {
			if (*deck)[i].Value == v {
				*deck = append((*deck)[:i], (*deck)[i+1:]...)
			} else {
				i++
			}
		}
	}
}

func RemoveRanks(ranks []Rank) func(*[]Card) {
	return func(deck *[]Card) {
		for _, r := range ranks {
			for i := 0; i < len(*deck); {
				if (*deck)[i].Value == r {
					*deck = append((*deck)[:i], (*deck)[i+1:]...)
				} else {
					i++
				}
			}
		}
	}
}

func RemoveSuit(s Symbol) func(*[]Card) {
	return func(deck *[]Card) {
		for i := 0; i < len(*deck); {
			if (*deck)[i].Suit == s {
				*deck = append((*deck)[:i], (*deck)[i+1:]...)
			} else {
				i++
			}
		}
	}
}

func RemoveSuits(suits []Symbol) func(*[]Card) {
	return func(deck *[]Card) {
		for _, s := range suits {
			for i := 0; i < len(*deck); {
				if (*deck)[i].Suit == s {
					*deck = append((*deck)[:i], (*deck)[i+1:]...)
				} else {
					i++
				}
			}
		}
	}
}

func Sort(less func([]Card) func(i, j int) bool) func(*[]Card) {
	return func(deck *[]Card) {
		sort.SliceStable(*deck, less(*deck))
	}
}

func DefaultSort() func(*[]Card) {
	return func(deck *[]Card) {
		sort.SliceStable(*deck, func(i, j int) bool {
			if (*deck)[i].Value == Joker {
				return false
			}
			if (*deck)[i].Value < (*deck)[j].Value {
				return true
			}
			if (*deck)[i].Suit < (*deck)[j].Suit {
				return true
			}
			return false
		})
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}
