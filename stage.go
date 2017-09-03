package main

import (
	"fmt"
	"sort"
)

const (
	HandRoyalStraightFlush = iota
	HandStraightFlush
	HandFiveCard
	HandFourCard
	HandFullHouse
	HandFlush
	HandStraight
	HandThreeCard
	HandTwoPair
)

type Stage []Card

func (s *Stage) Print() {
	for i := 0; i < len(*s); i++ {
		fmt.Printf("%d | %6s | %2s\n", i+1, (*s)[i].MarkString(), (*s)[i].NumberString())
	}
}

func (s *Stage) CheckHand() int {
	switch {
	case s.IsRoyalStraightFlush():
		return HandRoyalStraightFlush
	case s.IsStraightFlush():
		return HandStraightFlush
	case s.IsFiveCard():
		return HandFiveCard
	case s.IsFourCard():
		return HandFourCard
	case s.IsFullHouse():
		return HandFullHouse
	case s.IsFlush():
		return HandFlush
	case s.IsStraight():
		return HandStraight
	case s.IsThreeCard():
		return HandThreeCard
	case s.IsTwoPair():
		return HandTwoPair
	}
	return -1
}

func (s *Stage) IsFlush() bool {
	for i := 0; i < len(*s)-1; i++ {
		if (*s)[i].Mark() == MarkJoker || (*s)[i+1].Mark() == MarkJoker {
			continue
		}

		if (*s)[i].Mark() != (*s)[i+1].Mark() {
			return false
		}
	}
	return true
}

func (s *Stage) checkStraight(expectedNumbers []int) bool {
	foundJoker := false
	usedJoker := false
	for _, num := range expectedNumbers {
		var found bool
		for _, card := range *s {
			if num == card.Number() {
				found = true
				break
			}

			if card.Number() == NumberJoker {
				foundJoker = true
				continue
			}
		}

		if !found {
			if foundJoker && !usedJoker {
				usedJoker = true
			} else {
				return false
			}
		}
	}
	return true
}

func (s *Stage) isSpecialCaseOfStraight() bool {
	expectedNumbers := []int{
		10,
		NumberKing,
		NumberQueen,
		NumberJack,
		NumberAce,
	}

	return s.checkStraight(expectedNumbers)
}

func (s *Stage) isTypicallyStraight() bool {
	min := NumberJoker
	for _, card := range *s {
		if card.Number() < min {
			min = card.Number()
		}
	}

	if min > 9 {
		return false
	}

	expectedNumbers := []int{
		min,
		min + 1,
		min + 2,
		min + 3,
		min + 4,
	}

	return s.checkStraight(expectedNumbers)
}

func (s *Stage) IsRoyalStraightFlush() bool {
	if !s.IsFlush() {
		return false
	}

	return s.isSpecialCaseOfStraight()
}

func (s *Stage) IsStraightFlush() bool {
	if !s.IsFlush() {
		return false
	}

	return s.isTypicallyStraight()
}

func (s *Stage) IsFiveCard() bool {
	for i := 0; i < len(*s)-1; i++ {
		if (*s)[i].Number() == NumberJoker || (*s)[i+1].Number() == NumberJoker || ((*s)[i].Number() == (*s)[i+1].Number()) {
			continue
		}
		return false
	}
	return true
}

func (s *Stage) sort() {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i].Number() < (*s)[j].Number()
	})
}

func newcards(cards []Card, delcards []Card) []Card {
	newCards := make([]Card, 0)
	for _, card := range cards {
		hit := false
		for _, delcard := range delcards {
			if card.Number() == delcard.Number() && card.Mark() == delcard.Mark() {
				hit = true
				break
			}
		}

		if !hit {
			newCards = append(newCards, card)
		}
	}
	return newCards
}

func (s *Stage) isNCard(n int, cards []Card) (bool, []Card) {
	counter := map[int][]Card{}
	for _, card := range cards {
		key := card.Number()
		if _, ok := counter[key]; ok == false {
			counter[key] = make([]Card, 0)
		}
		counter[key] = append(counter[key], card)
	}

	for k, v := range counter {
		if k == NumberJoker {
			continue
		}

		if len(v) == n {
			newCards := newcards(cards, v)
			return true, newCards
		}

		if len(v) == n-1 {
			if _, ok := counter[NumberJoker]; ok == true {
				newCards := newcards(cards, v)
				newCards = newcards(newCards, counter[NumberJoker])
				return true, newCards
			}
		}
	}
	return false, cards
}

func (s *Stage) IsFourCard() bool {
	ret, _ := s.isNCard(4, []Card(*s))
	return ret
}

func (s *Stage) IsThreeCard() bool {
	ret, _ := s.isNCard(3, []Card(*s))
	return ret
}

func (s *Stage) IsFullHouse() bool {
	if ret, newCards := s.isNCard(3, []Card(*s)); ret {
		if ret, _ := s.isNCard(2, newCards); ret == true {
			return ret
		}
	}
	return false
}

func (s *Stage) IsStraight() bool {
	if s.isTypicallyStraight() {
		return true
	} else {
		if s.isSpecialCaseOfStraight() {
			return true
		}
	}
	return false
}

func (s *Stage) IsTwoPair() bool {
	pairs := map[int][]Card{}

	for i := 0; i < len(*s); i++ {
		key := (*s)[i].Number()
		if _, ok := pairs[key]; ok == false {
			pairs[key] = make([]Card, 0)
		}
		pairs[key] = append(pairs[key], (*s)[i])
	}

	counter := 0
	usedJoker := false
	for k, v := range pairs {
		if k == NumberJoker {
			continue
		}

		if len(v) == 2 {
			counter++
		}

		if len(v) == 1 {
			if _, ok := pairs[NumberJoker]; ok && !usedJoker {
				usedJoker = true
				counter++
			}
		}
	}

	if counter == 2 {
		return true
	}
	return false
}
