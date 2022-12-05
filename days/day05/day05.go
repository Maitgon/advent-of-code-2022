// days/day05/day05.go

package day05

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	n    int
	from int
	to   int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day05/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n\n")

	stacks := parseStacks(strings.Split(input[0], "\n"))
	stacks2 := make([]utils.Stack[byte], len(stacks))
	for i := 0; i < len(stacks); i++ {
		stackAux := make(utils.Stack[byte], len(stacks[i]))
		copy(stackAux, stacks[i])
		stacks2[i] = stackAux
	}
	instructs := parseInstructions(strings.Split(input[1], "\n"))
	sol1 := part1(stacks, instructs)
	sol2 := part2(stacks2, instructs)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func parseInstructions(inputs []string) []instruction {
	regex := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	instruct := make([]instruction, 0)
	for _, inp := range inputs {
		res := regex.FindStringSubmatch(inp)
		n, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])
		ins := instruction{n: n, from: from, to: to}
		instruct = append(instruct, ins)
	}
	return instruct
}

func parseStacks(inputs []string) []utils.Stack[byte] {
	nStacks := (len(inputs[len(inputs)-1])+1)/4 + 1
	stacks := make([]utils.Stack[byte], nStacks)
	for i := len(inputs) - 2; i >= 0; i-- {
		for j := 0; j < nStacks-1; j++ {
			letter := inputs[i][4*j+1]
			if letter != ' ' {
				stacks[j+1].Push(letter)
			}
		}
	}
	return stacks
}

func part1(stacks []utils.Stack[byte], instructs []instruction) string {
	for _, instruct := range instructs {
		for i := 0; i < instruct.n; i++ {
			elem := stacks[instruct.from].Pop()
			stacks[instruct.to].Push(elem)
		}
	}

	sol := make([]byte, len(stacks)-1)
	for i := 0; i < len(sol); i++ {
		sol[i] = stacks[i+1].Last()
	}
	return string(sol[:])
}

func part2(stacks []utils.Stack[byte], instructs []instruction) string {
	for _, instruct := range instructs {
		elems := stacks[instruct.from].PopN(instruct.n)
		stacks[instruct.to].PushN(elems)
	}
	sol := make([]byte, len(stacks)-1)
	for i := 0; i < len(sol); i++ {
		sol[i] = stacks[i+1].Last()
	}
	return string(sol[:])
}
