package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func convertToIntArray(numStrings []string) ([]int, error) {
	var numbers []int
	for _, numString := range numStrings {
		num, err := strconv.Atoi(numString)
		if err != nil || num < 1 || num > 5 {
			return []int{}, errors.New("入力された数値は有効な数字ではありません。")
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}

func numberIsExist(numbers []int, checkN int) bool {
	for _, num := range numbers {
		if checkN == (num - 1) {
			return true
		}
	}
	return false
}

func getNumbers() []int {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		numStrings := strings.Split(text, " ")
		var err error
		numbers, err = convertToIntArray(numStrings)
		if err != nil {
			fmt.Println(err)
			fmt.Println("1～5の数字で、スペース区切りにしてください。")
			continue
		}
		break
	}
	return numbers
}

func changeIndexes(numbers []int, length int) []int {
	chIndexes := make([]int, 0)
	for i := 0; i < length; i++ {
		if !numberIsExist(numbers, i) {
			chIndexes = append(chIndexes, i)
		}
	}
	return chIndexes
}

func run(deck Deck, counter int) {
	fmt.Printf("=== 第 %d 回目 ===\n", counter)
	deck.Shuffle()
	stage := Stage(deck[0:5])
	stage.Print()
	fmt.Printf("どれを残しますか？ 1～5の数字を入力してください。スペースで区切ることで複数選択できます。\n")
	numbers := getNumbers()
	sort.Ints(numbers)
	for i, idx := range changeIndexes(numbers, len(stage)) {
		stage[idx] = deck[len(stage)+i]
	}

	stage.Print()

	switch stage.CheckHand() {
	case HandRoyalStraightFlush:
		fmt.Println("!! Royal Straight Flush !!")
	case HandStraightFlush:
		fmt.Println("** Straight Flush **")
	case HandFiveCard:
		fmt.Println("++ Five Card ++")
	case HandFourCard:
		fmt.Println("++ Five Card ++")
	case HandFullHouse:
		fmt.Println("-- Full House --")
	case HandFlush:
		fmt.Println("[[ Flush ]]")
	case HandStraight:
		fmt.Println("$$ Straight $$")
	case HandThreeCard:
		fmt.Println("\\\\ Three card \\\\")
	case HandTwoPair:
		fmt.Println("== Two pair ==")
	default:
		fmt.Println("残念、ブタです")
	}

	fmt.Println("エンターキーでもう一回できます。ゲームをやめる場合はCTRL+Cを押してください")
	fmt.Scanln()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	deck := MakeDeck()
	for i := 1; ; i++ {
		run(deck, i)
	}
}
