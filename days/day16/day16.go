// days/day16/day16.go

package day16

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type valve struct {
	flow    int
	open    bool
	tunnels []string
}

type path struct {
	value int
	p     uint16
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day16/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")
	//fmt.Println(inputS)
	input := parseInput(inputS)

	//sol1 := part1(input)
	//graph := floydWarshall(input)
	//fmt.Println(0xff)
	//fmt.Println(graph)

	sol1, sol2 := part(input)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}

// byte is the same as uint8
func floydWarshall(valves map[string]valve) map[string]map[string]byte {
	// Devuelve una matriz (como representacion de un grafo) con las distancias
	// entre cada valve
	graph := make(map[string]map[string]byte)
	// Rellenamos el grafo
	for i, valv := range valves {
		graph[i] = make(map[string]byte)
		for j := range valves {
			if utils.Contains(valv.tunnels, j) {
				graph[i][j] = 1
			} else if i == j {
				graph[i][i] = 0
			} else {
				// All nodes are connected somehow, so you don't need more
				// than a byte to represent the distance
				graph[i][j] = 100
			}
		}
	}

	for k := range valves {
		for i := range valves {
			for j := range valves {
				graph[i][j] = utils.Min(graph[i][j], graph[i][k]+graph[k][j])
			}
		}
	}

	//fmt.Println(graph["AA"]["NC"])
	//fmt.Println(graph)

	return graph
}

func parseInput(inputS []string) map[string]valve {
	regex := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w ,]+)`)
	valves := map[string]valve{}
	for _, valv := range inputS {
		res := regex.FindStringSubmatch(valv)
		name := res[1]
		flowRate, _ := strconv.Atoi(res[2])
		tunnelsS := res[3]

		tunnels := strings.Split(tunnelsS, ", ")
		valves[name] = valve{flow: flowRate, open: false, tunnels: tunnels}
	}
	return valves
}

func part(input map[string]valve) (int, int) {
	// We create the graph
	graph := floydWarshall(input)
	//fmt.Println(graph["AA"]["NC"])

	// We filter the valves that we are going to use
	goodValves := []string{}
	for name, valv := range input {
		if valv.flow > 0 || name == "AA" {
			goodValves = append(goodValves, name)
		}
	}

	//fmt.Println(goodValves)

	// There are 16 gooValves so we can mask every valve as a bit of a uint16
	bitMask := make(map[string]uint16)
	for i, v := range goodValves {
		bitMask[v] = 1 << i
	}

	var dfs func(minutes, totalFlow int, opened uint16, node string) int
	dfs = func(minutes, totalFlow int, opened uint16, node string) int {
		flow := 0
		for _, valv := range goodValves {
			// If the valve is at the start, the valve has been opened or
			// the valve is the same we dont go there
			if valv == "AA" || bitMask[valv]&opened != 0 || valv == node {
				continue
			}
			// minutos que tardas en llegar + minuto de abrirlo
			//fmt.Println(node, valv)
			//fmt.Println(graph[node][valv])
			newMinutes := minutes - int(graph[node][valv]) - 1

			// Te quedas con el mejor camino
			if newMinutes >= 0 {
				if aux := dfs(newMinutes, newMinutes*input[valv].flow, opened|bitMask[valv], valv); aux > flow {
					flow = aux
				}
			}
		}
		return flow + totalFlow
	}

	sol1 := int(dfs(30, 0, 0, "AA"))

	// For part 2, we can encode all the possible paths as uint16 and
	// check all pair of paths, we can check if two paths have been in the
	// same valve using path1&path2 != 0

	var dfs2 func(minutes, totalFlow int, opened uint16, node string, p uint16) []path
	dfs2 = func(minutes, totalFlow int, opened uint16, node string, p uint16) []path {
		paths := []path{{value: totalFlow, p: p}}
		for _, valv := range goodValves {
			// If the valve is at the start, the valve has been opened or
			// the valve is the same we dont go there
			if valv == "AA" || bitMask[valv]&opened != 0 || valv == node {
				continue
			}
			// minutos que tardas en llegar + minuto de abrirlo
			//fmt.Println(node, valv)
			//fmt.Println(graph[node][valv])
			newMinutes := minutes - int(graph[node][valv]) - 1

			// Te quedas con el mejor camino
			// To add a valve to a path, we use bitMask[valve] | path
			if newMinutes >= 0 {
				paths = append(paths, dfs2(newMinutes, totalFlow+newMinutes*input[valv].flow, opened|bitMask[valv], valv, p|bitMask[valv])...)
			}
		}
		return paths
	}

	// We reduce the paths, we assumen, that at least, we are having paths that are
	// 1/2 out initial solution. If you don't do this this goes for LONG

	paths := dfs2(26, 0, 0, "AA", 0)

	//fmt.Println(paths)

	var reducedPaths []path
	for _, p := range paths {
		if p.value > sol1/2 {
			reducedPaths = append(reducedPaths, p)
		}
	}

	//fmt.Println(reducedPaths)

	sol2 := 0
	for i := 0; i < len(reducedPaths); i++ {
		// We avoid getting the same path and getting the mirrored path
		for j := i + 1; j < len(reducedPaths); j++ {
			if reducedPaths[i].p&reducedPaths[j].p == 0 {
				if aux := reducedPaths[i].value + reducedPaths[j].value; aux > sol2 {
					//fmt.Println(reducedPaths[i].p, reducedPaths[j].p)
					sol2 = aux
				}
			}
		}
	}

	return sol1, sol2

}
