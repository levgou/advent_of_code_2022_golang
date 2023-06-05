package day_06

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_06/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_06/input.txt"
)

const WindowSize = 14

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	msg := lines[0]

	windowIdx := WindowSize

	for windowIdx < len(msg)+1 {
		window := msg[windowIdx-WindowSize : windowIdx]
		windowSet := mapset.NewSet(strings.Split(window, "")...)
		if windowSet.Cardinality() == WindowSize {
			fmt.Println(window, windowIdx)
			break
		}
		windowIdx++
	}
}
