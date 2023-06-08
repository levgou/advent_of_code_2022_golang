package day_16

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"math/big"
	"strings"
	"time"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_16/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_16/demo_input.txt"
)

type Valve struct {
	flow       int
	neighbours []string
}

func parse(lines []string) map[string]Valve {
	// parse lines that look like:
	// Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE

	var valves = map[string]Valve{}

	for _, line := range lines {
		parts := strings.Split(line, " ")

		curValve := parts[1]

		// parse the flow rate
		flowRate := shared.ParseInt(
			strings.TrimSuffix(strings.TrimPrefix(parts[4], `rate=`), ";"))

		// parse the neighbours
		neighbours := make([]string, 0)
		for j := 9; j < len(parts); j++ {
			n := strings.TrimSuffix(parts[j], ",")
			neighbours = append(neighbours, n)
		}

		valves[curValve] = Valve{flowRate, neighbours}
	}

	return valves
}

type ValveAtTime struct {
	valve string
	time  int
}

const MINUTES = 30

var MEMO = map[string]int{}
var FLOW_SUM = 0

func findBestSolution(valves map[string]Valve) int {
	curValve := "AA"
	curFlow := 0
	openValves := mapset.NewSet[string]()

	minLeft := MINUTES

	for _, v := range valves {
		FLOW_SUM += v.flow
	}

	//return findBestSolutionRecursive2(
	//	curValve, curFlow, 0, openValves, &valves, minLeft)

	return findBestSolutionTwoPlayers(
		curValve, curValve, curFlow, 0, openValves, &valves, minLeft-4, true)

}

var pruned = 0
var noPruned = 0
var prunedSmart = 0

var bestSolution = 0

func findBestSolutionTwoPlayers(
	curValveA string,
	curValveB string,
	curFlow int,
	pressureReleased int,
	openValves mapset.Set[string],
	valves *map[string]Valve,
	minLeft int,
	turnA bool,
) int {
	if minLeft == 0 && turnA {
		if pressureReleased > bestSolution {
			//fmt.Println("NEW BEST", pressureReleased)
			bestSolution = pressureReleased
		}
		return pressureReleased
	}

	key := fmt.Sprintf(
		"%t-%s-%s-%s-%d", turnA, curValveA, curValveB, openValves.String(), minLeft)
	if val, ok := MEMO[key]; ok {
		pruned++
		return val
	} else {
		noPruned++
	}

	//prune in case we magically opened all valves but still got
	//worse than the best result so far
	if turnA && pressureReleased+minLeft*FLOW_SUM < bestSolution {
		//fmt.Println(
		//	pressureReleased, minLeft, FLOW_SUM,
		//	pressureReleased+minLeft*FLOW_SUM, bestSolution)

		prunedSmart++
		return 0
	}
	//else {
	//	fmt.Println("NO PRUNE", pressureReleased+minLeft*FLOW_SUM, bestSolution)
	//}

	possibilities := []int{}

	newPressureReleased := pressureReleased
	newTime := minLeft
	if !turnA {
		newPressureReleased += curFlow
		newTime = minLeft - 1
	}

	if turnA {
		valveFlowA := (*valves)[curValveA].flow
		// opened a valve and didnt move
		if valveFlowA > 0 && !openValves.Contains(curValveA) {
			open := openValves.Clone()
			open.Add(curValveA)
			newFlow := curFlow + valveFlowA
			possibilities = append(possibilities, findBestSolutionTwoPlayers(
				curValveA, curValveB, newFlow, newPressureReleased, open, valves, newTime, !turnA))
		}

		// move to a new valve
		for _, neighbour := range (*valves)[curValveA].neighbours {
			possibilities = append(possibilities, findBestSolutionTwoPlayers(
				neighbour, curValveB, curFlow, newPressureReleased, openValves, valves, newTime, !turnA))
		}

		bestVal := shared.Max(possibilities)
		MEMO[key] = bestVal

		return bestVal
	} else {
		valveFlowB := (*valves)[curValveA].flow
		// opened a valve and didnt move
		if valveFlowB > 0 && !openValves.Contains(curValveB) {
			open := openValves.Clone()
			open.Add(curValveB)
			newFlow := curFlow + valveFlowB
			possibilities = append(possibilities, findBestSolutionTwoPlayers(
				curValveA, curValveB, newFlow, newPressureReleased, open, valves, newTime, !turnA))
		}

		// move to a new valve
		for _, neighbour := range (*valves)[curValveB].neighbours {
			possibilities = append(possibilities, findBestSolutionTwoPlayers(
				curValveA, neighbour, curFlow, newPressureReleased, openValves, valves, newTime, !turnA))
		}

		bestVal := shared.Max(possibilities)
		MEMO[key] = bestVal

		return bestVal
	}
}

func findBestSolutionRecursive2(
	curValve string,
	curFlow int,
	pressureReleased int,
	openValves mapset.Set[string],
	valves *map[string]Valve,
	minLeft int,
) int {

	if minLeft == 0 {
		if pressureReleased > bestSolution {
			//fmt.Println("NEW BEST", pressureReleased)
			bestSolution = pressureReleased
		}
		return pressureReleased
	}

	key := fmt.Sprintf("%s-%s-%d", curValve, openValves.String(), minLeft)
	if val, ok := MEMO[key]; ok {
		pruned++
		return val
	} else {
		noPruned++
	}

	//prune in case we magically opened all valves but still got
	//worse than the best result so far
	if pressureReleased+minLeft*FLOW_SUM < bestSolution {
		//fmt.Println(
		//	pressureReleased, minLeft, FLOW_SUM,
		//	pressureReleased+minLeft*FLOW_SUM, bestSolution)

		prunedSmart++
		return 0
	}
	//else {
	//	fmt.Println("NO PRUNE", pressureReleased+minLeft*FLOW_SUM, bestSolution)
	//}

	possibilities := []int{}
	newTime := minLeft - 1
	newPressureReleased := pressureReleased + curFlow

	valveFlow := (*valves)[curValve].flow

	// opened a valve and didnt move
	if valveFlow > 0 && !openValves.Contains(curValve) {
		open := openValves.Clone()
		open.Add(curValve)
		newFlow := curFlow + valveFlow
		possibilities = append(possibilities, findBestSolutionRecursive2(
			curValve, newFlow, newPressureReleased, open, valves, newTime))
	}

	// move to a new valve
	for _, neighbour := range (*valves)[curValve].neighbours {
		possibilities = append(possibilities, findBestSolutionRecursive2(
			neighbour, curFlow, newPressureReleased, openValves, valves, newTime))
	}

	bestVal := shared.Max(possibilities)
	MEMO[key] = bestVal

	return bestVal
}

func Solution() {
	lines, err := shared.ReadLines(Demo)
	if err != nil {
		log.Fatal(err)
		return
	}

	valves := parse(lines)
	start := time.Now()

	pressureReleased := findBestSolution(valves)

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)
	fmt.Println(pruned, prunedSmart, noPruned)

	fmt.Println(pressureReleased)
}
