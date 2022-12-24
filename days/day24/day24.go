// days/day24/day24.go

package day24

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
	"time"
)

type pos struct {
	x, y int
}

// Where:                Down      UP    Right    Left     Stay
var directions = [5]pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {0, 0}}

const (
	U int = 1
	D int = 0
	L int = 3
	R int = 2
	S int = 4
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day24/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := [][][]byte{}
	for i := range inputS {
		aux := [][]byte{}
		for j := range inputS[i] {
			val := inputS[i][j]
			add := []byte{}
			if val != '.' {
				add = append(add, val)
			}
			aux = append(aux, add)
		}
		input = append(input, aux)
	}

	//printThis(input)
	sol1, sol2 := part(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part(input [][][]byte) (int, int) {
	base := pos{x: 0, y: 1}
	finish := pos{x: len(input) - 1, y: len(input[0]) - 2}

	sol1, input := travel(input, base, finish)

	//printThis(input)

	part2, input := travel(input, finish, base)

	part3, _ := travel(input, base, finish)

	return sol1, sol1 + part2 + part3
}

func travel(input [][][]byte, initPos, target pos) (int, [][][]byte) {
	positions := []pos{initPos}
	time := 1
	for ; ; time++ {
		snow := move(input)
		newPositions := []pos{}
		for _, p := range positions {
			for _, dPos := range directions {
				newPos := pos{x: p.x + dPos.x, y: p.y + dPos.y}
				// Si es la salida terminamos
				if newPos == target {
					return time, snow
				}

				// Si se sale del mapa no lo agregas
				if newPos.x < 0 || newPos.x >= len(input) || newPos.y < 0 || newPos.y >= (len(input[0])) {
					continue
				}
				// Si se choca contra pared o nieve no lo agregas
				if len(snow[newPos.x][newPos.y]) != 0 {
					continue
				}

				// Ahora si lo agregas
				newPositions = append(newPositions, newPos)
			}
		}

		newPositions = utils.UniqueConv(newPositions)
		// Si tenemos más de 75 posiciones nuevas nos quedamos solo con las 50 mejores
		if len(newPositions) >= 75 {
			sort.Slice(newPositions, func(i, j int) bool {
				return manhattanDistance(newPositions[i], target) < manhattanDistance(newPositions[j], target)
			})

			newPositions = newPositions[:75]
		}

		positions = newPositions
		input = snow
		//fmt.Println(positions[0])
		//printThis(snow)
	}
}

func manhattanDistance(p1, p2 pos) int {
	return int(math.Abs(float64(p1.x-p2.x))) + int(math.Abs(float64(p1.y-p2.y)))
}

func move(snow [][][]byte) [][][]byte {
	sol := [][][]byte{}
	for _, aux := range snow {
		vals := [][]byte{}
		for range aux {
			vals = append(vals, []byte{})
		}
		sol = append(sol, vals)
	}

	for i, aux := range snow {
		for j, vals := range aux {
			for _, val := range vals {
				switch val {
				case '#':
					sol[i][j] = append(sol[i][j], '#')
				case '>':
					if j < len(snow[i])-2 {
						sol[i][j+1] = append(sol[i][j+1], '>')
					} else {
						sol[i][1] = append(sol[i][1], '>')
					}
				case '<':
					if j > 1 {
						sol[i][j-1] = append(sol[i][j-1], '<')
					} else {
						sol[i][len(snow[i])-2] = append(sol[i][len(snow[i])-2], '<')
					}
				case '^':
					if i > 1 {
						sol[i-1][j] = append(sol[i-1][j], '^')
					} else {
						sol[len(snow)-2][j] = append(sol[len(snow)-2][j], '^')
					}
				case 'v':
					if i < len(snow)-2 {
						sol[i+1][j] = append(sol[i+1][j], 'v')
					} else {
						sol[1][j] = append(sol[1][j], 'v')
					}
				}
			}
		}
	}

	return sol
}

/*
func printThis(input [][][]byte) {
	str := ""
	for _, vals := range input {
		for _, val := range vals {
			if len(val) == 0 {
				str += "."
			} else if len(val) == 1 {
				str += string(val[0])
			} else {
				str += fmt.Sprint(len(val))
			}
		}
		str += "\n"
	}

	fmt.Println(str)
}
*/
