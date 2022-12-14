// days/day13/day13.go

package day13

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day13/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	input := strings.Split(string(bs), "\n\n")

	pairs := make([][2]utils.Tree, len(input))
	for i, pair := range input {
		aux := strings.Split(pair, "\n")
		pairs[i][0] = utils.ParseTree(&aux[0])
		pairs[i][1] = utils.ParseTree(&aux[1])
	}

	//fmt.Println(pairs[3][0].Show())
	//fmt.Println(pairs[3][1].Show())

	//fmt.Println(comp(pairs[3][0], pairs[3][1]))

	sol1 := part1(pairs)
	sol2 := part2(pairs)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(pairs [][2]utils.Tree) int {
	sol := 0
	for i, pair := range pairs {
		if comp(pair[0], pair[1]) == LE {
			sol += i + 1
		}
	}
	return sol
}

func part2(pairs [][2]utils.Tree) int {
	// We get all the pairs into a single array
	newArray := make([]utils.Tree, len(pairs)*2)
	for i, pair := range pairs {
		newArray[2*i] = pair[0]
		newArray[2*i+1] = pair[1]
	}

	// We create the dividers
	aux1 := []utils.Tree{{Val: 2, Tree: nil, Deep: 2}}
	aux2 := []utils.Tree{{Val: -1, Tree: &aux1, Deep: 1}}
	divider1 := utils.Tree{Val: -1, Tree: &aux2, Deep: 0}
	aux3 := []utils.Tree{{Val: 6, Tree: nil, Deep: 2}}
	aux4 := []utils.Tree{{Val: -1, Tree: &aux3, Deep: 1}}
	divider2 := utils.Tree{Val: -1, Tree: &aux4, Deep: 0}
	newArray = append(newArray, divider1, divider2)

	// We sort the new array
	sort.Slice(newArray, func(i, j int) bool { return comp(newArray[i], newArray[j]) == LE })

	sol := 1
	for i, tree := range newArray {
		if tree.IsDividerPacket() {
			//fmt.Println(tree.Show())
			sol *= i + 1
		}
	}

	return sol
}

func comp(iz, dr utils.Tree) compare {
	//fmt.Printf("Comparing: %s and %s, deeps: %d %d\n", iz.Show(), dr.Show(), iz.Deep, dr.Deep)
	if iz.IsVal() && dr.IsVal() {
		// Here both values are values
		//fmt.Println("Here bitch")
		if iz.Val == dr.Val {
			return EQ
		} else if iz.Val < dr.Val {
			return LE
		} else {
			return GT
		}
	} else if iz.IsLeaf() && dr.IsLeaf() {
		// Here both values are trees
		for i := range *iz.Tree {
			//fmt.Println(i)
			if i >= len(*dr.Tree) {
				//fmt.Println("No items left in dr")
				return GT
			}
			res := comp((*iz.Tree)[i], (*dr.Tree)[i])
			if res == LE || res == GT {
				return res
			}
		}
		if len(*dr.Tree) > len(*iz.Tree) {
			//fmt.Println("No items left in iz")
			return LE
		}
	} else {
		// If it goes here, it means that one is a value and the other is a tree
		//fmt.Println("Here, fucker")
		if iz.IsVal() {
			newTree := []utils.Tree{{Val: iz.Val, Tree: nil, Deep: iz.Deep + 1}}
			return comp(utils.Tree{Val: -1, Tree: &newTree, Deep: iz.Deep}, dr)
		} else {
			newTree := []utils.Tree{{Val: dr.Val, Tree: nil, Deep: dr.Deep + 1}}
			return comp(iz, utils.Tree{Val: -1, Tree: &newTree, Deep: iz.Deep})
		}
	}

	return EQ
}

type compare int

const (
	EQ compare = iota
	GT
	LE
)
