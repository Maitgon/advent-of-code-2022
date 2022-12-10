// days/day10/day10.go

package day10

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day10/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Printf("The solution to part 2 is:\n%v\n", sol2)
	fmt.Println("Time: ", end)
}

func part1(input []string) int {
	x := 1
	cycle := 1
	sol := 0
	for _, vals := range input {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			sol += cycle * x
		}
		spl := strings.Split(vals, " ")
		if len(spl) == 1 {
			cycle++
		} else {
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				sol += cycle * x
			}
			cycle++
			val, _ := strconv.Atoi(spl[1])
			x += val
		}
	}
	return sol
}

func part2(input []string) string {
	x := 1
	cycle := 0
	chart := [6][40]rune{}
	for _, vals := range input {
		row, column := getPos(cycle)
		if math.Abs(float64(x-cycle%40)) <= 1 {
			chart[row][column] = '#'
		} else {
			chart[row][column] = '.'
		}
		spl := strings.Split(vals, " ")
		if len(spl) == 1 {
			cycle++
		} else {
			cycle++
			row, column := getPos(cycle)
			if math.Abs(float64(x-cycle%40)) <= 1 {
				chart[row][column] = '#'
			} else {
				chart[row][column] = '.'
			}
			cycle++
			val, _ := strconv.Atoi(spl[1])
			x += val
		}
	}
	sol := ""
	for _, aux := range chart {
		sol += string(aux[:]) + "\n"
	}
	return sol
}

func getPos(cycle int) (int, int) {
	return cycle / 40, cycle % 40
}
