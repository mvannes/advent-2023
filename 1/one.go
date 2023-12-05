package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	scan := bufio.NewScanner(f)

	numberStringsToNumberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	result := 0
	for scan.Scan() {
		line := scan.Text()

		numbers := []string{}
		for index, token := range line {
			stringToken := string(token)
			_, err := strconv.Atoi(stringToken)
			if nil == err {
				numbers = append(numbers, stringToken)
				continue
			}

			for numberAsString, number := range numberStringsToNumberMap {
				if index+len(numberAsString) > len(line) {
					continue
				}
				if numberAsString == line[index:index+len(numberAsString)] {
					numbers = append(numbers, number)
					break
				}
			}
		}
		combined := numbers[0] + numbers[len(numbers)-1]
		i, _ := strconv.Atoi(combined)
		result = result + i
	}

	fmt.Println(result)
}
