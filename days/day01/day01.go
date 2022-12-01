// day01/day01.go

package day01

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
	bs, err := ioutil.ReadFile("./days/day01/input.txt")

	fmt.Println(err)

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := parseGood(bs)

	sol1, sol2 := part(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part(input [][]int) (int, int) {
	var max [3]int
	for _, row := range input {
		value := 0
		for _, val := range row {
			value += val
		}
		if max[0] <= value {
			max[2] = max[1]
			max[1] = max[0]
			max[0] = value
		} else if max[1] <= value {
			max[2] = max[1]
			max[1] = value
		} else if max[2] <= value {
			max[2] = value
		}
	}
	fmt.Println(max)
	return max[0], max[0] + max[1] + max[2]
}

func parseGood(bs []byte) [][]int {
	input := strings.Split(string(bs), "\n\n")

	var result [][]int
	for _, value := range input {
		var vals []int
		for _, val := range strings.Split(value, "\n") {
			val1, _ := strconv.ParseInt(val, 10, 64)
			vals = append(vals, int(val1))
		}
		result = append(result, vals)
	}

	return result
}
