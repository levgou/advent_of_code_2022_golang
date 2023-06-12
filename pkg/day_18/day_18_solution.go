package day_18

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/phf/go-queue/queue"
	"log"
	"strings"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_18/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_18/demo_input.txt"
)

type Particle struct {
	x int
	y int
	z int
}

func parse(lines []string) []Particle {
	ps := make([]Particle, 0)
	for _, line := range lines {
		parts := strings.Split(line, ",")
		ps = append(ps, Particle{
			x: shared.ParseInt(parts[0]),
			y: shared.ParseInt(parts[1]),
			z: shared.ParseInt(parts[2]),
		})
	}

	return ps
}

func checkNonTouchingSide(ps []Particle) int {
	set := mapset.NewSet[Particle]()
	for _, p := range ps {
		set.Add(p)
	}

	count := 0
	for _, p := range ps {
		openSides := 6
		neighbours := getNeighbours(p)
		for _, n := range neighbours {
			if set.Contains(n) {
				openSides--
			}
		}
		count += openSides
	}

	return count
}

const directionTries = 10

func tryReach(src Particle, target Particle, set mapset.Set[Particle]) bool {
	visited := mapset.NewSet[Particle]()
	q := queue.New()
	q.PushBack(src)

	for q.Len() > 0 {
		cur := q.PopFront()
		curP := cur.(Particle)
		visited.Add(curP)

		if cur == target {
			return true
		}

		for _, n := range getNeighbours(curP) {
			if !visited.Contains(n) && !set.Contains(n) {
				q.PushBack(n)
			}
		}

	}

	return false
}

func checkNonTouchingSideNoPockets(ps []Particle) int {
	set := mapset.NewSet[Particle]()
	for _, p := range ps {
		set.Add(p)
	}

	flooded := floodAir(set)

	count := 0
	for _, p := range ps {
		openSides := 6
		neighbours := getNeighbours(p)
		for _, n := range neighbours {
			if set.Contains(n) {
				openSides--
			} else if !flooded.Contains(n) {
				openSides--
			}
		}
		count += openSides
	}

	return count
}

func floodAir(set mapset.Set[Particle]) mapset.Set[Particle] {
	lowerLimit := -1
	src := Particle{lowerLimit, lowerLimit, lowerLimit}
	maxX := 0
	maxY := 0
	maxZ := 0

	for _, p := range set.ToSlice() {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.z > maxZ {
			maxZ = p.z
		}
	}

	// to flood can "go around max"
	maxX += 2
	maxY += 2
	maxZ += 2

	fmt.Println("flooding upto", maxX, maxY, maxZ)

	visited := mapset.NewSet[Particle]()
	q := queue.New()
	q.PushBack(src)

	for q.Len() > 0 {
		cur := q.PopFront()
		curP := cur.(Particle)
		if visited.Contains(curP) {
			continue
		}
		visited.Add(curP)

		for _, n := range getNeighbours(curP) {
			if !visited.Contains(n) && !set.Contains(n) &&
				n.x <= maxX && n.y <= maxY && n.z <= maxZ &&
				n.x >= lowerLimit && n.y >= lowerLimit && n.z >= lowerLimit {

				q.PushBack(n)
			}
		}

	}

	return visited
}

func getNeighbours(p Particle) []Particle {
	return []Particle{
		{p.x + 1, p.y, p.z},
		{p.x - 1, p.y, p.z},
		{p.x, p.y + 1, p.z},
		{p.x, p.y - 1, p.z},
		{p.x, p.y, p.z + 1},
		{p.x, p.y, p.z - 1},
	}
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	//lines = []string{
	//	"1,1,1",
	//	"2,1,1",
	//}

	ps := parse(lines)
	count := checkNonTouchingSide(ps)
	count2 := checkNonTouchingSideNoPockets(ps)
	fmt.Println(count, count2)
}
