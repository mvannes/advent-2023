package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	f, err := os.Open("input")
	if nil != err {
		log.Fatal(err.Error())
	}
	scan := bufio.NewScanner(f)

	result := 0
	for scan.Scan() {
		line := scan.Text()
		r := regexp.MustCompile("([0-9])]")

		res := r.FindAllString(line, -1)

		combined := res[0] + res[len(res)-1]
		i, _ := strconv.Atoi(combined)
		result = result + i
	}

	fmt.Println(result)
}
