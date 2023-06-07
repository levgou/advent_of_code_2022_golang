package day_12

import (
	"advent_of_code/pkg/shared"
	"fmt"
	graph2 "gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"log"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_12/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_12/demo_input.txt"
)

func genStepsMap(matrix [][]rune) [][]string {
	stepsMap := make([][]string, len(matrix))
	for i, line := range matrix {
		stepsMap[i] = make([]string, len(line))
		for j, _ := range line {
			stepsMap[i][j] = "."
		}
	}

	return stepsMap
}

type Point struct {
	row int
	col int
}

func genHeightMatrix(lines []string) ([][]rune, Point, Point, []Point) {
	end := Point{}
	start := Point{}
	aPoints := make([]Point, 0)

	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = make([]rune, len(line))
		for j, char := range line {
			matrix[i][j] = char
			if char == 'E' {
				end = Point{i, j}
			} else if char == 'S' {
				start = Point{i, j}
			} else if char == 'a' {
				aPoints = append(aPoints, Point{i, j})
			}
		}
	}

	return matrix, start, end, aPoints
}

func printMatrix[T string | rune](matrix [][]T) {
	for _, line := range matrix {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func genGraph(matrix [][]rune) (*simple.DirectedGraph, map[graph2.Node]Point, map[Point]graph2.Node) {
	nodeToCoords := make(map[graph2.Node]Point)
	coordsToNode := make(map[Point]graph2.Node)

	inMatrixAndReachable := func(p, other Point) bool {
		inside := other.row >= 0 && other.row < len(matrix) &&
			other.col >= 0 && other.col < len(matrix[0])

		if !inside {
			return false
		}

		reachable := matrix[other.row][other.col]-matrix[p.row][p.col] <= 1
		return reachable
	}

	graph := simple.NewDirectedGraph()
	for row, line := range matrix {
		for col, _ := range line {
			node := graph.NewNode()
			graph.AddNode(node)
			nodeToCoords[node] = Point{row, col}
			coordsToNode[Point{row, col}] = node
		}
	}

	for row, line := range matrix {
		for col, _ := range line {
			cur := Point{row, col}

			top := Point{row - 1, col}
			bottom := Point{row + 1, col}
			left := Point{row, col - 1}
			right := Point{row, col + 1}

			if inMatrixAndReachable(cur, top) {
				graph.SetEdge(simple.Edge{F: coordsToNode[cur], T: coordsToNode[top]})
			}
			if inMatrixAndReachable(cur, bottom) {
				graph.SetEdge(simple.Edge{F: coordsToNode[cur], T: coordsToNode[bottom]})
			}
			if inMatrixAndReachable(cur, left) {
				graph.SetEdge(simple.Edge{F: coordsToNode[cur], T: coordsToNode[left]})
			}
			if inMatrixAndReachable(cur, right) {
				graph.SetEdge(simple.Edge{F: coordsToNode[cur], T: coordsToNode[right]})
			}
		}
	}

	return graph, nodeToCoords, coordsToNode
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	matrix, start, end, aPoints := genHeightMatrix(lines)
	matrix[start.row][start.col] = 'a'
	matrix[end.row][end.col] = 'z'

	fmt.Println("Start: ", start, "End: ", end)

	graph, nodeToCoords, coordsToNode := genGraph(matrix)
	shortest := path.DijkstraAllPaths(graph)
	p, weight, _ := shortest.Between(coordsToNode[start].ID(), coordsToNode[end].ID())

	fmt.Println(weight)

	for _, node := range p {
		fmt.Println(nodeToCoords[node])
	}

	minWeight := weight
	for _, aPoint := range aPoints {
		_, weight, _ := shortest.Between(coordsToNode[aPoint].ID(), coordsToNode[end].ID())
		if weight < minWeight {
			minWeight = weight
		}
	}

	fmt.Println(minWeight)
}
