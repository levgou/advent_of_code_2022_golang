package day_04

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_04/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_04/input.txt"
)

func rangeContainsOtherRange(firstStart, firstEnd, secondStart, secondEnd int64) bool {
	return firstStart <= secondStart && firstEnd >= secondEnd
}

func anyRangeContainsOther(firstStart, firstEnd, secondStart, secondEnd int64) bool {
	return rangeContainsOtherRange(firstStart, firstEnd, secondStart, secondEnd) ||
		rangeContainsOtherRange(secondStart, secondEnd, firstStart, firstEnd)
}

func rangesOverlap(firstStart, firstEnd, secondStart, secondEnd int64) bool {
	return firstStart <= secondEnd && secondStart <= firstEnd
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	containNum := 0
	overlapNum := 0
	for _, pair := range lines {
		splitPair := strings.Split(pair, ",")
		first := splitPair[0]
		second := splitPair[1]

		firstSplitRange := strings.Split(first, "-")
		secondSplitRange := strings.Split(second, "-")

		firstStart, err1 := strconv.ParseInt(firstSplitRange[0], 10, 64)
		firstEnd, err2 := strconv.ParseInt(firstSplitRange[1], 10, 64)
		secondStart, err3 := strconv.ParseInt(secondSplitRange[0], 10, 64)
		secondEnd, err4 := strconv.ParseInt(secondSplitRange[1], 10, 64)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal(err1, err2, err3, err4)
			return
		}

		if anyRangeContainsOther(firstStart, firstEnd, secondStart, secondEnd) {
			containNum++
		}

		if rangesOverlap(firstStart, firstEnd, secondStart, secondEnd) {
			overlapNum++
		}
	}

	fmt.Println(containNum, overlapNum)
}
