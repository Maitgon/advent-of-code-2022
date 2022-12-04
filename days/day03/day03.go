// days/day03/day03.go

package day03

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day03/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(input []string) (total int) {
	total = 0
	for _, rucksack := range input {
		mid := len(rucksack) / 2
		letter := utils.Intersection([]byte(rucksack[0:mid]), ([]byte(rucksack[mid:])))[0]
		total += getValue(letter)
	}
	return
}

func part2(input []string) (total int) {
	total = 0
	for i := 0; i < len(input); i += 3 {
		letter := utils.Intersection3([]byte(input[i]), []byte(input[i+1]), []byte(input[i+2]))[0]
		total += getValue(letter)
	}
	return
}

func getValue(letter byte) int {
	if letter >= 'a' {
		return int(letter - 'a' + 1)
	} else {
		return int(letter - 'A' + 27)
	}
}
