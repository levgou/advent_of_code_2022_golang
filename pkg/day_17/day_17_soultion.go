package day_17

import (
	"advent_of_code/pkg/shared"
	"fmt"
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
		//'🟥',
		'#',
	},
	{
		[][]rune{
			{' ', '@', ' '},
			{'@', '@', '@'},
			{' ', '@', ' '},
		},
		//'🟨',
		'#',
	},
	{
		[][]rune{
			{'@', '@', '@'},
			{'@', ' ', ' '},
			{'@', ' ', ' '},
		},
		//'🟩',

		'#',
	},
	{
		[][]rune{
			{'@'},
			{'@'},
			{'@'},
			{'@'},
		},
		//'🟦',
		'#',
	},
	{
		[][]rune{
			{'@', '@'},
			{'@', '@'},
		},
		//'🟪',
		'#',
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

func trimWorld(level int) {
	World = World[level:]
}

func isRock(char rune) bool {
	return char == '#'

	for i := 0; i < len(Rocks); i++ {
		if char == Rocks[i] {
			return true
		}
	}
	return false

	//return slices.Index(Rocks, char) != -1
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

func findClosedLevel() int {
	found := false

	for row := len(World) - 1; row > 0; row-- {
		found = true
		for col := 0; col < len(World[row]); col++ {
			if !isRock(World[row][col]) {
				//&& !isRock(World[row-1][col])
				//{
				found = false
				break
			}
		}

		if found {
			return row - 1
		}
	}

	return -1
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

type Memo struct {
	shape    int
	jet      int
	horizPos int
}

type MemoVal struct {
	idx    int
	height int
}

func run(jets string, rockCount int) int {
	curShape := 0
	curJet := 0
	highest := -1
	linesRemoved := 0
	memo := make(map[Memo]MemoVal)

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

				// todo: this was an attempt to save on memory
				//blockLevel := findClosedLevel()
				//if blockLevel > 5 {
				//printWorld()
				//println("block level", blockLevel)
				//trimWorld(blockLevel)
				//linesRemoved += blockLevel
				//printWorld()
				//}

				highest = findHighest()
				//fmt.Println("highest", highest)

				// We are detection a cycle
				// if we find that cur shape finished in the same place horiz
				// and the jet idx is the same, we check how many rocks passed since
				// in case what is left is a multiple of this, we calc the cycle height gain
				// and multiply it by (what's left) / (cycle length)
				stt := Memo{curShape, curJet, afterDown.x}
				if val, ok := memo[stt]; ok {
					idxDiff := i - val.idx
					heightGain := highest - val.height
					rocksRemain := rockCount - i - 1
					if rocksRemain%idxDiff == 0 {
						fmt.Println("Rocks left are a mult of memo H", i, idxDiff, rocksRemain, heightGain, rocksRemain/idxDiff)
						return highest + (rocksRemain/idxDiff)*heightGain + 1
					}
				} else {
					memo[stt] = MemoVal{i, highest}
				}

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

	return linesRemoved + highest + 1
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
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	jets := lines[0]

	//1_000_000_000_000

	println(run(jets, 1_000_000_000_000))
	//printWorld()
}
