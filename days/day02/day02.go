// days/day02/day02.go

package day02

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day02/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol1 := combatAll(input)
	sol2 := combatAll2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func combatAll(input []string) int {
	total := 0
	for _, clash := range input {
		total += combat(clash)
	}
	return total
}

func combat(clash string) int {
	result := 0
	if clash[2] == 'X' {
		result += 1
	} else if clash[2] == 'Y' {
		result += 2
	} else if clash[2] == 'Z' {
		result += 3
	}

	if clash[0]+23 == clash[2] {
		result += 3
	} else if clash[0] == 'A' && clash[2] == 'Y' ||
		clash[0] == 'B' && clash[2] == 'Z' ||
		clash[0] == 'C' && clash[2] == 'X' {
		result += 6
	}

	return result
}

func combatAll2(input []string) int {
	total := 0
	for _, clash := range input {
		total += combat2(clash)
	}
	return total
}

func combat2(clash string) int {
	if clash[2] == 'X' {
		if clash[0] == 'A' {
			return 3
		} else if clash[0] == 'B' {
			return 1
		} else {
			return 2
		}
	} else if clash[2] == 'Y' {
		if clash[0] == 'A' {
			return 4
		} else if clash[0] == 'B' {
			return 5
		} else {
			return 6
		}
	} else {
		if clash[0] == 'A' {
			return 8
		} else if clash[0] == 'B' {
			return 9
		} else {
			return 7
		}
	}
}
