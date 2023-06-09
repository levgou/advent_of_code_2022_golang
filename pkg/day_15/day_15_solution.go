package day_15

import (
	"advent_of_code/pkg/shared"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/slices"
	"log"
	"math"
	"math/big"
	"strings"
	"time"
)

const (
	Input = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_15/input.txt"
	Demo  = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_15/demo_input.txt"
)

type Point struct {
	x int
	y int
}

type Beacon struct {
	x int
	y int
}

type Sensor struct {
	x                 int
	y                 int
	beacon            Beacon
	manhattanDistance int
}

func parse(lines []string) []Sensor {
	// parse rows that look like:
	// Sensor at x=20, y=1: closest beacon is at x=15, y=3
	//	0	  1	  2	    3	   4	  5	    6  7  8	    9

	sensors := make([]Sensor, 0)
	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '=' || r == ',' || r == ':' || r == ' '
		})

		x := shared.ParseInt(parts[3])
		y := shared.ParseInt(parts[5])

		beaconX := shared.ParseInt(parts[11])
		beaconY := shared.ParseInt(parts[13])

		manhattanDistance := int(math.Abs(float64(beaconX-x)) + math.Abs(float64(beaconY-y)))

		sensor := Sensor{x, y, Beacon{beaconX, beaconY}, manhattanDistance}
		sensors = append(sensors, sensor)
	}

	return sensors
}

func allPointsWithinManhattanDistance(sensor Sensor, onLine int) mapset.Set[Point] {
	points := mapset.NewSet[Point]()
	distance := sensor.manhattanDistance

	if sensor.y-distance > onLine ||
		sensor.y+distance < onLine {
		return points
	}

	if sensor.y < onLine {
		distanceToLine := onLine - sensor.y
		xDistance := distance - distanceToLine
		for addX := 0; addX <= xDistance; addX++ {
			points.Add(Point{sensor.x + addX, onLine})
			points.Add(Point{sensor.x - addX, onLine})
		}

	} else if sensor.y > onLine {
		distanceToLine := sensor.y - onLine
		xDistance := distance - distanceToLine
		for addX := 0; addX <= xDistance; addX++ {
			points.Add(Point{sensor.x + addX, onLine})
			points.Add(Point{sensor.x - addX, onLine})
		}

	} else {
		points.Add(Point{sensor.x + distance, onLine})
		points.Add(Point{sensor.x - distance, onLine})
	}

	return points
}

func allPointsNoBeaconExists(sensors []Sensor, onLine int) mapset.Set[Point] {
	points := mapset.NewSet[Point]()

	for _, sensor := range sensors {
		newPoints := allPointsWithinManhattanDistance(sensor, onLine)
		points = points.Union(newPoints)
	}

	for _, sensor := range sensors {
		points.Remove(Point{sensor.beacon.x, sensor.beacon.y})
		points.Remove(Point{sensor.x, sensor.y})
	}

	return points
}

func noBeaconOnLineYCount(noBeacon mapset.Set[Point], y int) int {
	count := 0
	for point := range noBeacon.Iter() {
		if point.y == y {
			count++
		}
	}

	return count
}

func manhattanDistance(sensor Sensor, beacon Beacon) int {
	return int(math.Abs(float64(beacon.x-sensor.x)) + math.Abs(float64(beacon.y-sensor.y)))
}

func findBeaconUpTo(sensors []Sensor, maxXYbeacon int) (int, int) {
	slices.SortFunc(sensors, func(i, j Sensor) bool {
		return i.y < j.y
	})

	for y := 0; y < maxXYbeacon; y++ {

		x := 0
		for x < maxXYbeacon {

			maxDistOverhead := -1
			s := Sensor{}
			for _, sensor := range sensors {
				dist := manhattanDistance(sensor, Beacon{x, y})
				distanceOverhead := sensor.manhattanDistance - dist

				if dist <= sensor.manhattanDistance &&
					distanceOverhead > maxDistOverhead {
					maxDistOverhead = distanceOverhead
					s = sensor
				}
			}

			if maxDistOverhead == -1 {
				fmt.Println("found it", x, y)
				return x, y
			} else {
				xDiff := s.x - x
				// tested point is to the left of the sensor
				if xDiff > 0 {
					x += xDiff
					x += maxDistOverhead
				} else {
					x += maxDistOverhead
				}
			}

			x++
		}
	}

	return -1, -1
}

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	maxXYbeacon := 4_000_000
	//maxXYbeacon = 20

	sensors := parse(lines)

	//fmt.Println("sensors", sensors)

	//line := 2000000
	//noBeacon := allPointsNoBeaconExists(sensors, line)
	//noCount := noBeaconOnLineYCount(noBeacon, line)
	//fmt.Println("noCount", noCount)

	start := time.Now()

	x, y := findBeaconUpTo(sensors, maxXYbeacon)
	fmt.Println("Ans:", x*4000000+y)

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)
}
