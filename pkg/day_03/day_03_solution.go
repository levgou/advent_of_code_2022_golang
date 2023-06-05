package day_03

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_03/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_03/input.txt"
)

const (
	LowercaseA = int('a')
	LowercaseZ = int('z')
	UppercaseA = int('A')
	UppercaseZ = int('Z')
)

func charVal(c string) int {
	val := int(c[0])

	if val >= LowercaseA && val <= LowercaseZ {
		return val - LowercaseA + 1
	} else if val >= UppercaseA && val <= UppercaseZ {
		return val - UppercaseA + 27
	} else {
		return 0
	}
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	commonSum := 0
	for _, line := range lines {
		l := len(line)
		left := line[:l/2]
		right := line[l/2:]

		leftSet := mapset.NewSet(strings.Split(right, "")...)
		rightSet := mapset.NewSet(strings.Split(left, "")...)

		common, hasVal := leftSet.Intersect(rightSet).Pop()
		if !hasVal {
			log.Fatal("No common value found")
			return
		}

		commonSum += charVal(common)
	}

	fmt.Println(commonSum)
}

func Solution2() {
	lines, err := shared.ReadLines(Input)

	if err != nil {
		log.Fatal(err)
		return
	}

	commonSum := 0
	groups := shared.ChunkSlice(lines, 3)
	for _, group := range groups {
		first := mapset.NewSet(strings.Split(group[0], "")...)
		second := mapset.NewSet(strings.Split(group[1], "")...)
		third := mapset.NewSet(strings.Split(group[2], "")...)

		common, valExists := first.Intersect(second).Intersect(third).Pop()
		if !valExists {
			log.Fatal("No common value found")
			return
		}

		commonSum += charVal(common)
	}

	fmt.Println(commonSum)
}
