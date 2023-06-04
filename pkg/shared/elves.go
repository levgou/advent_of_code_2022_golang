package shared

import (
	"log"
	"strconv"
)

func ElfSums(portionedLines [][]string) []int {
	elves := []int{}

	for _, portion := range portionedLines {
		elf := 0

		for _, line := range portion {
			val, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
				return nil
			}

			elf += val
		}
		elves = append(elves, elf)
	}

	return elves
}
