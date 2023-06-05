package day_05

import (
	"advent_of_code/pkg/shared"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_05/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_05/input.txt"
)

func columnNumLineNum(lines []string) (int, error) {
	for i, line := range lines {
		if len(line) > 0 && strings.Trim(line, " ")[0] == '1' {
			return i, nil
		}
	}

	return 0, errors.New("Line not found")
}

func columnNum(lines []string) (int64, error) {
	columnLineIndex, err := columnNumLineNum(lines)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	columnNumsLine := lines[columnLineIndex]
	colNums := strings.Split(strings.Trim(columnNumsLine, " "), " ")
	return strconv.ParseInt(colNums[len(colNums)-1], 10, 64)

	return 0, errors.New("No column found")
}

func indexForCol(col int) int {
	return 1 + 4*(col-1)
}

func colCrates(lines []string, colNum int) (map[string][]string, error) {
	res := make(map[string][]string)
	columns := make([]string, colNum)
	for i := 0; i < colNum; i++ {
		columns[i] = strconv.Itoa(i + 1)
		res[columns[i]] = []string{}
	}

	for i := len(lines) - 1; i >= 0; i-- {
		curLine := lines[i] + strings.Repeat(" ", 100)
		for i, col := range columns {
			index := indexForCol(i + 1)
			colChar := string(curLine[index])
			if colChar != " " {
				res[col] = append(res[col], colChar)
			}
		}
	}

	return res, nil
}

type Move struct {
	Amount  int
	FromCol string
	ToCol   string
}

func parseMoveLines(lines []string) []Move {
	moves := make([]Move, len(lines))
	replacer := strings.NewReplacer("move ", "", "from ", "", "to ", "")

	for i, line := range lines {
		numbersLine := replacer.Replace(line)
		numbers := strings.Split(numbersLine, " ")
		amount, _ := strconv.ParseInt(string(numbers[0]), 10, 64)
		moves[i] = Move{
			Amount:  int(amount),
			FromCol: string(numbers[1]),
			ToCol:   string(numbers[2]),
		}
	}

	return moves
}

func performMoves(cratesMap map[string][]string, moves []Move, pickMoreThanOne bool) map[string][]string {
	curMoveMap := make(map[string][]string)
	for k, v := range cratesMap {
		curMoveMap[k] = make([]string, len(v))
		copy(curMoveMap[k], v)
	}

	for _, move := range moves {
		fmt.Println(move)
		moveIndex := len(curMoveMap[move.FromCol]) - move.Amount
		cratesToMove := curMoveMap[move.FromCol][moveIndex:]
		curMoveMap[move.FromCol] = curMoveMap[move.FromCol][:moveIndex]

		if pickMoreThanOne {
			curMoveMap[move.ToCol] = append(curMoveMap[move.ToCol], cratesToMove...)
		} else {
			curMoveMap[move.ToCol] = append(curMoveMap[move.ToCol], shared.Reverse(cratesToMove)...)
		}

		fmt.Println(curMoveMap)
	}

	return curMoveMap

}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	columnLineIndex, err := columnNumLineNum(lines)
	if err != nil {
		log.Fatal(err)
		return
	}

	colNum, err := columnNum(lines)
	if err != nil {
		log.Fatal(err)
		return
	}

	stackLines := lines[:columnLineIndex]

	cratesMap, err := colCrates(stackLines, int(colNum))
	if err != nil {
		log.Fatal(err)
		return
	}

	moveLines := lines[columnLineIndex+2:]
	moves := parseMoveLines(moveLines)

	afterMoveMap := performMoves(cratesMap, moves, true)

	topLine := make([]string, colNum)
	for i := 0; i < len(topLine); i++ {
		col := strconv.Itoa(i + 1)
		topLine[i] = shared.Last(afterMoveMap[col])
	}

	fmt.Println(topLine, strings.Join(topLine, ""))
}
