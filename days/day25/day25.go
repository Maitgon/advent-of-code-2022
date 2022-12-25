// days/day25/day25.go

package day25

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var values = map[byte]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
var valuesInv = [5]byte{'0', '1', '2', '=', '-'}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day25/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n")

	sol := part1(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol)
	fmt.Println("Time: ", end)

}

func part1(input []string) string {
	solInt := int64(0)
	for _, val := range input {
		solInt += int64(snafuToInt(val))
	}
	return intToSnafu(solInt)
}

func snafuToInt(snafu string) int {
	fact := 1
	num := 0
	for i := len(snafu) - 1; i >= 0; i-- {
		num += fact * values[snafu[i]]
		fact *= 5
	}
	return num
}

func intToSnafu(num int64) string {
	sol := ""
	for num != 0 {
		newDigit := valuesInv[num%5]
		if newDigit == '=' || newDigit == '-' {
			num += 5
		}
		num /= 5
		sol = string(newDigit) + sol
	}
	return sol
}
