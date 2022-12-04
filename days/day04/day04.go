// days/day04/day04.go

package day04

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type pair struct {
	left  int
	right int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day04/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol1, sol2 := part1(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(input []string) (sol1, sol2 int) {
	for _, str := range input {
		var pair1, pair2 pair
		fmt.Sscanf(str, "%d-%d,%d-%d", &pair1.left, &pair1.right, &pair2.left, &pair2.right)
		if pair1.left <= pair2.left && pair1.right >= pair2.right ||
			pair2.left <= pair1.left && pair2.right >= pair1.right {
			sol1++
		}
		if pair1.right >= pair2.left && pair2.right >= pair1.left {
			sol2++
		}
	}
	return
}
