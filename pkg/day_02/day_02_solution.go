package day_02

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"log"
	"strings"
)

const InputFilePath = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_02/input.txt"
const DemoInputFilePath = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_02/demo_input.txt"

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"
)

const (
	SelfRock     = "X"
	SelfPaper    = "Y"
	SelfScissors = "Z"
)

const (
	Lose = "X"
	Tie  = "Y"
	Win  = "Z"
)

var values = map[string]int{
	SelfRock:         1,
	OpponentRock:     1,
	SelfPaper:        2,
	OpponentPaper:    2,
	SelfScissors:     3,
	OpponentScissors: 3,
}

func winRound(self string, opponent string) bool {
	return self == SelfRock && opponent == OpponentScissors ||
		self == SelfPaper && opponent == OpponentRock ||
		self == SelfScissors && opponent == OpponentPaper
}

func tieRound(self string, opponent string) bool {
	return self == SelfRock && opponent == OpponentRock ||
		self == SelfPaper && opponent == OpponentPaper ||
		self == SelfScissors && opponent == OpponentScissors
}

func roundScore(self string, opponent string) int {
	if winRound(self, opponent) {
		return 6
	} else if tieRound(self, opponent) {
		return 3
	} else {
		return 0
	}
}

func scoreBasedOnResult(result string) int {
	if result == Win {
		return 6
	} else if result == Tie {
		return 3
	} else {
		return 0
	}
}

var moveBasedOnResult = map[string]map[string]string{
	Win: {
		OpponentRock:     SelfPaper,
		OpponentPaper:    SelfScissors,
		OpponentScissors: SelfRock,
	},
	Tie: {
		OpponentRock:     SelfRock,
		OpponentPaper:    SelfPaper,
		OpponentScissors: SelfScissors,
	},
	Lose: {
		OpponentRock:     SelfScissors,
		OpponentPaper:    SelfRock,
		OpponentScissors: SelfPaper,
	},
}

func Solution() {
	lines, err := shared.ReadLines(InputFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	score := 0

	for _, line := range lines {
		round := 0

		split := strings.Split(line, " ")
		opponent, self := split[0], split[1]

		round += values[self]
		round += roundScore(self, opponent)

		score += round
	}

	fmt.Println(score)
}

func Solution2() {

	lines, err := shared.ReadLines(InputFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	score := 0

	for _, line := range lines {
		round := 0

		split := strings.Split(line, " ")
		opponent, result := split[0], split[1]

		round += scoreBasedOnResult(result)
		move := moveBasedOnResult[result][opponent]
		round += values[move]

		score += round
	}

	fmt.Println(score)
}
