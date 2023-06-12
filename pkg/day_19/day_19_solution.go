package day_19

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"log"
	"strings"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_19/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_19/demo_input.txt"
)

type AllResources struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type Resources struct {
	ore      int
	clay     int
	obsidian int
}

type Blueprint struct {
	oreRobot      Resources
	clayRobot     Resources
	obsidianRobot Resources
	geodeRobot    Resources
}

func parse(lines []string) []Blueprint {
	blueprints := make([]Blueprint, 0)
	for _, line := range lines {
		instructions := strings.Split(strings.Split(line, ":")[1], ".")
		blueprint := Blueprint{
			oreRobot: Resources{
				ore:      shared.ParseInt(strings.Split(instructions[0], " ")[5]),
				clay:     0,
				obsidian: 0,
			},

			clayRobot: Resources{
				ore:      shared.ParseInt(strings.Split(instructions[1], " ")[5]),
				clay:     0,
				obsidian: 0,
			},

			obsidianRobot: Resources{
				ore:      shared.ParseInt(strings.Split(instructions[2], " ")[5]),
				clay:     shared.ParseInt(strings.Split(instructions[2], " ")[8]),
				obsidian: 0,
			},

			geodeRobot: Resources{
				ore:      shared.ParseInt(strings.Split(instructions[3], " ")[5]),
				clay:     0,
				obsidian: shared.ParseInt(strings.Split(instructions[3], " ")[8]),
			},
		}
		blueprints = append(blueprints, blueprint)
	}

	return blueprints
}

type MemoKey struct {
	Resources AllResources
	robots    AllResources
	timeLeft  int
}

var memHit = 0
var memHit2 = 0
var memMiss = 0

func findBestBlueprintRec(
	blueprint Blueprint,
	resources AllResources,
	robots AllResources,
	timeLeft int,
	memo *map[MemoKey]int,
	mostGeodeRobotsAtTime *map[int][2]int,
) int {

	if timeLeft == 0 {
		return resources.geode
	}

	for moreTime := timeLeft; moreTime <= 24; moreTime++ {
		if val, ok := (*mostGeodeRobotsAtTime)[moreTime]; ok {
			if val[0] > robots.geode && val[1] > resources.geode {
				//fmt.Println("hit", timeLeft, moreTime, val, robots)
				memHit2 += 1
				return 0
			}
		}
	}
	(*mostGeodeRobotsAtTime)[timeLeft] = [2]int{robots.geode, resources.geode}

	memKey := MemoKey{resources, robots, timeLeft}
	if val, ok := (*memo)[memKey]; ok {
		memHit += 1
		return val
	} else {
		memMiss += 1
	}

	otherSteps := make([]int, 0)

	updatedResources := AllResources{
		ore:      resources.ore + robots.ore,
		clay:     resources.clay + robots.clay,
		obsidian: resources.obsidian + robots.obsidian,
		geode:    resources.geode + robots.geode,
	}

	// when we can construct a geoRobot itll be the best option
	if blueprint.geodeRobot.ore <= resources.ore &&
		blueprint.geodeRobot.obsidian <= resources.obsidian {
		withGeodeRobot := robots
		withGeodeRobot.geode += 1
		withGeodeRobotResources := updatedResources
		withGeodeRobotResources.ore -= blueprint.geodeRobot.ore
		withGeodeRobotResources.obsidian -= blueprint.geodeRobot.obsidian
		makeGeodeRobot := findBestBlueprintRec(blueprint, withGeodeRobotResources, withGeodeRobot, timeLeft-1, memo, mostGeodeRobotsAtTime)
		return makeGeodeRobot
	}

	waitResult := findBestBlueprintRec(blueprint, updatedResources, robots, timeLeft-1, memo, mostGeodeRobotsAtTime)
	otherSteps = append(otherSteps, waitResult)

	if blueprint.oreRobot.ore <= resources.ore {
		withOreRobot := robots
		withOreRobot.ore += 1
		withOreRobotResources := updatedResources
		withOreRobotResources.ore -= blueprint.oreRobot.ore
		makeOreRobot := findBestBlueprintRec(blueprint, withOreRobotResources, withOreRobot, timeLeft-1, memo, mostGeodeRobotsAtTime)
		otherSteps = append(otherSteps, makeOreRobot)
	}

	if blueprint.clayRobot.ore <= resources.ore {
		withClayRobot := robots
		withClayRobot.clay += 1
		withClayRobotResources := updatedResources
		withClayRobotResources.ore -= blueprint.clayRobot.ore
		makeClayRobot := findBestBlueprintRec(blueprint, withClayRobotResources, withClayRobot, timeLeft-1, memo, mostGeodeRobotsAtTime)
		otherSteps = append(otherSteps, makeClayRobot)
	}

	if blueprint.obsidianRobot.ore <= resources.ore &&
		blueprint.obsidianRobot.clay <= resources.clay {
		withObsidianRobot := robots
		withObsidianRobot.obsidian += 1
		withObsidianRobotResources := updatedResources
		withObsidianRobotResources.ore -= blueprint.obsidianRobot.ore
		withObsidianRobotResources.clay -= blueprint.obsidianRobot.clay
		makeObsidianRobot := findBestBlueprintRec(blueprint, withObsidianRobotResources, withObsidianRobot, timeLeft-1, memo, mostGeodeRobotsAtTime)
		otherSteps = append(otherSteps, makeObsidianRobot)
	}

	max := shared.Max(otherSteps)
	(*memo)[memKey] = max
	return max
}

func findBestBlueprint(blueprints []Blueprint) {

	scores := make([]int, 0)
	for i, bp := range blueprints[:3] {
		resouces := AllResources{0, 0, 0, 0}
		robots := AllResources{1, 0, 0, 0}
		memo := make(map[MemoKey]int)
		mostGeodeRobotsAtTime := make(map[int][2]int)

		score := findBestBlueprintRec(bp, resouces, robots, 32, &memo, &mostGeodeRobotsAtTime)
		fmt.Println(i, score)
		scores = append(scores, score)
	}

	fmt.Println(memHit, memHit2, memMiss)
	fmt.Println(scores)
	sum := 0
	for i, score := range scores {
		sum += score * (i + 1)
	}
	fmt.Println(sum)
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	blueprints := parse(lines)
	findBestBlueprint(blueprints)
}
