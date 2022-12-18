// days/day18/day18.go

package day18

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type position struct {
	x, y, z int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day18/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	var input []position
	for _, val := range inputS {
		var pos position
		fmt.Sscanf(val, "%d,%d,%d", &pos.x, &pos.y, &pos.z)
		input = append(input, pos)
	}

	/*
		var maxX, minX, maxY, minY, maxZ, minZ = input[0].x, input[0].x, input[0].y, input[0].y, input[0].z, input[0].z

		for _, p := range input {
			if maxX < p.x {
				maxX = p.x
			}
			if minX > p.x {
				minX = p.x
			}
			if maxY < p.y {
				maxY = p.y
			}
			if minY > p.y {
				minY = p.y
			}
			if maxZ < p.z {
				maxZ = p.z
			}
			if minZ < p.z {
				minZ = p.z
			}
		}

		fmt.Println(maxX, minX, maxY, minY, maxZ, minZ)
	*/

	//fmt.Println(input)

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(input []position) int {
	sol := 0
	for _, lava := range input {
		faces := 6
		for _, other := range input {
			if nextTo(lava, other) {
				faces--
			}
		}
		sol += faces
	}

	return sol
}

func part2(input []position) int {
	var newDrops []position
	var toSee []position
	pos := position{-1, -1, -1}
	seen := map[position]bool{}
	toSee = append(toSee, pos)

	for len(toSee) > 0 {
		p := toSee[0]
		toSee = toSee[1:]
		seen[p] = true
		newDrops = append(newDrops, p)

		for _, newP := range getAdjacents(p) {
			if !seen[newP] && !utils.Contains(toSee, newP) && !utils.Contains(input, newP) {
				toSee = append(toSee, newP)
			}
		}
	}

	faces := part1(newDrops)

	return faces - 6*24*24
}

func nextTo(p1, p2 position) bool {
	return (p1.x == p2.x && p1.y == p2.y && (p1.z == p2.z-1 || p1.z == p2.z+1)) ||
		(p1.x == p2.x && p1.z == p2.z && (p1.y == p2.y-1 || p1.y == p2.y+1)) ||
		(p1.z == p2.z && p1.y == p2.y && (p1.x == p2.x-1 || p1.x == p2.x+1))
}

func getAdjacents(p position) []position {
	sol := []position{}
	if p.x-1 >= -1 {
		sol = append(sol, position{p.x - 1, p.y, p.z})
	}
	if p.y-1 >= -1 {
		sol = append(sol, position{p.x, p.y - 1, p.z})
	}
	if p.z-1 >= -1 {
		sol = append(sol, position{p.x, p.y, p.z - 1})
	}
	if p.x+1 <= 22 {
		sol = append(sol, position{p.x + 1, p.y, p.z})
	}
	if p.y+1 <= 22 {
		sol = append(sol, position{p.x, p.y + 1, p.z})
	}
	if p.z+1 <= 22 {
		sol = append(sol, position{p.x, p.y, p.z + 1})
	}
	return sol
}
