// days/day14/day14.go

package day14

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type position struct {
	x int
	y int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day14/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputAux := strings.Split(string(bs), "\n")
	input := [][]position{}
	for _, lines := range inputAux {
		inputAux := []position{}
		for _, pos := range strings.Split(lines, " -> ") {
			aux := strings.Split(pos, ",")
			posX, _ := strconv.Atoi(aux[0])
			posY, _ := strconv.Atoi(aux[1])
			inputAux = append(inputAux, position{x: posX, y: posY})
		}
		input = append(input, inputAux)
	}

	//fmt.Println(input)

	grid1 := gridCreation1(input)
	sol1 := fillGrid(grid1)

	grid2 := gridCreation2(input)
	sol2 := fillGrid(grid2)

	/*
		grid2 = utils.Transpose(grid2)
		for _, val := range grid2[0:20] {
			fmt.Println(string(val[490:510]))
		}
	*/

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func gridCreation1(input [][]position) [][]byte {
	grid := make([][]byte, 1000)
	for i := range grid {
		gridAux := make([]byte, 1000)
		for j := range gridAux {
			gridAux[j] = '.'
		}
		grid[i] = gridAux
	}
	grid[500][0] = '+'

	for _, p := range input {
		for i := 0; i < len(p)-1; i++ {
			fst := p[i]
			snd := p[i+1]

			// If the x position is different, then the y position is equal and viceversa
			if fst.x < snd.x {
				for dx := fst.x; dx <= snd.x; dx++ {
					grid[dx][fst.y] = '#'
				}
			} else if fst.x > snd.x {
				for dx := fst.x; dx >= snd.x; dx-- {
					grid[dx][fst.y] = '#'
				}
			} else if fst.y < snd.y {
				for dy := fst.y; dy <= snd.y; dy++ {
					grid[fst.x][dy] = '#'
				}
			} else if fst.y > snd.y {
				for dy := fst.y; dy >= snd.y; dy-- {
					grid[fst.x][dy] = '#'
				}
			}
		}
	}

	return grid
}

func gridCreation2(input [][]position) [][]byte {
	grid := make([][]byte, 1000)
	for i := range grid {
		gridAux := make([]byte, 1000)
		for j := range gridAux {
			gridAux[j] = '.'
		}
		grid[i] = gridAux
	}
	grid[500][0] = '+'

	for _, p := range input {
		for i := 0; i < len(p)-1; i++ {
			fst := p[i]
			snd := p[i+1]

			// If the x position is different, then the y position is equal and viceversa
			if fst.x < snd.x {
				for dx := fst.x; dx <= snd.x; dx++ {
					grid[dx][fst.y] = '#'
				}
			} else if fst.x > snd.x {
				for dx := fst.x; dx >= snd.x; dx-- {
					grid[dx][fst.y] = '#'
				}
			} else if fst.y < snd.y {
				for dy := fst.y; dy <= snd.y; dy++ {
					grid[fst.x][dy] = '#'
				}
			} else if fst.y > snd.y {
				for dy := fst.y; dy >= snd.y; dy-- {
					grid[fst.x][dy] = '#'
				}
			}
		}
	}

	maxY := 0
	for _, pAux := range input {
		for _, p := range pAux {
			if p.y > maxY {
				maxY = p.y
			}
		}
	}
	for i := 0; i < 1000; i++ {
		grid[i][maxY+2] = '#'
	}

	return grid
}

func fillGrid(grid [][]byte) int {
	numSand := 0
	for ; ; numSand++ {
		sand := position{x: 500, y: 0}
		canMove := true
		for i := 0; i <= 200 && canMove; i++ {
			if grid[sand.x][sand.y+1] == '.' {
				sand = position{x: sand.x, y: sand.y + 1}
			} else if grid[sand.x-1][sand.y+1] == '.' {
				sand = position{x: sand.x - 1, y: sand.y + 1}
			} else if grid[sand.x+1][sand.y+1] == '.' {
				sand = position{x: sand.x + 1, y: sand.y + 1}
			} else {
				grid[sand.x][sand.y] = 'o'
				canMove = false
			}
		}
		if canMove {
			break
		}
		if sand == (position{x: 500, y: 0}) {
			numSand++
			break
		}
	}
	return numSand
}
