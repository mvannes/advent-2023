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

type Gear struct {
	Symbol          Symbol
	AdjacentNumbers []Number
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
	gears := []Gear{}
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
			s := Symbol{
				y: y,
				x: index,
			}
			symbols = append(symbols, s)
			if token == "*" {
				gears = append(gears, Gear{
					Symbol:          s,
					AdjacentNumbers: nil,
				})
			}
			index++
		}

		y++
	}

	partNumbers := []Number{}
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
			partNumbers = append(partNumbers, number)
		}
	}

	// I could combine this with the previous loop, but cba.
	for i, gear := range gears {
		for _, number := range partNumbers {
			isNumberInXRange := number.xStart-1 <= gear.Symbol.x && gear.Symbol.x <= number.xEnd+1
			isNumberInYRange := number.y-1 <= gear.Symbol.y && gear.Symbol.y <= number.y+1

			if isNumberInYRange && isNumberInXRange {
				gear.AdjacentNumbers = append(gear.AdjacentNumbers, number)
			}
		}
		gears[i] = gear
	}

	fmt.Println(gears)

	// This loop is unnecessary.
	res := 0
	for _, gear := range gears {
		if len(gear.AdjacentNumbers) != 2 {
			continue
		}
		gearRes := 1
		for _, n := range gear.AdjacentNumbers {
			gearRes *= n.Number
		}
		res += (gearRes)
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
