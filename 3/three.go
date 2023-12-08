package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Number struct {
	Number int
	xStart int
	xEnd   int
	y      int
}

type Symbol struct {
	y int
	x int
}

func main() {
	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	scan := bufio.NewScanner(f)

	y := 0
	numbers := []Number{}
	symbols := []Symbol{}
	for scan.Scan() {
		line := scan.Text()

		index := 0
		for index < len(line) {
			token := string(line[index])
			if token == "." {
				index++
				continue
			}

			_, err := strconv.Atoi(token)
			// Resolve full int, with start index (index) and end index
			if err == nil {
				fullInt, endIndex := getInt(line, index)

				numbers = append(numbers, Number{
					y:      y,
					xStart: index,
					xEnd:   endIndex,
					Number: fullInt,
				})

				index = endIndex + 1
				continue
			}

			// Symbol
			symbols = append(symbols, Symbol{
				y: y,
				x: index,
			})
			index++
		}

		y++
	}

	res := 0
	for _, number := range numbers {
		isAdjacentToSymbol := false
		for _, symbol := range symbols {

			// Number x y + 1 should envelop symbol
			isNumberInXRange := number.xStart-1 <= symbol.x && symbol.x <= number.xEnd+1
			isNumberInYRange := number.y-1 <= symbol.y && symbol.y <= number.y+1

			if isNumberInYRange && isNumberInXRange {
				isAdjacentToSymbol = true
				break
			}
		}
		if isAdjacentToSymbol {
			fmt.Println(number, "adjacent")
			res += number.Number
		}
		if !isAdjacentToSymbol {
			fmt.Println(number, "not adjacent")
		}
	}

	fmt.Println(res)
}

func getInt(line string, startIndex int) (int, int) {
	index := startIndex
	intstring := ""
	for index < len(line) {
		token := string(line[index])
		_, err := strconv.Atoi(token)
		if err == nil {
			intstring = intstring + token
			index++
			continue
		}
		break
	}

	intt, _ := strconv.Atoi(intstring)
	return intt, index - 1
}
