// days/day15/day15.go

package day15

import (
	"AOC2022-Go/utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type position struct {
	x, y int
}

func Solve() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("./days/day15/input.txt")

	if err != nil {
		bs, _ = ioutil.ReadFile("input.txt")
	}

	inputS := strings.Split(string(bs), "\n")
	sensors := make([]position, len(inputS))
	beacons := make([]position, len(inputS))
	for i, inp := range inputS {
		var sensor, beacon position
		fmt.Sscanf(inp, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		sensors[i] = sensor
		beacons[i] = beacon
	}

	sol1 := part1(sensors, beacons)
	sol2 := part2(sensors, beacons)

	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)
}

func manhattanDist(p1, p2 position) int {
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func part1(sensors, beacons []position) int {
	newIntervals := getLine(sensors, beacons, 2000000)
	nBeacons := 0
	for _, inter := range newIntervals {
		nBeacons += inter.Len()
	}

	for _, beacon := range utils.UniqueConv(beacons) {
		if beacon.y == 2000000 {
			nBeacons--
		}
	}

	return nBeacons
}

func part2(sensors, beacons []position) int64 {
	var sol int64
	for y := 0; y <= 4000000; y++ {
		newIntervals := getLine(sensors, beacons, y)
		if len(newIntervals) == 2 {
			x := newIntervals[0].R + 1
			sol = int64(x)*4000000 + int64(y)
			break
		}
	}
	return sol
}

func getLine(sensors, beacons []position, yPos int) []utils.Interval {
	/*
		if yPos%40000 == 0 {
			fmt.Println(yPos / 40000)
		}
	*/
	intervals := []utils.Interval{}
	for i, sensor := range sensors {
		// Hallamos el intervalo que miran
		beacon := beacons[i]
		dist := manhattanDist(beacon, sensor)
		distY := int(math.Abs(float64(sensor.y - yPos)))
		newWidth := dist - distY
		if newWidth >= 0 {
			intervals = append(intervals, utils.Interval{L: sensor.x - newWidth, R: sensor.x + newWidth})
		}
	}

	//fmt.Println(intervals)

	newIntervals := utils.SweepLine(intervals)

	//fmt.Println(newIntervals)

	/*
		nBeacons := 0
		for _, inter := range newIntervals {
			nBeacons += inter.Len()
		}

		for _, beacon := range utils.UniqueConv(beacons) {
			if beacon.y == yPos {
				nBeacons--
			}
		}
	*/

	return newIntervals
}
