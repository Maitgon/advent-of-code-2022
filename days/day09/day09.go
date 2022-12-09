// days/day09/day09.go

package day09

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type direction int

const (
	D direction = iota
	L
	U
	R
)

type move struct {
	dir   direction
	steps int
}

type position struct {
	x int
	y int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day09/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	// Parsing
	inputS := strings.Split(string(bs), "\n")
	input := make([]move, len(inputS))
	for i := range inputS {
		aux := strings.Split(inputS[i], " ")
		n, _ := strconv.Atoi(aux[1])
		switch aux[0] {
		case "D":
			input[i] = move{dir: D, steps: n}
		case "L":
			input[i] = move{dir: L, steps: n}
		case "U":
			input[i] = move{dir: U, steps: n}
		case "R":
			input[i] = move{dir: R, steps: n}
		}
	}

	sol1, sol2 := part(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part(input []move) (int, int) {
	table1 := map[position]bool{}
	table2 := map[position]bool{}
	H := [10]position{}
	for i := range H {
		H[i] = position{0, 0}
	}
	table1[H[1]] = true
	table2[H[9]] = true

	for _, mov := range input {
		for mov.steps > 0 {
			// We move H[0]
			switch mov.dir {
			case D:
				H[0].y--
			case U:
				H[0].y++
			case L:
				H[0].x--
			case R:
				H[0].x++
			}
			mov.steps--

			for i := 1; i < 10; i++ {
				if math.Abs(float64(H[i].y-H[i-1].y)) <= 1 && math.Abs(float64(H[i].x-H[i-1].x)) <= 1 {
					break
				}
				dx := H[i-1].x - H[i].x
				dy := H[i-1].y - H[i].y
				if dx > 0 {
					H[i].x++
				} else if dx < 0 {
					H[i].x--
				}
				if dy > 0 {
					H[i].y++
				} else if dy < 0 {
					H[i].y--
				}
			}
			table1[H[1]] = true
			table2[H[9]] = true
		}
	}
	return len(table1), len(table2)
}
