// days/day23/day23.go

package day23

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type dir = uint16

type pos struct {
	x, y int
}

const (
	N dir = 1
	S dir = 5
	W dir = 7
	E dir = 3
)

var directions [8]pos = [8]pos{{-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}}

//               012
// directions:   7#3
//               654

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day23/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := make(map[pos]bool)
	for i, val := range inputS {
		for j, aux := range val {
			if aux == '#' {
				//input[pos{x: j, y: i}] = [4]dir{N, S, W, E}
				input[pos{x: j, y: i}] = true
			}
		}
	}

	//fmt.Println(input)

	sol1, sol2 := part1(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(input map[pos]bool) (int, int) {
	var sol1, sol2 int
	dirs := [4]dir{N, S, W, E}
	for i := 0; ; i++ {
		newArray := make(map[pos]bool)   // El nuevo array que será input
		newPos := make(map[pos]pos)      // Dice a donde van ahora
		newPosCount := make(map[pos]int) // Contea cuantas veces aparecen nuevas posiciones
		for key := range input {

			//fmt.Print("Se mueve ")
			// Si no puede moverse no se mueve :D
			if !canMove(key, input) {
				//fmt.Println("Se esta aislado")
				newArray[key] = true
				//newPos[key] = key
				//newPosCount[key] += 1
				continue
			}

			// Proponer movimientos
			moved := false
			for _, d := range dirs {
				if canMoveDir(key, input, d) {
					//fmt.Println("Se mueve")
					dPos := directions[d]
					newPos[key] = pos{x: key.x + dPos.x, y: key.y + dPos.y}
					newPosCount[newPos[key]] += 1
					moved = true
					break
				}
			}

			if !moved {
				//fmt.Println("Se queda quieto")
				newArray[key] = true
				continue
				//newPos[key] = key
				//newPosCount[key] += 1
			}
		}
		// Vemos si pueden moverse
		for old, new := range newPos {
			if newPosCount[new] == 1 {
				newArray[new] = true
			} else {
				newArray[old] = true
			}
		}

		if i == 9 {
			// Ahora buscamos los valores mas grandes y pequeños para hallar el área
			var maxX, minX, maxY, minY int
			for key := range input {
				maxX, minX, maxY, minY = key.x, key.x, key.y, key.y
				break
			}

			for key := range input {
				if key.x > maxX {
					maxX = key.x
				}
				if key.x < minX {
					minX = key.x
				}
				if key.y > maxY {
					maxY = key.y
				}
				if key.y < minY {
					minY = key.y
				}
			}
			sol1 = (maxX-minX+1)*(maxY-minY+1) - len(input)
		}

		equal := true
		for key := range input {
			if input[key] != newArray[key] {
				equal = false
			}
		}

		if equal {
			sol2 = i + 1
			break
		}

		// Actualizamos las direcciones
		dirs[0], dirs[1], dirs[2], dirs[3] = dirs[1], dirs[2], dirs[3], dirs[0]

		// Actualizamos el array input
		input = make(map[pos]bool)
		for key := range newArray {
			input[key] = true
		}
	}

	return sol1, sol2
}

func canMove(p pos, input map[pos]bool) bool {
	for d := range directions {
		dPos := directions[d]
		newPos := pos{x: p.x + dPos.x, y: p.y + dPos.y}
		_, ok := input[newPos]
		if ok {
			return true
		}
	}

	return false
}

func canMoveDir(p pos, input map[pos]bool, d dir) bool {
	//fmt.Println(p, d)
	for i := -1; i <= 1; i++ {
		dPos := directions[(int(d)+i)%len(directions)]
		newPos := pos{x: p.x + dPos.x, y: p.y + dPos.y}
		//fmt.Println(newPos)

		if input[newPos] {
			return false
		}
	}
	return true
}
