package main

import "log"

type Card int

const (
	None Card = iota
	Joker
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
	Ace
)

func FromString(r rune) Card {
	switch r {
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return Ten
	case 'J':
		return Jack
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		log.Fatalf("invalid card: %c", r)
		return None
	}
}

func FromStringPart2(r rune) Card {
	switch r {
	case 'J':
		return Joker
	case '2':
		return Two
	case '3':
		return Three
	case '4':
		return Four
	case '5':
		return Five
	case '6':
		return Six
	case '7':
		return Seven
	case '8':
		return Eight
	case '9':
		return Nine
	case 'T':
		return Ten
	case 'Q':
		return Queen
	case 'K':
		return King
	case 'A':
		return Ace
	default:
		log.Fatalf("invalid card: %c", r)
		return None
	}
}

func (c *Card) Value() int {
	switch *c {
	case Joker:
		return 1
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 11
	case Queen:
		return 12
	case King:
		return 13
	case Ace:
		return 14
	default:
		log.Fatalf("invalid card: %d", *c)
		return 0
	}
}

func (c *Card) String() string {
	switch *c {
	case Joker:
		return "J"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "T"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		log.Fatalf("invalid card: %d", *c)
		return ""
	}
}
