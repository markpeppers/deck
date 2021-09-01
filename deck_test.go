package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	deck := New()
	assert.Equal(t, 52, len(deck), "Should have 52 cards")
	assert.Equal(t, "Ace of Spades", fmt.Sprintf("%s", deck[0]), "First card is Ace of Spades")
	assert.Equal(t, "King of Hearts", fmt.Sprintf("%s", deck[51]), "Last card is King of Hearts")
}

func TestMulti(t *testing.T) {
	deck := New(MultiDeck(4))
	assert.Equal(t, 52*4, len(deck), "Should have 4 full decks")
	assert.Equal(t, "Ace of Spades", fmt.Sprintf("%s", deck[52]), "First card of 2nd deck is Ace of Spades")
	assert.Equal(t, "King of Hearts", fmt.Sprintf("%s", deck[52*4-1]), "Last card of 4th deck is King of Hearts")
}

func TestShuffle(t *testing.T) {
	deck := New(Shuffle)
	assert.Equal(t, 52, len(deck), "Shuffled deck should have 52 cards")
}

func TestJokers(t *testing.T) {
	deck := New(AddJokers(2))
	assert.Equal(t, "Joker of None", fmt.Sprintf("%s", deck[52]), "First Joker should be present")
	assert.Equal(t, "Joker of None", fmt.Sprintf("%s", deck[53]), "Second Joker should be present")
	shuffledDeck := New(AddJokers(2), Shuffle)
	jokerCount := 0
	for _, card := range shuffledDeck {
		if card.Value == Joker {
			jokerCount++
		}
	}
	assert.Equal(t, 2, jokerCount, "Shuffled deck should contain 2 Jokers")
}

func TestRemoveRank(t *testing.T) {
	deck := New(RemoveRank(Two), RemoveRank(Three))
	count := 0
	for _, c := range deck {
		if c.Value == Two || c.Value == Three {
			count++
		}
	}
	assert.Equal(t, 0, count, "Twos and Threes should be removed")
	assert.Equal(t, 52-4*2, len(deck), "Deck should have all cards but Twos and Threes")
}

func TestRemoveRanks(t *testing.T) {
	deck := New(RemoveRanks([]Rank{King, Queen, Jack}))
	count := 0
	for _, c := range deck {
		if c.Value == King || c.Value == Queen || c.Value == Jack {
			count++
		}
	}
	assert.Equal(t, 0, count, "Face cards should be removed")
	assert.Equal(t, 52-4*3, len(deck), "Deck should have all non-face cards")
}

func TestRemoveSuit(t *testing.T) {
	deck := New(RemoveSuit(Clubs))
	count := 0
	for _, c := range deck {
		if c.Suit == Clubs {
			count++
		}
	}
	assert.Equal(t, 0, count, "Clubs should be removed")
	assert.Equal(t, 52-13, len(deck), "No non-Clubs should be removed")
}

func TestRemoveSuits(t *testing.T) {
	deck := New(RemoveSuits([]Symbol{Diamonds, Hearts}))
	count := 0
	for _, c := range deck {
		if c.Suit == Diamonds || c.Suit == Hearts {
			count++
		}
	}
	assert.Equal(t, 0, count, "All Red suits should be removed")
	assert.Equal(t, 52-13*2, len(deck), "All black suits should remain")
}
