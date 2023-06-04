package day_01

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"log"
)

const InputFilePath = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_01/input.txt"

func Solution() {
	lines, err := shared.ReadLines(InputFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	portionedLines := shared.GroupLines(lines)
	elves := shared.ElfSums(portionedLines)

	max := shared.Max(elves)

	fmt.Println(max)
}

func Solution2() {
	lines, err := shared.ReadLines(InputFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	portionedLines := shared.GroupLines(lines)
	elves := shared.ElfSums(portionedLines)

	max3, err := shared.MaxN(elves, 3)

	fmt.Println(shared.Sum(max3))
}
