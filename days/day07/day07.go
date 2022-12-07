// days/day07/day07.go

package day07

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// "AOC2022-Go/utils"

type directory struct {
	name string
	dirs *[]directory
	data []int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day07/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")
	p := 0
	parsed := parse(input, &p, "")

	sol1, sol2 := part(parsed)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func parse(input []string, p *int, currName string) directory {
	line := input[*p]
	command := strings.Split(line, " ")
	name := currName + command[2]
	dirs := &[]directory{}
	var data []int
	*p++
	for {
		if *p >= len(input) {
			break
		}
		line = input[*p]
		command = strings.Split(line, " ")

		if len(command) == 2 {
			if val, err := strconv.Atoi(command[0]); err == nil {
				data = append(data, val)
			}
			*p++
		} else if len(command) == 3 {
			if command[2] == ".." {
				*p++
				break
			}
			*dirs = append(*dirs, parse(input, p, name))
		}
	}
	return directory{name: name, dirs: dirs, data: data}
}

func part(dir directory) (int, int) {
	vals := map[string]int{}
	fillMap(dir, &vals)
	sol1 := 0
	for _, val := range vals {
		if val < 100000 {
			sol1 += val
		}
	}

	limit := 30000000 - (70000000 - vals["/"])
	sol2 := 999999999
	for _, size := range vals {
		if size >= limit && size < sol2 {
			sol2 = size
		}
	}

	return sol1, sol2
}

func fillMap(dir directory, vals *map[string]int) {
	valueOf := 0
	for _, subDir := range *dir.dirs {
		fillMap(subDir, vals)
		valueOf += (*vals)[subDir.name]
	}
	for _, size := range dir.data {
		valueOf += size
	}
	(*vals)[dir.name] = valueOf
}
