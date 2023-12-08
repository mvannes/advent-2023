package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winningNumbers []int
	myNumbers      []int
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	scan := bufio.NewScanner(f)

	cards := []Card{}
	for scan.Scan() {
		line := scan.Text()
		cardAndNumbers := strings.Split(line, ":")
		winningAndMyNumbers := strings.Split(cardAndNumbers[1], "|")
		winning := strings.Split(winningAndMyNumbers[0], " ")
		my := strings.Split(winningAndMyNumbers[1], " ")

		card := Card{
			winningNumbers: nil,
			myNumbers:      nil,
		}

		for _, w := range winning {
			if w == "" {
				continue
			}
			wnum, err := strconv.Atoi(w)
			if err != nil {
				log.Fatal(err)
			}

			card.winningNumbers = append(card.winningNumbers, wnum)
		}
		for _, m := range my {
			if m == "" {
				continue
			}
			mnum, err := strconv.Atoi(m)
			if err != nil {
				log.Fatal(err)
			}

			card.myNumbers = append(card.myNumbers, mnum)
		}

		cards = append(cards, card)
	}

	totalScore := 0
	for _, card := range cards {
		winningNumbers := 0
		for _, my := range card.myNumbers {
			for _, win := range card.winningNumbers {
				if win == my {
					winningNumbers++
					break
				}
			}
		}

		if winningNumbers == 0 {
			continue
		}
		cardScore := 1
		for i := 1; i < winningNumbers; i++ {
			cardScore = cardScore * 2
		}
		totalScore += cardScore

	}

	fmt.Println(totalScore)
}
