// days/day19/day19.go

package day19

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type blueprint struct {
	id       int
	ore      int
	clay     int
	obsidian [2]int
	geode    [2]int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day19/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")

	input := parseInput(inputS)

	fmt.Println(input)

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

func parseInput(inputS []string) []blueprint {
	var blueprints []blueprint
	for _, str := range inputS {
		var b blueprint
		fmt.Sscanf(str, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.id, &b.ore, &b.clay, &b.obsidian[0], &b.obsidian[1], &b.geode[0], &b.geode[1])
		blueprints = append(blueprints, b)
	}
	return blueprints
}

// Tried to parallelise it but didn't work oops
func part1(input []blueprint) int {
	sol := make([]int, len(input))
	//var wg sync.WaitGroup
	//wg.Add(len(input))
	for i, b := range input {
		//go func(i int) {
		//wg.Done()
		sol[i] = qualityLvl(b)
		//}(i)
	}
	//wg.Wait()
	goodSol := 0
	for _, val := range sol {
		goodSol += val
	}
	return goodSol
}

func part2(input []blueprint) int {
	robots := [4]int{1, 0, 0, 0}
	minerals := [4]int{0, 0, 0, 0}
	qualities := [3]int{}
	for i := 0; i < 3; i++ {
		best := 0
		qualities[i] = qualityAux(input[i], robots, minerals, 32, &best, [3]bool{true, true, true})
	}
	sol := qualities[0] * qualities[1] * qualities[2]

	return sol
}

func qualityLvl(b blueprint) int {
	robots := [4]int{1, 0, 0, 0}
	minerals := [4]int{0, 0, 0, 0}
	best := 0
	q := qualityAux(b, robots, minerals, 24, &best, [3]bool{true, true, true})
	return b.id * q
}

func qualityAux(b blueprint, robots, minerals [4]int, min int, best *int, blocked [3]bool) int {

	// Stop if we can't get enough geodes than the best
	// found so far
	possibleGeodes := minerals[3]
	for i := 0; i < min; i++ {
		possibleGeodes += robots[3] + i
	}
	if possibleGeodes < *best {
		return 0
	}

	// Si el tiempo se agota terminamos.
	if min == 0 {
		return minerals[3]
	}

	// Si el número de robots de obsidiana y ore es mayor o igual
	// que los que refieren los geodes, entonces solo vamos a construir geodes
	if robots[0] >= b.geode[0] && robots[2] >= b.geode[1] {
		nGeodes := minerals[3]
		for i := 0; i < min; i++ {
			nGeodes += robots[3] + i
		}
		return nGeodes
	}

	newMinerals := [4]int{}
	newRobots := [4]int{}

	// Si podemos construir un robot de geodes lo hacemos:
	if minerals[0] >= b.geode[0] && minerals[2] >= b.geode[1] {
		for i := range minerals {
			newMinerals[i] = minerals[i] + robots[i]
		}
		newMinerals[0] -= b.geode[0]
		newMinerals[2] -= b.geode[1]
		newRobots = robots
		newRobots[3]++

		nGeodes := qualityAux(b, newRobots, newMinerals, min-1, best, [3]bool{true, true, true})
		if nGeodes > *best {
			*best = nGeodes
		}
		return nGeodes
	}

	// Si no nos da tiempo a construir un robot de geodes nuevo
	// por no tener obsidiana o ores devolver el maximo de geodes

	// Teniendo en cuenta que no construimos un robot de geodes
	// podemos pillar obsidiana
	// obisidiana = obsidiana (min ) + obsidiana (min -1)  + obsidiana + 1 (min - 2) + ...
	possibleObsidian := minerals[2]
	for i := 0; i < min-1; i++ {
		possibleObsidian += robots[2] + i
	}
	if possibleObsidian <= b.geode[1] {
		return robots[3]*min + minerals[3]
	}

	maxGeodes := 0

	// Miramos el resto de nuestras opciones y elegimos la mejor:

	// Si tenemos mas numero de robots de cualquier elemento
	// de los que hagan falta para construir algo entonces
	// no construimos robot de ese elemento

	maxRobotsOre := utils.Max(b.ore, utils.Max(b.clay, utils.Max(b.obsidian[0], b.geode[0])))
	maxRobotsClay := b.obsidian[1]
	maxRobotsObsidian := b.geode[1]

	// Guardamos también si podemos construir o no robots
	// Ya que si no construimos ningún robot pero podíamos construir alguno
	// En el siguiente minuto no tenemos que volverlo a contruir

	// Construimos robot de ore
	canConstructOre := true
	if minerals[0] >= b.ore && robots[0] <= maxRobotsOre && blocked[0] {
		canConstructOre = false
		for i := range minerals {
			newMinerals[i] = minerals[i] + robots[i]
		}
		newMinerals[0] -= b.ore
		newRobots = robots
		newRobots[0]++
		maxGeodes = utils.Max(maxGeodes, qualityAux(b, newRobots, newMinerals, min-1, best, [3]bool{true, true, true}))
	}

	// Construimos robot de clay
	canConstructClay := true
	if minerals[0] >= b.clay && robots[1] <= maxRobotsClay && blocked[1] {
		canConstructClay = false
		for i := range minerals {
			newMinerals[i] = minerals[i] + robots[i]
		}
		newMinerals[0] -= b.clay
		newRobots = robots
		newRobots[1]++
		maxGeodes = utils.Max(maxGeodes, qualityAux(b, newRobots, newMinerals, min-1, best, [3]bool{true, true, true}))
	}

	// Construimos robot de obsidiana
	canConstructObsidian := true
	if minerals[0] >= b.obsidian[0] && minerals[1] >= b.obsidian[1] && robots[2] <= maxRobotsObsidian && blocked[2] {
		canConstructObsidian = false
		for i := range minerals {
			newMinerals[i] = minerals[i] + robots[i]
		}
		newMinerals[0] -= b.obsidian[0]
		newMinerals[1] -= b.obsidian[1]
		newRobots = robots
		newRobots[2]++
		maxGeodes = utils.Max(maxGeodes, qualityAux(b, newRobots, newMinerals, min-1, best, [3]bool{true, true, true}))
	}

	// No construimos ningun robot
	for i := range minerals {
		newMinerals[i] = minerals[i] + robots[i]
	}
	// If we don't construct any robot, we block the ones that we could construct before
	maxGeodes = utils.Max(maxGeodes, qualityAux(b, robots, newMinerals, min-1, best, [3]bool{canConstructOre, canConstructClay, canConstructObsidian}))

	if maxGeodes > *best {
		*best = maxGeodes
	}
	return maxGeodes

}
