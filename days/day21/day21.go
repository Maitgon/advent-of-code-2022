// days/day21/day21.go

package day21

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type monkey struct {
	n  int
	op operation
	m1 string
	m2 string
}

type operation = int

const (
	NONE operation = iota
	SUM
	MULT
	SUB
	DIV
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day21/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := parseInput(inputS)
	/*
		input2 := map[string]monkey{}
		for k, v := range input {
			input2[k] = v
		}
	*/
	//fmt.Println(input)

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func part1(input map[string]monkey) int {
	return getValue("root", input)
}

func getValue(target string, input map[string]monkey) int {
	if input[target].op == NONE || input[target].n != -1 {
		return input[target].n
	}
	val1 := getValue(input[target].m1, input)
	val2 := getValue(input[target].m2, input)
	var sol int
	switch input[target].op {
	case SUM:
		sol = val1 + val2
	case SUB:
		sol = val1 - val2
	case MULT:
		sol = val1 * val2
	case DIV:
		sol = val1 / val2
	}

	return sol
}

func part2(input map[string]monkey) int {
	initMonkey := input["root"]
	if hasHuman(initMonkey.m1, input) {
		val := getValue(initMonkey.m2, input)
		return humanValue(val, initMonkey.m1, input)
	} else {
		val := getValue(initMonkey.m1, input)
		return humanValue(val, initMonkey.m2, input)
	}
}

func humanValue(sol int, target string, input map[string]monkey) int {
	if target == "humn" {
		return sol
	}
	if hasHuman(input[target].m1, input) {
		val2 := getValue(input[target].m2, input)

		switch input[target].op {
		case SUM: // sol = val1 + val2 -> val1 = sol - val2
			return humanValue(sol-val2, input[target].m1, input)
		case SUB: // sol = val1 - val2 -> val1 = sol + val2
			return humanValue(sol+val2, input[target].m1, input)
		case MULT: // sol = val1 * val2 -> val1 = sol/val2
			return humanValue(sol/val2, input[target].m1, input)
		case DIV: // sol = val1 / val2 -> val1 = sol * val2
			return humanValue(sol*val2, input[target].m1, input)
		}
	} else {
		val1 := getValue(input[target].m1, input)
		switch input[target].op {
		case SUM: // sol = val1 + val2 -> val2 = sol - val1
			return humanValue(sol-val1, input[target].m2, input)
		case SUB: // sol = val1 - val2 -> val2 = val1 - sol
			return humanValue(val1-sol, input[target].m2, input)
		case MULT: // sol = val1 * val2 -> val2 = sol / val1
			return humanValue(sol/val1, input[target].m2, input)
		case DIV: // sol = val1 / val2 -> val2 = val1 / sol
			return humanValue(val1/sol, input[target].m2, input)
		}
	}

	return 0
}

func parseInput(inputS []string) map[string]monkey {
	re := regexp.MustCompile(`(\w{4}): (\d+|(\w{4}) ([\+\*\-/]) (\w{4}))`)
	sol := map[string]monkey{}
	for _, val := range inputS {
		var m monkey
		res := re.FindStringSubmatch(val)
		n, err := strconv.Atoi(res[2])
		if err != nil {
			op := res[4]
			switch op {
			case "+":
				m.op = SUM
			case "-":
				m.op = SUB
			case "*":
				m.op = MULT
			case "/":
				m.op = DIV
			}
			m.m1 = res[3]
			m.m2 = res[5]
			m.n = -1
		} else {
			m.n = n
			m.op = NONE
		}
		sol[res[1]] = m

	}

	return sol

}

func hasHuman(target string, input map[string]monkey) bool {
	if input[target].op == NONE {
		return target == "humn"
	}
	return hasHuman(input[target].m1, input) || hasHuman(input[target].m2, input)
}
