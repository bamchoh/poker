package main

import (
	"fmt"
	"testing"
)

func TestIsFlush(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, 1},
			Card{MarkHeart, 9},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkHeart, 5},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkHeart, 2},
			Card{MarkHeart, 9},
			Card{MarkHeart, 4},
			Card{MarkHeart, 5},
		},
		Stage{
			Card{MarkClover, 9},
			Card{MarkClover, 3},
			Card{MarkJoker, NumberJoker},
			Card{MarkClover, 4},
			Card{MarkClover, 5},
		},
		Stage{
			Card{MarkSpade, 2},
			Card{MarkSpade, 3},
			Card{MarkSpade, 9},
			Card{MarkSpade, 5},
			Card{MarkJoker, NumberJoker},
		},
	}

	for _, stage := range stagetests {
		if stage.CheckHand() != HandFlush {
			t.Error("does not same mark")
		}
	}
}

func TestIsNotSameMark(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, 1},
			Card{MarkHeart, 2},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkSpade, 5},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkHeart, 2},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkSpade, 5},
		},
		Stage{
			Card{MarkHeart, 2},
			Card{MarkClover, 3},
			Card{MarkJoker, NumberJoker},
			Card{MarkSpade, 4},
			Card{MarkDia, 5},
		},
		Stage{
			Card{MarkClover, 2},
			Card{MarkSpade, 3},
			Card{MarkSpade, 4},
			Card{MarkSpade, 5},
			Card{MarkJoker, NumberJoker},
		},
	}

	for _, stage := range stagetests {
		if stage.CheckHand() == HandFlush {
			t.Error("does not same mark")
		}
	}
}

func TestIsRoyalStraightFlush(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, 10},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
	}

	for _, stage := range stagetests {
		if stage.CheckHand() != HandRoyalStraightFlush {
			t.Error("not royal straight flush")
		}
	}
}

func TestIsNotRoyalStraightFlush(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 10},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkDia, 10},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkDia, 2},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 2},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 9},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkHeart, 9},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 2},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkJoker, NumberJoker},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() == HandRoyalStraightFlush {
			errstr := fmt.Sprintf("%d is royal straight flush", i)
			t.Error(errstr)
		}
	}
}

func TestIsStraightFlush(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, NumberAce},
			Card{MarkHeart, 2},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkHeart, 5},
		},
		Stage{
			Card{MarkClover, NumberKing},
			Card{MarkClover, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkClover, 10},
			Card{MarkClover, 9},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() != HandStraightFlush {
			errstr := fmt.Sprintf("%d is not straight flush", i)
			t.Error(errstr)
		}
	}
}

func TestIsNotStraightFlush(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkClover, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkClover, 10},
			Card{MarkClover, 9},
		},
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, 2},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkHeart, 5},
		},
		Stage{
			Card{MarkHeart, NumberAce},
			Card{MarkHeart, 2},
			Card{MarkHeart, 3},
			Card{MarkHeart, 4},
			Card{MarkHeart, 6},
		},
		Stage{
			Card{MarkClover, NumberKing},
			Card{MarkClover, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkClover, 10},
			Card{MarkClover, 8},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkClover, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkClover, 10},
			Card{MarkClover, 7},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() == HandStraightFlush {
			errstr := fmt.Sprintf("%d is straight flush", i)
			t.Error(errstr)
		}
	}
}

func TestIsFiveCard(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, NumberAce},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkDia, 9},
			Card{MarkSpade, 9},
			Card{MarkClover, 9},
			Card{MarkHeart, 9},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkClover, NumberKing},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberKing},
			Card{MarkDia, NumberKing},
			Card{MarkJoker, NumberJoker},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() != HandFiveCard {
			errstr := fmt.Sprintf("%d is not five card", i)
			t.Error(errstr)
		}
	}
}

func TestIsNotFiveCard(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, NumberAce},
			Card{MarkJoker, 2},
		},
		Stage{
			Card{MarkDia, 8},
			Card{MarkSpade, 9},
			Card{MarkClover, 9},
			Card{MarkHeart, 9},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkClover, NumberKing},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberKing},
			Card{MarkDia, NumberQueen},
			Card{MarkJoker, NumberJoker},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() == HandFiveCard {
			errstr := fmt.Sprintf("%d is five card", i)
			t.Error(errstr)
		}
	}
}

func TestIsFourCard(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, 2},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, NumberAce},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkDia, 9},
			Card{MarkSpade, 9},
			Card{MarkClover, 9},
			Card{MarkHeart, 8},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkClover, NumberKing},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberKing},
			Card{MarkDia, NumberKing},
			Card{MarkJoker, NumberQueen},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() != HandFourCard {
			errstr := fmt.Sprintf("%d is not four card", i)
			t.Error(errstr)
		}
	}
}

func TestIsNotFourCard(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, NumberAce},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkDia, 9},
			Card{MarkSpade, 9},
			Card{MarkClover, 8},
			Card{MarkHeart, 8},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkClover, NumberQueen},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberKing},
			Card{MarkDia, NumberKing},
			Card{MarkJoker, NumberQueen},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() == HandFourCard {
			errstr := fmt.Sprintf("%d is four card", i)
			t.Error(errstr)
		}
	}
}

func TestIsFullHouse(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, 2},
			Card{MarkHeart, 2},
		},
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, 2},
			Card{MarkSpade, 2},
			Card{MarkJoker, NumberJoker},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() != HandFullHouse {
			errstr := fmt.Sprintf("%d is not full house", i)
			t.Error(errstr)
		}
	}
}

func TestIsNotFullHouse(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, 4},
			Card{MarkSpade, 2},
			Card{MarkHeart, 2},
		},
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, 2},
			Card{MarkHeart, 3},
		},
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, NumberAce},
			Card{MarkDia, NumberAce},
			Card{MarkSpade, 2},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkClover, NumberAce},
			Card{MarkHeart, 3},
			Card{MarkDia, 2},
			Card{MarkSpade, 4},
			Card{MarkJoker, NumberJoker},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() == HandFullHouse {
			errstr := fmt.Sprintf("%d is full house", i)
			t.Error(errstr)
		}
	}
}

func TestIsStraight(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkHeart, NumberJack},
			Card{MarkHeart, NumberQueen},
			Card{MarkHeart, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 10},
			Card{MarkDia, NumberJack},
			Card{MarkClover, NumberQueen},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkDia, NumberJack},
			Card{MarkClover, NumberQueen},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkHeart, 10},
			Card{MarkDia, NumberJack},
			Card{MarkClover, NumberQueen},
			Card{MarkSpade, NumberKing},
			Card{MarkJoker, NumberJoker},
		},
		Stage{
			Card{MarkHeart, NumberAce},
			Card{MarkDia, 2},
			Card{MarkClover, 3},
			Card{MarkSpade, 4},
			Card{MarkHeart, 5},
		},
		Stage{
			Card{MarkHeart, NumberKing},
			Card{MarkDia, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkSpade, 10},
			Card{MarkHeart, 9},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkDia, NumberQueen},
			Card{MarkClover, NumberJack},
			Card{MarkHeart, 10},
			Card{MarkSpade, 9},
		},
	}

	for _, stage := range stagetests {
		if stage.CheckHand() != HandStraight {
			t.Error("does not same mark")
		}
	}
}

func TestIsThreeCard(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, 10},
			Card{MarkDia, 10},
			Card{MarkClover, 10},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberAce},
		},
		Stage{
			Card{MarkJoker, NumberJoker},
			Card{MarkDia, 10},
			Card{MarkClover, 10},
			Card{MarkSpade, NumberKing},
			Card{MarkHeart, NumberAce},
		},
	}

	for i, stage := range stagetests {
		if stage.CheckHand() != HandThreeCard {
			errstr := fmt.Sprintf("%d is not three card.", i)
			t.Error(errstr)
		}
	}
}

func TestIsTowPair(t *testing.T) {
	stagetests := []Stage{
		Stage{
			Card{MarkHeart, 10},
			Card{MarkDia, 10},
			Card{MarkClover, 9},
			Card{MarkSpade, 9},
			Card{MarkHeart, NumberAce},
		},
	}

	for i, stage := range stagetests {
		hand := stage.CheckHand()
		if hand != HandTwoPair {
			errstr := fmt.Sprintf("%d is not two pair.", i)
			t.Error(errstr)
		}
	}
}
