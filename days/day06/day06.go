// days/day06/day06.go

package day06

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day06/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := string(bs)

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(input string) int {
	pointer := 0
	curr := 0
	values := make([]byte, 4)
	for ; ; pointer++ {
		if !utils.Unique(values) || pointer < 4 {
			values[curr] = input[pointer]
		} else {
			break
		}
		curr = (curr + 1) % 4
	}
	return pointer
}

func part2(input string) int {
	pointer := 0
	curr := 0
	values := make([]byte, 14)
	for ; ; pointer++ {
		if !utils.Unique(values) || pointer < 14 {
			values[curr] = input[pointer]
		} else {
			break
		}
		curr = (curr + 1) % 14
	}
	return pointer
}
