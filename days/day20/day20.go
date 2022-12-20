// days/day20/day20.go

package day20

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day20/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := make([]int, len(inputS))
	for i, val := range inputS {
		num, _ := strconv.Atoi(val)
		input[i] = num
	}

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(input []int) int {
	positions := make([]int, len(input))

	for i := range positions {
		positions[i] = i
	}

	// We calculate where are the final positions of each value
	for i, val := range input {
		prio := positions[i]
		newPos := (prio + (val % (len(input) - 1))) % len(input)
		if newPos == prio {
			continue
		}
		if newPos < 0 {
			newPos += len(input)
		}

		if val > 0 {
			if prio < newPos {
				for j := range positions {
					if prio < positions[j] && positions[j] <= newPos {
						positions[j]--
					}
				}
			} else {
				for j := range positions {
					if (prio < positions[j] || positions[j] <= newPos) && positions[j] != 0 {
						positions[j]--
					} else if positions[j] == 0 {
						positions[j] = len(input) - 1
					}
				}
			}
		} else if val < 0 {
			if newPos < prio {
				for j := range positions {
					if newPos <= positions[j] && positions[j] < prio {
						positions[j]++
					}
				}
			} else {
				for j := range positions {
					if (newPos <= positions[j] || positions[j] < prio) && positions[j] != len(input)-1 {
						positions[j]++
					} else if positions[j] == len(input)-1 {
						positions[j] = 0
					}
				}
			}
		}
		positions[i] = newPos

		/*
			newPositions := make([]int, len(input))
			for i := range newPositions {
				newPositions[positions[i]] = input[i]
			}
			fmt.Println(newPositions)
		*/

	}

	// We create the new array with the given positions
	var zeroPos int
	newPositions := make([]int, len(input))
	for i := range newPositions {
		newPositions[positions[i]] = input[i]
		if input[i] == 0 {
			zeroPos = positions[i]
		}
	}

	// We get the new values:
	return newPositions[(zeroPos+1000)%len(input)] +
		newPositions[(zeroPos+2000)%len(input)] +
		newPositions[(zeroPos+3000)%len(input)]
}

func part2(nInput []int) int64 {
	input := make([]int64, len(nInput))
	for i, val := range nInput {
		input[i] = int64(val) * 811589153
	}

	positions := make([]int64, len(input))

	for i := range positions {
		positions[i] = int64(i)
	}

	// We calculate where are the final positions of each value
	for k := 0; k < 10; k++ {
		for i, val := range input {
			prio := positions[i]
			newPos := (prio + (val % int64((len(input) - 1)))) % int64(len(input))
			if newPos == prio {
				continue
			}
			if newPos < 0 {
				newPos += int64(len(input))
			}

			if val > 0 {
				if prio < newPos {
					for j := range positions {
						if prio < positions[j] && positions[j] <= newPos {
							positions[j]--
						}
					}
				} else {
					for j := range positions {
						if (prio < positions[j] || positions[j] <= newPos) && positions[j] != 0 {
							positions[j]--
						} else if positions[j] == 0 {
							positions[j] = int64(len(input) - 1)
						}
					}
				}
			} else if val < 0 {
				if newPos < prio {
					for j := range positions {
						if newPos <= positions[j] && positions[j] < prio {
							positions[j]++
						}
					}
				} else {
					for j := range positions {
						if (newPos <= positions[j] || positions[j] < prio) && positions[j] != int64(len(input)-1) {
							positions[j]++
						} else if positions[j] == int64(len(input)-1) {
							positions[j] = 0
						}
					}
				}
			}
			positions[i] = newPos

			/*
				newPositions := make([]int, len(input))
				for i := range newPositions {
					newPositions[positions[i]] = input[i]
				}
				fmt.Println(newPositions)
			*/
		}
	}

	// We create the new array with the given positions
	var zeroPos int64
	newPositions := make([]int64, len(input))
	for i := range newPositions {
		newPositions[positions[i]] = input[i]
		if input[i] == 0 {
			zeroPos = positions[i]
		}
	}

	// We get the new values:
	return newPositions[(zeroPos+1000)%int64(len(input))] +
		newPositions[(zeroPos+2000)%int64(len(input))] +
		newPositions[(zeroPos+3000)%int64(len(input))]
}
