package day_14

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"log"
	"strings"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_14/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_14/demo_input.txt"
)

type Point struct {
	x int
	y int
}

var sendOrigin = Point{500, 0}

func readStructures(lines []string) ([][]Point, int, int, int, int) {
	var structures [][]Point
	maxX, maxY, minX, minY := 0, 0, 1000000, 1000000

	for _, line := range lines {
		points := make([]Point, 0)
		pointsStr := strings.Split(line, "->")
		for i := 0; i < len(pointsStr); i++ {
			pointsStr[i] = strings.TrimSpace(pointsStr[i])
		}

		for _, pointStr := range pointsStr {
			parts := strings.Split(pointStr, ",")

			x := shared.ParseInt(parts[0])
			y := shared.ParseInt(parts[1])

			if x > maxX {
				maxX = x
			}
			if x < minX {
				minX = x
			}
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y

			}

			points = append(points, Point{x, y})
		}

		structures = append(structures, points)
	}

	const floorBuffer = 800

	floorStructure := []Point{
		{minX - floorBuffer, maxY + 2},
		{maxX + floorBuffer, maxY + 2},
	}

	structures = append(structures, floorStructure)

	return structures, maxX + floorBuffer, maxY + 2, minX - floorBuffer, minY
}

func createWorld(minX int, maxX int, minY int, maxY int) ([][]string, int, int) {
	xDeduct := minX
	yDeduct := minY

	xEnd := maxX - xDeduct
	yEnd := maxY - yDeduct

	world := make([][]string, yEnd)
	for i := 0; i < yEnd; i++ {
		world[i] = make([]string, xEnd)
		for j := 0; j < xEnd; j++ {
			world[i][j] = "."
		}
	}

	return world, xDeduct, yDeduct
}

func printWorld(world [][]string) {
	fmt.Print("  ")
	for i := 0; i < len(world[0]); i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i, line := range world {
		fmt.Print(i, " ")
		for _, point := range line {
			fmt.Print(point, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func addStructures(world *[][]string, structures [][]Point) {
	for _, structure := range structures {
		//fmt.Println(structure)
		for i := 1; i < len(structure); i++ {
			endPoint := structure[i]
			startPoint := structure[i-1]
			//fmt.Println(startPoint, endPoint)

			lowerPoint := startPoint
			higherPoint := endPoint
			if startPoint.y > endPoint.y {
				lowerPoint = endPoint
				higherPoint = startPoint
			}

			leftPoint := startPoint
			rightPoint := endPoint
			if startPoint.x > endPoint.x {
				leftPoint = endPoint
				rightPoint = startPoint
			}

			if startPoint.x == endPoint.x {
				//fmt.Println("vertical")
				for y := lowerPoint.y; y <= higherPoint.y; y++ {
					(*world)[y][lowerPoint.x] = "#"
				}
			} else {
				//fmt.Println("horizontal")
				for x := leftPoint.x; x <= rightPoint.x; x++ {
					(*world)[leftPoint.y][x] = "#"
				}
			}
			//printWorld(*world)
		}

	}
}

func adjustStructures(structures *[][]Point, deductX int) {
	for i := 0; i < len(*structures); i++ {
		for j := 0; j < len((*structures)[i]); j++ {
			(*structures)[i][j].x -= deductX
		}
	}
}

func runSimulation(world *[][]string, origin Point) int {
	particlesToSimulated := 0
	WORLD := *world

	for {

		if WORLD[origin.y][origin.x] == "o" {
			fmt.Println("BIG GAIN")
			return particlesToSimulated
		}

		particlePlace := origin
		WORLD[origin.y][origin.x] = "o"

		for {
			if WORLD[particlePlace.y+1][particlePlace.x] == "." {
				WORLD[particlePlace.y][particlePlace.x] = "."
				particlePlace.y += 1
				WORLD[particlePlace.y][particlePlace.x] = "o"
			} else if particlePlace.y+1 == len(WORLD) || particlePlace.x-1 < 0 {
				return particlesToSimulated

			} else if WORLD[particlePlace.y+1][particlePlace.x-1] == "." {
				WORLD[particlePlace.y][particlePlace.x] = "."
				particlePlace.y += 1
				particlePlace.x -= 1
				WORLD[particlePlace.y][particlePlace.x] = "o"
			} else if particlePlace.x+1 == len(WORLD[0]) {

				return particlesToSimulated
			} else if WORLD[particlePlace.y+1][particlePlace.x+1] == "." {
				WORLD[particlePlace.y][particlePlace.x] = "."
				particlePlace.y += 1
				particlePlace.x += 1
				WORLD[particlePlace.y][particlePlace.x] = "o"

			} else {
				break
			}
		}

		particlesToSimulated++
		//printWorld(WORLD)
	}
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	structures, maxX, maxY, minX, minY := readStructures(lines)
	//fmt.Println(structures)

	fmt.Println(minX, maxX, minY, maxY)

	world, deductX, deductY := createWorld(minX, maxX+1, 0, maxY+1)
	fmt.Println(deductX, deductY)
	//printWorld(world)

	adjustStructures(&structures, deductX)
	//fmt.Println(deductX, structures)
	addStructures(&world, structures)
	sendOrigin.x -= deductX
	world[sendOrigin.y][sendOrigin.x] = "+"
	//printWorld(world)

	simuCount := runSimulation(&world, sendOrigin)
	//printWorld(world)
	fmt.Println(simuCount)
}
