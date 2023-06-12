package day_17

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_17/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_17/demo_input.txt"
)

var Rocks = []rune{'🟥', '🟨', '🟩', '🟦', '🟪'}

const (
	//Rock       = '#'
	MovingRock = '@'
	Air        = ' '
)

type Shape struct {
	repr  [][]rune
	color rune
}

var Shapes = []Shape{
	{
		[][]rune{{'@', '@', '@', '@'}},
		'🟥',
	},
	{
		[][]rune{
			{' ', '@', ' '},
			{'@', '@', '@'},
			{' ', '@', ' '},
		},
		'🟨',
	},
	{
		[][]rune{
			{'@', '@', '@'},
			{'@', ' ', ' '},
			{'@', ' ', ' '},
		},
		'🟩',
	},
	{
		[][]rune{
			{'@'},
			{'@'},
			{'@'},
			{'@'},
		},
		'🟦',
	},
	{
		[][]rune{
			{'@', '@'},
			{'@', '@'},
		},
		'🟪',
	},
}

const Right = '>'
const Left = '<'

const StartLeftDist = 2
const StartBottomDist = 3 + 1

const Width = 7

var World = [][Width]rune{
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
	{' ', ' ', ' ', ' ', ' ', ' ', ' '},
}

var Line = [Width]rune{' ', ' ', ' ', ' ', ' ', ' ', ' '}

type Point struct {
	x int
	y int
}

func isRock(char rune) bool {
	return slices.Index(Rocks, char) != -1
}

func findHighest() int {
	for row := len(World) - 1; row >= 0; row-- {
		for col := 0; col < len(World[row]); col++ {
			if isRock(World[row][col]) {
				return row
			}
		}
	}

	return 0
}

func removeShape(shape [][]rune, topLeft Point) {
	for i, row := range shape {
		for j, _ := range row {
			if World[topLeft.y+i][topLeft.x+j] == MovingRock {
				World[topLeft.y+i][topLeft.x+j] = Air
			}
		}
	}
}

func paintShapeStatic(shape Shape, topLeft Point) {
	for i, row := range shape.repr {
		for j, char := range row {
			if char == MovingRock {
				char = shape.color
			}
			if char != Air {
				World[topLeft.y+i][topLeft.x+j] = char
			}
		}
	}
}

func paintShape(shape [][]rune, topLeft Point) {
	for i, row := range shape {
		for j, char := range row {
			if char != Air {
				World[topLeft.y+i][topLeft.x+j] = char
			}
		}
	}
}

func shapeTouchesOther(shape [][]rune, topLeft Point) bool {
	for i, row := range shape {
		for j, char := range row {
			if char == MovingRock && isRock(World[topLeft.y+i][topLeft.x+j]) {
				return true
			}
		}
	}
	return false
}

func moveShapeDown(shape [][]rune, topLeft Point) Point {
	afterMove := topLeft
	if afterMove.y > 0 {
		afterMove.y--
	}

	if shapeTouchesOther(shape, afterMove) {
		return topLeft
	}

	return afterMove
}

func moveShapeHoriz(shape [][]rune, topLeft Point, direction rune) Point {
	shapeWidth := len(shape[0])
	afterMove := topLeft

	if direction == Right && afterMove.x > 0 {
		//fmt.Println("moving right")
		afterMove.x--

	} else if direction == Left && afterMove.x+shapeWidth < Width {
		//fmt.Println("moving left")
		afterMove.x++
	}

	if shapeTouchesOther(shape, afterMove) {
		//fmt.Println("touching other")
		return topLeft
	}

	return afterMove
}

func run(jets string, rockCount int) {
	curShape := 0
	curJet := 0
	highest := -1

	for i := 0; i < rockCount; i++ {
		shape := Shapes[curShape]

		//shapeHeight := len(shape)
		shapeWidth := len(shape.repr[0])

		shapeTopLeft := Point{
			Width - shapeWidth - StartLeftDist,
			highest + StartBottomDist,
		}

		//println("\n---- NEW")
		paintShape(shape.repr, shapeTopLeft)
		//printWorld()

		for {
			jet := jets[curJet]
			//fmt.Printf("jet >>>>>>>>> [%d] [%c] \n", curJet, rune(jet))

			afterHoriz := moveShapeHoriz(shape.repr, shapeTopLeft, rune(jet))
			curJet = (curJet + 1) % len(jets)
			afterDown := moveShapeDown(shape.repr, afterHoriz)

			removeShape(shape.repr, shapeTopLeft)
			paintShape(shape.repr, afterDown)

			if afterDown == afterHoriz {
				paintShapeStatic(shape, afterDown)
				highest = findHighest()
				//fmt.Println("highest", highest)
				break

			} else {
				shapeTopLeft = afterDown
			}
		}
		if len(World)-highest < 8 {
			for j := 0; j < 3; j++ {
				World = append(World, Line)
			}
		}
		//printWorld()
		curShape = (curShape + 1) % len(Shapes)
	}
}

func printWorld() {

	for i := len(World) - 1; i >= 0; i-- {
		fmt.Printf("%05d| ", i)
		for j := Width - 1; j >= 0; j-- {
			char := World[i][j]
			if char == ' ' {
				char = '.'
			} else if char == MovingRock {
				char = '💠'
			}
			print(string(char))
			if char == '.' {
				print(" ")
			}
		}
		println("|")
	}

	println("     + - - - - - - - +\n")

	return

	for i, row := range World {
		fmt.Printf("%05d| ", i)
		for _, char := range row {
			if char == ' ' {
				char = '.'
			} else if char == MovingRock {
				char = '💠'
			}
			print(string(char))
			if char == '.' {
				print(" ")
			}
		}
		println("|")
	}
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	jets := lines[0]
	run(jets, 2022)

	//printWorld()
}
