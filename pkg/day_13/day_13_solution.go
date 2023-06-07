package day_13

import (
	"advent_of_code/pkg/shared"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"reflect"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_13/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_13/demo_input.txt"
)

const (
	OK   = "OK"
	FAIL = "FAIL"
	CONT = "CONT"
)

func compareArrays(leftArr, rightArr []interface{}) string {
	if len(leftArr) == 0 && len(rightArr) == 0 {
		return CONT
	} else if len(leftArr) == 0 {
		return OK
	} else if len(rightArr) == 0 {
		return FAIL
	}

	firstLeft := leftArr[0]
	firstRight := rightArr[0]

	firstComp := compareMsgs(firstLeft, firstRight)
	if firstComp == OK || firstComp == FAIL {
		return firstComp
	}

	return compareArrays(leftArr[1:], rightArr[1:])
}

func compareMsgs(left, right interface{}) string {

	leftArr, ok1 := left.([]interface{})
	rightArr, ok2 := right.([]interface{})

	if ok1 && ok2 {
		return compareArrays(leftArr, rightArr)
	}

	if ok1 {
		return compareArrays(leftArr, []interface{}{right})
	} else if ok2 {
		return compareArrays([]interface{}{left}, rightArr)
	}

	leftInt, ok1 := left.(float64)
	rightInt, ok2 := right.(float64)

	if ok1 && ok2 {
		if leftInt < rightInt {
			return OK
		} else if leftInt == rightInt {
			return CONT
		} else {
			return FAIL
		}
	}

	fmt.Println("left", left, reflect.TypeOf(left))
	fmt.Println("right", right, reflect.TypeOf(right))
	panic("oi vey voy")
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	grouped := shared.GroupLines(lines)
	//fmt.Println(grouped)

	okIndeces := 0
	for i, pair := range grouped {
		left, right := pair[0], pair[1]
		var leftParsed, rightParsed interface{}
		json.Unmarshal([]byte(left), &leftParsed)
		json.Unmarshal([]byte(right), &rightParsed)

		if i+1 == 127 {
			fmt.Println("left", left)
			fmt.Println("right", right)
		}

		res := compareMsgs(leftParsed, rightParsed)
		fmt.Println(i+1, res)
		if res == OK {
			okIndeces += i + 1
		}
	}

	fmt.Println(okIndeces)
}

func Solution2() {

	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	grouped := shared.GroupLines(lines)
	grouped = append(grouped, []string{"[[2]]", "[[6]]"})

	var allUnmarshalled []interface{}
	for _, pair := range grouped {
		left, right := pair[0], pair[1]
		var leftParsed, rightParsed interface{}
		json.Unmarshal([]byte(left), &leftParsed)
		json.Unmarshal([]byte(right), &rightParsed)
		allUnmarshalled = append(allUnmarshalled, leftParsed, rightParsed)
	}

	slices.SortFunc(allUnmarshalled, func(a, b interface{}) bool {
		return compareMsgs(a, b) == OK
	})

	mult := 1

	for i, msg := range allUnmarshalled {
		fmt.Println(i+1, msg)
		if fmt.Sprintf("%v", msg) == "[[6]]" || fmt.Sprintf("%v", msg) == "[[2]]" {
			fmt.Println("found", i+1)
			mult *= i + 1
		}
	}

	fmt.Println(mult)
}
