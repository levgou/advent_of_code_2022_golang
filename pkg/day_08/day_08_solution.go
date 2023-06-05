package day_08

import (
	"advent_of_code/pkg/shared"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_08/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_08/input.txt"
)

func emptyMatrixSize(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func Solution() {
	lines, err := shared.ReadLines(DemoInput)
	if err != nil {
		log.Fatal(err)
		return
	}

	grid := make([][]int64, len(lines))
	for i, line := range lines {
		grid[i] = make([]int64, len(line))
		splitLine := strings.Split(line, "")
		for j, num := range splitLine {
			height, _ := strconv.ParseInt(num, 10, 64)
			grid[i][j] = height
		}
	}

	maxFromLeft := emptyMatrixSize(len(grid), len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if col == 0 {
				continue
			}
			maxFromLeft[row][col] =
				int(math.Max(float64(maxFromLeft[row][col-1]), float64(grid[row][col-1])))
		}
	}

	maxFromRight := emptyMatrixSize(len(grid), len(grid[0]))
	for row := len(grid) - 1; row >= 0; row-- {
		for col := len(grid[0]) - 1; col >= 0; col-- {
			if col == len(grid[0])-1 {
				continue
			}
			maxFromRight[row][col] =
				int(math.Max(float64(maxFromRight[row][col+1]), float64(grid[row][col+1])))
		}
	}

	maxFromTop := emptyMatrixSize(len(grid), len(grid[0]))
	for row := 1; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			maxFromTop[row][col] =
				int(math.Max(float64(maxFromTop[row-1][col]), float64(grid[row-1][col])))
		}
	}

	maxFromBottom := emptyMatrixSize(len(grid), len(grid[0]))
	for row := len(grid) - 2; row >= 0; row-- {
		for col := 0; col < len(grid[0]); col++ {
			maxFromBottom[row][col] =
				int(math.Max(float64(maxFromBottom[row+1][col]), float64(grid[row+1][col])))
		}
	}

	seenCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[0])-1 {
				seenCount++
			} else {
				val := grid[row][col]
				if val > int64(maxFromLeft[row][col]) ||
					val > int64(maxFromRight[row][col]) ||
					val > int64(maxFromTop[row][col]) ||
					val > int64(maxFromBottom[row][col]) {
					seenCount++
				}
			}
		}
	}

	log.Printf("Solution: %d", seenCount)
}
func Solution2() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	grid := make([][]int64, len(lines))
	for i, line := range lines {
		grid[i] = make([]int64, len(line))
		splitLine := strings.Split(line, "")
		for j, num := range splitLine {
			height, _ := strconv.ParseInt(num, 10, 64)
			grid[i][j] = height
		}
	}

	maxFromLeft := emptyMatrixSize(len(grid), len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if col == 0 {
				continue
			}

			seen := 0
			for colBack := col - 1; colBack >= 0; colBack-- {
				seen++
				if grid[row][col] > grid[row][colBack] {
					continue
				} else {
					break
				}
			}

			maxFromLeft[row][col] = seen
		}
	}

	maxFromRight := emptyMatrixSize(len(grid), len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := len(grid[0]) - 2; col >= 0; col-- {

			seen := 0
			for colBack := col + 1; colBack < len(grid[0]); colBack++ {
				seen++
				if grid[row][col] > grid[row][colBack] {
					continue
				} else {
					break
				}
			}

			maxFromRight[row][col] = seen
		}
	}

	maxFromTop := emptyMatrixSize(len(grid), len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {

			seen := 0
			for rowBack := row - 1; rowBack >= 0; rowBack-- {
				seen++
				if grid[row][col] > grid[rowBack][col] {
					continue
				} else {
					break
				}
			}

			maxFromTop[row][col] = seen
		}
	}

	maxFromBottom := emptyMatrixSize(len(grid), len(grid[0]))
	for row := len(grid) - 2; row >= 0; row-- {
		for col := 0; col < len(grid[0]); col++ {

			seen := 0
			for rowBack := row + 1; rowBack < len(grid); rowBack++ {
				seen++
				if grid[row][col] > grid[rowBack][col] {
					continue
				} else {
					break
				}
			}

			maxFromBottom[row][col] = seen
		}
	}

	maxArea := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			curArea := maxFromTop[row][col] * maxFromBottom[row][col] * maxFromLeft[row][col] * maxFromRight[row][col]
			if curArea > maxArea {
				maxArea = curArea
			}
		}
	}

	log.Printf("Solution: %d", maxArea)
}
