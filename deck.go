package main

import "math/rand"

type Deck []Card

func MakeDeck() Deck {
	deck := make([]Card, 0)
	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			deck = append(deck, Card{i, j})
		}
	}
	deck = append(deck, Card{MarkJoker, NumberJoker})
	return deck
}

func (deck *Deck) Shuffle() {
	n := len(*deck)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	}
}
