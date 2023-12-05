package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	Blue  int
	Green int
	Red   int
}
type Game struct {
	GameNumber int
	Sets       []Set
}

func main() {

	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	scan := bufio.NewScanner(f)

	//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	//Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	//Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	//Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	//Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

	games := []Game{}
	for scan.Scan() {
		line := scan.Text()
		var gameNumber int

		splitted := strings.Split(line, ":")
		gameString := splitted[0]
		fmt.Sscanf(gameString, "Game %d:", &gameNumber)

		setsString := splitted[1]

		sets := strings.Split(setsString, ";")

		setsInGame := []Set{}
		for _, setString := range sets {
			set := Set{
				Blue:  0,
				Green: 0,
				Red:   0,
			}

			for _, item := range strings.Split(setString, ",") {
				itemSplit := strings.Split(strings.TrimSpace(item), " ")
				amount, _ := strconv.Atoi(itemSplit[0])
				color := itemSplit[1]
				switch color {
				case "red":
					set.Red = amount
				case "green":
					set.Green = amount
				case "blue":
					set.Blue = amount
				}
			}
			setsInGame = append(setsInGame, set)
		}

		games = append(games, Game{
			GameNumber: gameNumber,
			Sets:       setsInGame,
		})
	}

	result := 0
	for _, game := range games {
		gameIsPossible := true
		for _, set := range game.Sets {
			if set.Red > 12 || set.Blue > 14 || set.Green > 13 {
				gameIsPossible = false
				break
			}
		}
		if !gameIsPossible {
			continue
		}

		result += game.GameNumber
	}

	fmt.Println(result)
}
