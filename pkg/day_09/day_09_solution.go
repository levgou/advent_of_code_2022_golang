package day_09

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_09/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_09/input.txt"
)

const Visted = "[#]"
const NotVisited = "[.]"
const Start = "s"
const Head = "H"
const Tail = "T"

const WorldDim = 2000

var World = [WorldDim][WorldDim]string{}

func resetWorldNumbers() {
	for row := 0; row < len(World); row++ {
		for col := 0; col < len(World[0]); col++ {
			if World[row][col] != Visted {
				World[row][col] = NotVisited
			}
		}
	}

}

type Point struct {
	Row int
	Col int
}

const (
	UP    = "U"
	DOWN  = "D"
	LEFT  = "L"
	RIGHT = "R"
)

type MoveSeq struct {
	Direction string
	Steps     int
}

var StartPoint = Point{Row: WorldDim / 2, Col: WorldDim / 2}

func printWorld() {
	for row := 0; row < len(World); row++ {
		for col := 0; col < len(World[0]); col++ {
			fmt.Print(World[row][col])
		}
		fmt.Println()
	}
}

func parseInput(lines []string) []MoveSeq {
	directions := make([]MoveSeq, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		direction := MoveSeq{
			Direction: parts[0],
			Steps:     shared.ParseInt(parts[1]),
		}
		directions[i] = direction
	}

	return directions
}

func movePoint(direction string, headPoint Point) Point {
	switch direction {
	case UP:
		headPoint.Row--
	case DOWN:
		headPoint.Row++
	case LEFT:
		headPoint.Col--
	case RIGHT:
		headPoint.Col++
	}

	return headPoint
}

func pointsTouching(p1, p2 Point) bool {
	// same place
	if p1.Row == p2.Row && p1.Col == p2.Col {
		return true
	}

	// same row and adjacent columns
	if p1.Row == p2.Row && (p1.Col == p2.Col+1 || p1.Col == p2.Col-1) {
		return true
	}

	// same column and adjacent rows
	if p1.Col == p2.Col && (p1.Row == p2.Row+1 || p1.Row == p2.Row-1) {
		return true
	}

	// diagonally adjacent (exactly 1 diff in row and 1 diff in col)
	if (p1.Row == p2.Row+1 || p1.Row == p2.Row-1) && (p1.Col == p2.Col+1 || p1.Col == p2.Col-1) {
		return true
	}

	return false
}

func moveTail(tail, head Point) Point {
	if pointsTouching(tail, head) {
		return tail
	}

	if tail.Row == head.Row {
		if tail.Col < head.Col {
			tail.Col++
		} else {
			tail.Col--
		}
	} else if tail.Col == head.Col {
		if tail.Row < head.Row {
			tail.Row++
		} else {
			tail.Row--
		}
	} else {
		// the distance is 2 in some direction
		/*
			.....    .....    .....
			.....    ..H..    ..H..
			..H.. -> ..... -> ..T..
			.T...    .T...    .....
			.....    .....    .....

			.....    .....    .....
			.....    .....    .....
			..H.. -> ...H. -> ..TH.
			.T...    .T...    .....
			.....    .....    .....
		*/
		// the distance is 2 in diagonal direction
		if int(math.Abs(float64(head.Row-tail.Row))) == 2 &&
			int(math.Abs(float64(head.Col-tail.Col))) == 2 {
			tail.Row += (head.Row - tail.Row) / 2
			tail.Col += (head.Col - tail.Col) / 2
		} else if head.Row-tail.Row == 2 {
			tail.Row++
			tail.Col = head.Col
		} else if head.Row-tail.Row == -2 {
			tail.Row--
			tail.Col = head.Col
		} else if head.Col-tail.Col == 2 {
			tail.Col++
			tail.Row = head.Row
		} else if head.Col-tail.Col == -2 {
			tail.Col--
			tail.Row = head.Row
		}
	}

	return tail
}

func runSteps(steps []MoveSeq) {
	var headPoint = StartPoint
	var tailPoint = StartPoint

	for _, stepSeq := range steps {
		direction := stepSeq.Direction
		for i := 0; i < stepSeq.Steps; i++ {
			headPoint = movePoint(direction, headPoint)
			tailPoint = moveTail(tailPoint, headPoint)
			World[tailPoint.Row][tailPoint.Col] = Visted
		}

	}
}
func runSteps10(steps []MoveSeq) {
	viseted := mapset.NewSet[Point]()
	// todo: here should init 10 points insted of 2
	points := make([]Point, 10)
	for i := 0; i < 10; i++ {
		points[i] = StartPoint
	}

	for _, stepSeq := range steps {
		direction := stepSeq.Direction
		for i := 0; i < stepSeq.Steps; i++ {
			points[0] = movePoint(direction, points[0])
			World[points[0].Row][points[0].Col] = "[0]"
			for j := 1; j < 10; j++ {
				points[j] = moveTail(points[j], points[j-1])
				World[points[j].Row][points[j].Col] = "[" + strconv.Itoa(j) + "]"
			}
			viseted.Add(points[9])
		}
	}
	fmt.Println("Visited", viseted.Cardinality())
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	for row := 0; row < len(World); row++ {
		for col := 0; col < len(World[0]); col++ {
			World[row][col] = NotVisited
		}
	}

	World[StartPoint.Row][StartPoint.Col] = Head

	steps := parseInput(lines)
	runSteps(steps)
	countVisited := 0
	for row := 0; row < len(World); row++ {
		for col := 0; col < len(World[0]); col++ {
			if World[row][col] == Visted {
				countVisited++
			}
		}
	}

	fmt.Println("Visited", countVisited)
}

func Solution2() {

	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	for row := 0; row < len(World); row++ {
		for col := 0; col < len(World[0]); col++ {
			World[row][col] = NotVisited
		}
	}

	World[StartPoint.Row][StartPoint.Col] = Head

	steps := parseInput(lines)
	runSteps10(steps)
}
