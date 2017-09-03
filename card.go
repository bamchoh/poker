package main

import "strconv"

type Card struct {
	mark   int
	number int
}

func (c *Card) Mark() int {
	return c.mark
}

func (c *Card) Number() int {
	return c.number
}

func (c *Card) MarkString() string {
	switch c.mark {
	case MarkHeart:
		return "Heart"
	case MarkSpade:
		return "Spade"
	case MarkDia:
		return "Dia"
	case MarkClover:
		return "Clover"
	case MarkJoker:
		return "*"
	default:
		return ""
	}
}

func (c *Card) NumberString() string {
	switch c.number {
	case NumberJack:
		return "J"
	case NumberQueen:
		return "Q"
	case NumberKing:
		return "K"
	case NumberAce:
		return "A"
	case NumberJoker:
		return "Joker"
	}

	return strconv.Itoa(c.number)
}
