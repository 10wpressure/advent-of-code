package main

import (
	"fmt"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
	Nothing
)

func (ht HandType) String() string {
	switch ht {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	default:
		return "Nothing"
	}
}

func (ht HandType) Less(ht2 HandType) bool {
	return int64(ht) < int64(ht2)
}
func (ht HandType) Equal(ht2 HandType) bool {
	return int64(ht) == int64(ht2)

}

type Hand struct {
	cards    []Card
	handType HandType
	bid      int64
	score    int64
}

func NewHand(cards []Card, bid int64) *Hand {
	h := &Hand{
		cards: cards,
		bid:   bid,
	}
	h.GetType()
	return h
}

func (h *Hand) String() string {
	res := ""
	for _, c := range h.cards {
		res += c.String()
	}

	return res
}

func (h *Hand) Print() {
	fmt.Println(h.String())
}

func (h *Hand) GetType() HandType {
	hand := make(map[Card]int)

	for i := 0; i < len(h.cards); i++ {
		hand[h.cards[i]]++
	}

	maxOccurrences := 0
	for _, count := range hand {
		if count > maxOccurrences {
			maxOccurrences = count
		}
	}

	switch {
	case maxOccurrences == 5:
		h.handType = FiveOfAKind
	case maxOccurrences == 4:
		if hand[Joker] == 1 || hand[Joker] == 4 {
			h.handType = FiveOfAKind
		} else {
			h.handType = FourOfAKind
		}
	case maxOccurrences == 3:
		if hand[Joker] == 3 && len(hand) == 3 {
			h.handType = FourOfAKind
		} else if hand[Joker] == 3 && len(hand) == 2 {
			h.handType = FiveOfAKind
		} else if hand[Joker] == 2 {
			h.handType = FiveOfAKind
		} else if hand[Joker] == 1 {
			h.handType = FourOfAKind
		} else if len(hand) == 3 {
			h.handType = ThreeOfAKind
		} else {
			h.handType = FullHouse
		}
	case maxOccurrences == 2:
		switch len(hand) {
		case 4: // 2 1 1 1
			if hand[Joker] >= 1 {
				h.handType = ThreeOfAKind
			} else {
				h.handType = OnePair
			}
		default: // 2 2 1
			if hand[Joker] == 2 {
				h.handType = FourOfAKind
			} else if hand[Joker] == 1 {
				h.handType = FullHouse
			} else {
				h.handType = TwoPair
			}
		}
	case maxOccurrences == 1:
		if hand[Joker] == 1 {
			h.handType = OnePair
		} else {
			h.handType = HighCard
		}
	default:
		h.handType = HighCard
	}

	return h.handType
}
