// days/day17/day17.go

package day17

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
)

type pos struct {
	x, y int
}

var pieces [5][]pos = [5][]pos{
	{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
	{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
}

func Solve() {

	// start := time.Now()

	// Input reading
	input, err := ioutil.ReadFile("./days/day17/input.txt")

	if err != nil {
		input, _ = ioutil.ReadFile("input.txt")
	}

	sol1 := part1(input)
	fmt.Println(sol1)

}

func part1(input []byte) int64 {
	var grid [7][100000]byte
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	top := -1
	cur := 0
	i := 0
	fmt.Println(len(input))
	//dif := 0
	//prevTop := 0
	getLoop := 240 + (1_000_000_000_000-240)%1740
	fmt.Println(getLoop)
	var check [20000][5]bool
	for ; cur < 1180; cur++ {
		// i = 1447 starts loop with cur%5 = 0
		// height dif is 2681, it loops every 1740 and starts at 240
		// You need to loop 240 + (1_000_000_000_000 - 240) % 1740
		/*
			if check[i][cur%5] {
				fmt.Println(i, cur%5, top)
			}
		*/
		if i == 1447 && cur%5 == 0 {
			//fmt.Println(top)
			//dif = top - prevTop
			//prevTop = top
			//fmt.Println(dif)
			fmt.Println(cur)
		}
		check[i][cur%5] = true
		piece := pieces[cur%5]
		p := pos{2, top + 4}
		for canFall := true; canFall; i = (i + 1) % len(input) {
			//fmt.Println("Rock: ", cur, p.x, p.y)
			if canMove(grid, input[i], p, piece) {
				switch input[i] {
				case '<':
					//fmt.Println("LEFT")
					p = pos{p.x - 1, p.y}
				case '>':
					//fmt.Println("RIGHT")
					p = pos{p.x + 1, p.y}
				}
			}

			if canMove(grid, 'D', p, piece) {
				//fmt.Println("DOWN")
				p = pos{p.x, p.y - 1}
			} else {
				canFall = false
			}
		}

		for _, pix := range piece {
			grid[p.x+pix.x][p.y+pix.y] = '@'
		}

		// Update new top
		switch cur % 5 {
		case 0:
			top = utils.Max(top, p.y)
		case 1:
			top = utils.Max(top, p.y+2)
		case 2:
			top = utils.Max(top, p.y+2)
		case 3:
			top = utils.Max(top, p.y+3)
		case 4:
			top = utils.Max(top, p.y+1)
		}

		/*
			gridNew := [][]byte{}
			for _, aux := range grid[0:7] {
				gridNew = append(gridNew, aux[0:10000])
			}
			fmt.Println("PASE")
			grid2 := utils.Transpose(gridNew)
			fmt.Println("PASE")
		*/
	}

	for _, val := range grid[0:7] {
		fmt.Println(string(val[0:15]))
	}

	finalTop := int64(top) + (1_000_000_000_000/1740)*2681

	return finalTop + 1
}

func canMove(grid [7][100000]byte, dir byte, p pos, piece []pos) bool {
	switch dir {
	case '<':
		p = pos{p.x - 1, p.y}
		for _, pix := range piece {
			if p.x+pix.x < 0 || grid[p.x+pix.x][p.y+pix.y] == '@' {
				return false
			}
		}

	case '>':
		p = pos{p.x + 1, p.y}
		for _, pix := range piece {
			if p.x+pix.x > 6 || grid[p.x+pix.x][p.y+pix.y] == '@' {
				return false
			}
		}

	case 'D':
		p = pos{p.x, p.y - 1}
		for _, pix := range piece {
			if p.y < 0 || grid[p.x+pix.x][p.y+pix.y] == '@' {
				return false
			}
		}

	}

	return true
}
