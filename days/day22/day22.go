// days/day22/day22.go

package day22

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type pos struct {
	face, x, y, dir int
}

const (
	R int = 1000
	L int = 1001
	D int = 1002
	U int = 1003
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day22/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n\n")

	inputPart1 := strings.Split(inputS[0], "\n")
	inputPart2 := parseInstructions(inputS[1])

	//fmt.Println(inputPart2)

	//fmt.Println(len(inputPart1[0]))

	faces := parseFaces(inputPart1)
	/*
		for _, val := range faces[4] {
			fmt.Println(string(val[:]))
		}
	*/

	sol1 := part(inputPart2, faces, getDirection)
	sol2 := part(inputPart2, faces, getDirection2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part(instructions []int, faces [6][][]byte, getNewPos func(faces [6][][]byte, hPos pos, ins int) pos) int {
	hPos := pos{face: 0, x: 0, y: 0, dir: R}
	for _, ins := range instructions {

		// If the instruction is to change directions get new Direction
		if ins == L {
			switch hPos.dir {
			case D:
				hPos.dir = R
			case R:
				hPos.dir = U
			case U:
				hPos.dir = L
			case L:
				hPos.dir = D
			}
		} else if ins == R {
			switch hPos.dir {
			case D:
				hPos.dir = L
			case L:
				hPos.dir = U
			case U:
				hPos.dir = R
			case R:
				hPos.dir = D
			}
		} else {
			// If its not direction get the new direction you are getting
			hPos = getNewPos(faces, hPos, ins)
		}
	}

	// Faces:
	//         #12
	//         #3#
	//         45#
	//         6##

	sol := 0
	nRows := len(faces[hPos.face])
	nCols := len(faces[hPos.face][hPos.x])
	hPos.x++
	hPos.y++
	switch hPos.face {
	case 0:
		sol = 1000*hPos.x + 4*(nCols+hPos.y)
	case 1:
		sol = 1000*hPos.x + 4*(2*nCols+hPos.y)
	case 2:
		sol = 1000*(nRows+hPos.x) + 4*(nCols+hPos.y)
	case 3:
		sol = 1000*(2*nRows+hPos.x) + 4*hPos.y
	case 4:
		sol = 1000*(2*nRows+hPos.x) + 4*(nCols+hPos.y)
	case 5:
		sol = 1000*(3*nRows+hPos.x) + 4*hPos.y
	}

	switch hPos.dir {
	case D:
		sol += 1
	case L:
		sol += 2
	case U:
		sol += 3
	}

	return sol
}

func getDirection(faces [6][][]byte, hPos pos, ins int) pos {
	switch hPos.dir {
	case R:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x, y: hPos.y + 1, dir: hPos.dir}
			if newPos.y >= len(faces[hPos.face][hPos.x]) {
				switch newPos.face {
				case 0: // De cara 1 vas a cara 2
					newPos.face = 1
					newPos.y = 0
				case 1: // De cara 2 vas a cara 1
					newPos.face = 0
					newPos.y = 0
				case 2: // De cara 3 vas a cara 3
					newPos.y = 0
				case 3: // De cara 4 vas a cara 5
					newPos.face = 4
					newPos.y = 0
				case 4: // De cara 5 vas a cara 4
					newPos.face = 3
					newPos.y = 0
				case 5: // De cara 6 vas a cara 6
					newPos.y = 0
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
		}

	case L:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x, y: hPos.y - 1, dir: hPos.dir}
			if newPos.y < 0 {
				switch newPos.face {
				case 0: // De cara 1 vas a cara 2
					newPos.face = 1
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 1: // De cara 2 vas a cara 1
					newPos.face = 0
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 2: // De cara 3 vas a cara 3
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 3: // De cara 4 vas a cara 5
					newPos.face = 4
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 4: // De cara 5 vas a cara 4
					newPos.face = 3
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 5: // De cara 6 vas a cara 6
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
		}

	case U:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x - 1, y: hPos.y, dir: hPos.dir}
			if newPos.x < 0 {
				switch newPos.face {
				case 0: // De cara 1 vas a cara 5
					newPos.face = 4
					newPos.x = len(faces[newPos.face]) - 1
				case 1: // De cara 2 vas a cara 2
					newPos.x = len(faces[newPos.face]) - 1
				case 2: // De cara 3 vas a cara 1
					newPos.face = 0
					newPos.x = len(faces[newPos.face]) - 1
				case 3: // De cara 4 vas a cara 6
					newPos.face = 5
					newPos.x = len(faces[newPos.face]) - 1
				case 4: // De cara 5 vas a cara 3
					newPos.face = 2
					newPos.x = len(faces[newPos.face]) - 1
				case 5: // De cara 6 vas a cara 4
					newPos.face = 3
					newPos.x = len(faces[newPos.face]) - 1
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
		}

	case D:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x + 1, y: hPos.y, dir: hPos.dir}
			if newPos.x >= len(faces[newPos.face]) {
				switch newPos.face {
				case 0: // De cara 1 vas a cara 3
					newPos.face = 2
					newPos.x = 0
				case 1: // De cara 2 vas a cara 2
					newPos.x = 0
				case 2: // De cara 3 vas a cara 5
					newPos.face = 4
					newPos.x = 0
				case 3: // De cara 4 vas a cara 6
					newPos.face = 5
					newPos.x = 0
				case 4: // De cara 5 vas a cara 1
					newPos.face = 0
					newPos.x = 0
				case 5: // De cara 6 vas a cara 4
					newPos.face = 3
					newPos.x = 0
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
		}
	}

	return hPos
}

func getDirection2(faces [6][][]byte, hPos pos, ins int) pos {
	switch hPos.dir {
	case R:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x, y: hPos.y + 1, dir: hPos.dir}
			brk := false
			if newPos.y >= len(faces[hPos.face][hPos.x]) {
				switch newPos.face {
				case 0: // De cara 0R vas a cara 1L
					newPos.face = 1
					newPos.y = 0
				case 1: // De cara 1R vas a cara 4R
					newPos.face = 4
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
					newPos.x = len(faces[newPos.face]) - 1 - newPos.x
					newPos.dir = L
					brk = true
				case 2: // De cara 2R vas a cara 1D
					newPos.face = 1
					newPos.y = newPos.x
					newPos.x = len(faces[newPos.face]) - 1
					newPos.dir = U
					brk = true
				case 3: // De cara 3R vas a cara 4L
					newPos.face = 4
					newPos.y = 0
				case 4: // De cara 4R vas a cara 1R
					newPos.face = 1
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
					newPos.x = len(faces[newPos.face]) - 1 - newPos.x
					newPos.dir = L
					brk = true
				case 5: // De cara 5R vas a cara 4D
					newPos.y = 4
					newPos.y = newPos.x
					newPos.x = len(faces[newPos.face]) - 1
					newPos.dir = U
					brk = true
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
			if brk {
				hPos = getDirection2(faces, hPos, ins-i-1)
				break
			}
		}

	case L:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x, y: hPos.y - 1, dir: hPos.dir}
			brk := false
			if newPos.y < 0 {
				switch newPos.face {
				case 0: // De cara 0L vas a cara 3L
					newPos.face = 3
					newPos.y = 0
					newPos.x = len(faces[newPos.face]) - 1 - newPos.x
					newPos.dir = R
					brk = true
				case 1: // De cara 1L vas a cara 0R
					newPos.face = 0
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 2: // De cara 2L vas a cara 3U
					newPos.face = 3
					newPos.y = newPos.x
					newPos.x = 0
					newPos.dir = D
					brk = true
				case 3: // De cara 3L vas a cara 0L
					newPos.face = 0
					newPos.y = 0
					newPos.x = len(faces[newPos.face]) - 1 - newPos.x
					newPos.dir = R
					brk = true
				case 4: // De cara 4L vas a cara 3R
					newPos.face = 3
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
				case 5: // De cara 5L vas a cara 0U
					newPos.face = 0
					newPos.y = newPos.x
					newPos.x = 0
					newPos.dir = D
					brk = true
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
			if brk {
				hPos = getDirection2(faces, hPos, ins-i-1)
				break
			}
		}

	case U:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x - 1, y: hPos.y, dir: hPos.dir}
			brk := false
			if newPos.x < 0 {
				switch newPos.face {
				case 0: // De cara 0U vas a cara 5L
					newPos.face = 5
					newPos.x = newPos.y
					newPos.y = 0
					newPos.dir = R
					brk = true
				case 1: // De cara 1U vas a cara 5D
					newPos.face = 5
					newPos.x = len(faces[newPos.face]) - 1
				case 2: // De cara 2U vas a cara 0D
					newPos.face = 0
					newPos.x = len(faces[newPos.face]) - 1
				case 3: // De cara 3U vas a cara 2L
					newPos.face = 2
					newPos.x = newPos.y
					newPos.y = 0
					newPos.dir = R
					brk = true
				case 4: // De cara 4U vas a cara 2D
					newPos.face = 2
					newPos.x = len(faces[newPos.face]) - 1
				case 5: // De cara 5U vas a cara 3D
					newPos.face = 3
					newPos.x = len(faces[newPos.face]) - 1
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
			if brk {
				hPos = getDirection2(faces, hPos, ins-i-1)
				break
			}
		}

	case D:
		for i := 0; i < ins; i++ {
			newPos := pos{face: hPos.face, x: hPos.x + 1, y: hPos.y, dir: hPos.dir}
			brk := false
			if newPos.x >= len(faces[newPos.face]) {
				switch newPos.face {
				case 0: // De cara 0D vas a cara 2U
					newPos.face = 2
					newPos.x = 0
				case 1: // De cara 1D vas a cara 2R
					newPos.face = 2
					newPos.x = newPos.y
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
					newPos.dir = L
					brk = true
				case 2: // De cara 2D vas a cara 4U
					newPos.face = 4
					newPos.x = 0
				case 3: // De cara 3D vas a cara 5U
					newPos.face = 5
					newPos.x = 0
				case 4: // De cara 4D vas a cara 5R
					newPos.face = 5
					newPos.x = newPos.y
					newPos.y = len(faces[newPos.face][newPos.x]) - 1
					newPos.dir = L
					brk = true
				case 5: // De cara 5D vas a cara 1U
					newPos.face = 1
					newPos.x = 0
				}
			}
			if faces[newPos.face][newPos.x][newPos.y] == '#' {
				break
			}
			hPos = newPos
			if brk {
				hPos = getDirection2(faces, hPos, ins-i-1)
				break
			}
		}
	}

	return hPos
}

func parseInstructions(input string) []int {
	re := regexp.MustCompile(`\d+|(R|L)`)
	res := re.FindAllString(input, -1)

	parsed := make([]int, len(res))
	for i, val := range res {
		switch val {
		case "L":
			parsed[i] = L
		case "R":
			parsed[i] = R
		default:
			n, _ := strconv.Atoi(val)
			parsed[i] = n
		}
	}
	return parsed
}

func parseFaces(input []string) [6][][]byte {
	sideLength := len(input[0]) / 3
	// Faces:
	//         #12
	//         #3#
	//         45#
	//         6##

	sol := [6][][]byte{}

	// Rellenar cara 1
	cara1 := make([][]byte, sideLength)
	for i := 0; i < sideLength; i++ {
		line := make([]byte, sideLength)
		for j := sideLength; j < 2*sideLength; j++ {
			line[j-sideLength] = input[i][j]
		}
		cara1[i] = line
	}
	sol[0] = cara1

	// Rellenar cara 2
	cara2 := make([][]byte, sideLength)
	for i := 0; i < sideLength; i++ {
		line := make([]byte, sideLength)
		for j := 2 * sideLength; j < 3*sideLength; j++ {
			line[j-2*sideLength] = input[i][j]
		}
		cara2[i] = line
	}
	sol[1] = cara2

	// Rellenar cara 3
	cara3 := make([][]byte, sideLength)
	for i := sideLength; i < 2*sideLength; i++ {
		line := make([]byte, sideLength)
		for j := sideLength; j < 2*sideLength; j++ {
			line[j-sideLength] = input[i][j]
		}
		cara3[i-sideLength] = line
	}
	sol[2] = cara3

	// Rellenar cara 4
	cara4 := make([][]byte, sideLength)
	for i := 2 * sideLength; i < 3*sideLength; i++ {
		line := make([]byte, sideLength)
		for j := 0; j < sideLength; j++ {
			line[j] = input[i][j]
		}
		cara4[i-2*sideLength] = line
	}
	sol[3] = cara4

	// Rellenar cara 5
	cara5 := make([][]byte, sideLength)
	for i := 2 * sideLength; i < 3*sideLength; i++ {
		line := make([]byte, sideLength)
		for j := sideLength; j < 2*sideLength; j++ {
			line[j-sideLength] = input[i][j]
		}
		cara5[i-2*sideLength] = line
	}
	sol[4] = cara5

	// Rellenar cara 6
	cara6 := make([][]byte, sideLength)
	for i := 3 * sideLength; i < 4*sideLength; i++ {
		line := make([]byte, sideLength)
		for j := 0; j < sideLength; j++ {
			line[j] = input[i][j]
		}
		cara6[i-3*sideLength] = line
	}
	sol[5] = cara6

	return sol
}
