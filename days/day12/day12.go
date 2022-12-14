// days/day12/day12.go

package day12

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type point struct {
	x int
	y int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day12/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")
	var startP point
	var endP point
	grid := utils.Make2D[int](len(input[0]), len(input))
	for i, aux := range input {
		for j, let := range aux {
			switch let {
			case 'S':
				grid[j][i] = 0
				startP = point{x: j, y: i}
			case 'E':
				grid[j][i] = 27
				endP = point{x: j, y: i}
			default:
				grid[j][i] = int(input[i][j]) - 96
			}
		}
	}

	//fmt.Println(startP)
	//fmt.Println(endP)

	sol1 := dijkstra(grid, startP, 27, neightbours1)
	sol2 := dijkstra(grid, endP, 1, neightbours2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func dijkstra(grid [][]int, startP point, heightEnd int, neightbours func(*[][]int, point) []point) int {
	queue := utils.Queue[point]{}
	dist := utils.Make2D[int](len(grid), len(grid[0]))
	for i, vals := range grid {
		for j := range vals {
			dist[i][j] = 999
		}
	}
	dist[startP.x][startP.y] = 0
	queue.Put(point{x: startP.x, y: startP.y})

	for len(queue) != 0 {
		//fmt.Println(queue)
		//for _, val := range dist {
		//	fmt.Println(val)
		//}
		u := queue.Get()

		for _, neig := range neightbours(&grid, u) {
			aux := dist[u.x][u.y] + 1
			if grid[neig.x][neig.y] == heightEnd {
				return aux
			}
			if aux < dist[neig.x][neig.y] {
				dist[neig.x][neig.y] = aux
				queue.Put(neig)
			}
		}

	}

	return 0

}

func neightbours1(grid *[][]int, p point) []point {
	sol := []point{}
	if p.x-1 >= 0 {
		if (*grid)[p.x][p.y]-(*grid)[p.x-1][p.y] >= -1 {
			sol = append(sol, point{x: p.x - 1, y: p.y})
		}
	}
	if p.y-1 >= 0 {
		if (*grid)[p.x][p.y]-(*grid)[p.x][p.y-1] >= -1 {
			sol = append(sol, point{x: p.x, y: p.y - 1})
		}
	}
	if p.x+1 < len(*grid) {
		if (*grid)[p.x][p.y]-(*grid)[p.x+1][p.y] >= -1 {
			sol = append(sol, point{x: p.x + 1, y: p.y})
		}
	}
	if p.y+1 < len((*grid)[0]) {
		if (*grid)[p.x][p.y]-(*grid)[p.x][p.y+1] >= -1 {
			sol = append(sol, point{x: p.x, y: p.y + 1})
		}
	}
	return sol
}

func neightbours2(grid *[][]int, p point) []point {
	sol := []point{}
	if p.x-1 >= 0 {
		if (*grid)[p.x][p.y]-(*grid)[p.x-1][p.y] <= 1 {
			sol = append(sol, point{x: p.x - 1, y: p.y})
		}
	}
	if p.y-1 >= 0 {
		if (*grid)[p.x][p.y]-(*grid)[p.x][p.y-1] <= 1 {
			sol = append(sol, point{x: p.x, y: p.y - 1})
		}
	}
	if p.x+1 < len(*grid) {
		if (*grid)[p.x][p.y]-(*grid)[p.x+1][p.y] <= 1 {
			sol = append(sol, point{x: p.x + 1, y: p.y})
		}
	}
	if p.y+1 < len((*grid)[0]) {
		if (*grid)[p.x][p.y]-(*grid)[p.x][p.y+1] <= 1 {
			sol = append(sol, point{x: p.x, y: p.y + 1})
		}
	}
	return sol
}
