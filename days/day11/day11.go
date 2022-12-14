// days/day11/day11.go

package day11

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

type operation int

const (
	SUM operation = iota
	MUL
	SQR
)

type monkey struct {
	id        int
	items     []int64
	op        operation
	opN       int64
	test      int
	T         int
	F         int
	inspected int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day11/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n\n")
	input1 := make([]monkey, len(inputS))
	input2 := make([]monkey, len(inputS))
	for i := range inputS {
		input1[i] = parseMonkey(inputS[i])
		input2[i] = parseMonkey(inputS[i])
	}

	sol1 := part1(input1)
	sol2 := part2(input2)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func inspectOnce(monkeys []monkey) []monkey {
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			var newWorryLvl int64
			switch monkey.op {
			case SUM:
				newWorryLvl = (item + monkey.opN) / 3
			case MUL:
				newWorryLvl = (item * monkey.opN) / 3
			case SQR:
				newWorryLvl = (item * item) / 3
			}
			if newWorryLvl%int64(monkey.test) == 0 {
				monkeys[monkey.T].items = append(monkeys[monkey.T].items, newWorryLvl)
			} else {
				monkeys[monkey.F].items = append(monkeys[monkey.F].items, newWorryLvl)
			}
			monkeys[i].inspected++
		}
		monkeys[i].items = []int64{}
	}

	return monkeys
}

func inspectOnce2(monkeys []monkey) []monkey {
	mcm := 1
	for _, monkey := range monkeys {
		mcm *= monkey.test
	}
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			var newWorryLvl int64
			switch monkey.op {
			case SUM:
				newWorryLvl = (item + monkey.opN) % int64(mcm)
			case MUL:
				newWorryLvl = (item * monkey.opN) % int64(mcm)
			case SQR:
				newWorryLvl = (item * item) % int64(mcm)
			}
			if newWorryLvl%int64(monkey.test) == 0 {
				monkeys[monkey.T].items = append(monkeys[monkey.T].items, newWorryLvl)
			} else {
				monkeys[monkey.F].items = append(monkeys[monkey.F].items, newWorryLvl)
			}
			monkeys[i].inspected++
		}
		monkeys[i].items = []int64{}
	}

	return monkeys
}

func part1(monkeys []monkey) int {
	for i := 0; i < 20; i++ {
		monkeys = inspectOnce(monkeys)
	}

	/*
		for i := 0; i < len(monkeys); i++ {
			fmt.Println(monkeys[i].items)
		}
	*/

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspected > monkeys[j].inspected })

	//fmt.Printf("monkey %d: %d, monkey %d: %d\n", monkeys[0].id, monkeys[0].inspected, monkeys[1].id, monkeys[1].inspected)
	return monkeys[0].inspected * monkeys[1].inspected
}

func part2(monkeys []monkey) int {
	for i := 0; i < 10000; i++ {
		monkeys = inspectOnce2(monkeys)
	}

	/*
		for i := 0; i < len(monkeys); i++ {
			fmt.Println(monkeys[i].items)
		}
	*/

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].inspected > monkeys[j].inspected })

	//fmt.Printf("monkey %d: %d, monkey %d: %d\n", monkeys[0].id, monkeys[0].inspected, monkeys[1].id, monkeys[1].inspected)
	return monkeys[0].inspected * monkeys[1].inspected
}

func parseMonkey(input string) monkey {
	var it, opS, opNS string
	var test, T, F, id int
	lines := strings.Split(input, "\n")
	fmt.Sscanf(lines[0], "Monkey %d:", &id)
	aux1 := strings.Split(lines[1], ": ")
	it = aux1[1]
	fmt.Sscanf(lines[2], "  Operation: new = old %s %s", &opS, &opNS)
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &test)
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &T)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &F)

	//fmt.Println(it)
	var op operation
	var opN int64
	if opS == "*" && opNS == "old" {
		op = SQR
		opN = -1
	} else if opS == "*" {
		op = MUL
		aux, _ := strconv.Atoi(opNS)
		opN = int64(aux)
	} else {
		op = SUM
		aux, _ := strconv.Atoi(opNS)
		opN = int64(aux)
	}

	items := []int64{}
	itemsS := strings.Split(it, ", ")
	for _, item := range itemsS {
		aux, _ := strconv.Atoi(item)
		items = append(items, int64(aux))
	}

	return monkey{id: id, items: items, op: op, opN: opN, test: test, T: T, F: F, inspected: 0}
}
