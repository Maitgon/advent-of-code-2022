// days/day08/day08.go

package day08

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day08/input.txt")

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

func part1(input []string) int {
	sol := 0
	for i, row := range input {
		for j := range row {
			if notSeen(&input, i, j) {
				sol++
			}
		}
	}
	return sol
}

func part2(input []string) int {
	sol := 0
	for i, row := range input {
		for j := range row {
			if score := getScenicScore(&input, i, j); score >= sol {
				sol = score
			}
		}
	}
	return sol
}

func notSeen(input *[]string, i, j int) bool {
	if i == 0 || j == 0 || i == len(*input)-1 || j == len((*input)[0])-1 {
		return true
	}

	value := (*input)[i][j]

	// Look for trees in the top
	top := byte(0)
	for ix := 0; ix < i; ix++ {
		if (*input)[ix][j] > top {
			top = (*input)[ix][j]
		}
	}
	if top < value {
		return true
	}

	// Look for trees in the bottom
	bot := byte(0)
	for iy := len(*input) - 1; iy > i; iy-- {
		if (*input)[iy][j] > bot {
			bot = (*input)[iy][j]
		}
	}
	if bot < value {
		return true
	}

	// Look for trees in the left
	left := byte(0)
	for jx := 0; jx < j; jx++ {
		if (*input)[i][jx] > left {
			left = (*input)[i][jx]
		}
	}
	if left < value {
		return true
	}

	// Look for values on the right
	right := byte(0)
	for jy := len((*input)[0]) - 1; jy > j; jy-- {
		if (*input)[i][jy] > right {
			right = (*input)[i][jy]
		}
	}
	return right < value
}

func getScenicScore(input *[]string, i, j int) int {
	value := (*input)[i][j]

	// Top scenic
	ix := i
	var top int
	for ; ix > 0; ix-- {
		if (*input)[ix][j] >= value && ix != i {
			break
		}
	}
	top = i - ix

	// Bot scenic
	iy := i
	var bot int
	for ; iy < len(*input)-1; iy++ {
		if (*input)[iy][j] >= value && iy != i {
			break
		}
	}
	bot = iy - i

	// Left scenic
	jx := j
	var left int
	for ; jx > 0; jx-- {
		if (*input)[i][jx] >= value && jx != j {
			break
		}
	}
	left = j - jx

	// Right scenic
	jy := j
	var right int
	for ; jy < len((*input)[0])-1; jy++ {
		if (*input)[i][jy] >= value && jy != j {
			break
		}
	}
	right = jy - j

	return top * bot * left * right
}
