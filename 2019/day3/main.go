package main

import (
	"errors"
	"fmt"
	"github.com/GvandeSteeg/adventofcode/util"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Wire []Line
type Line [2]Point

// Sort points in line in ascending order
func Swap(line Line) Line {
	if line[0].x > line[1].x {
		return Line{Point{line[1].x, line[0].y}, Point{line[0].x, line[1].y}}
	} else if line[0].y > line[1].y {
		return Line{Point{line[0].x, line[1].y}, Point{line[1].x, line[0].y}}
	} else {
		return line
	}
}

func Intersection(l1, l2 Line) (Point, error) {
	l1 = Swap(l1)
	l2 = Swap(l2)

	if l2[0].x > l1[0].x && l2[0].x < l1[1].x && l1[0].y > l2[0].y && l1[0].y < l2[1].y {
		return Point{l2[0].x, l1[0].y}, nil
	}
	if l1[0].x > l2[0].x && l1[0].x < l2[1].x && l2[0].y > l1[0].y && l2[0].y < l1[1].y {
		return Point{l1[0].x, l2[0].y}, nil
	}

	return Point{}, errors.New("lines do not overlap")
}

func CalcLines(wire string) []Line {
	lines := []Line{}
	lastpoint := Point{0, 0}

	splitwire := strings.Split(wire, ",")
	for _, v := range splitwire {
		direction, strvalue := v[0], v[1:]
		value, err := strconv.Atoi(strvalue)
		if err != nil {
			log.Fatal(err)
		}

		newpoint := Point{lastpoint.x, lastpoint.y}
		switch direction {
		case 'R':
			newpoint.x += value
		case 'L':
			newpoint.x -= value
		case 'U':
			newpoint.y += value
		case 'D':
			newpoint.y -= value
		default:
			log.Fatal("Unexpected direction")
		}
		lines = append(lines, Line{lastpoint, newpoint})
		lastpoint = newpoint
	}

	return lines
}

func findClosest(intersections []Point) int {
	var distances []int
	for _, i := range intersections {
		distances = append(distances, util.Abs(i.x)+util.Abs(i.y))
	}
	_, min := util.Min(distances)
	return min
}

func findIntersections(wire1, wire2 Wire) map[Point]int {
	intersections := make(map[Point]int)
	var w1Cost, w2Cost int
	for _, l1 := range wire1 {
		w1Cost += findSteps(l1)
		for _, l2 := range wire2 {
			w2Cost += findSteps(l2)
			if result, err := Intersection(l1, l2); err == nil {
				intersections[result] = w1Cost - returnChange(l1, result) + w2Cost - returnChange(l2, result)
			}
		}
		w2Cost = 0
	}
	return intersections
}

func findSteps(line Line) int {
	// Lines are always straight, so only X or Y will ever change
	return util.Abs(line[0].x-line[1].x) + util.Abs(line[0].y-line[1].y)
}

// Returns the value to be subtracted from a line's cost once an intersection is encountered
func returnChange(line Line, intersection Point) int {
	xcost := util.Abs(line[0].x - intersection.x)
	xlen := util.Abs(line[0].x - line[1].x)
	xchange := xlen - xcost
	ycost := util.Abs(line[0].y - intersection.y)
	ylen := util.Abs(line[0].y - line[1].y)
	ychange := ylen - ycost

	return xchange + ychange
}

func PrintLines(wire1, wire2 Wire) {
	zip := func(a, b Wire) Wire {
		if len(a) > len(b) {
			b = append(b, Line{})
		} else if len(a) < len(b) {
			a = append(a, Line{})
		}

		r := Wire{}

		for i, e := range a {
			r = append(r, Line{e[0], b[i][0]})
			r = append(r, Line{e[1], b[i][1]})
		}

		return r
	}

	zipped := zip(wire1, wire2)
	for _, a := range zipped {
		fmt.Printf("%d,%d,%d,%d\n", a[0].x, a[0].y, a[1].x, a[1].y)
		time.Sleep(10 * time.Microsecond) // Had to add a limit otherwise a chunk of data is lost
	}
	os.Exit(1)
}

func part1(intersections map[Point]int) int {
	var transform []Point
	for i := range intersections {
		transform = append(transform, i)
	}
	return findClosest(transform)
}

func part2(intersections map[Point]int) int {
	var costs []int
	for _, cost := range intersections {
		costs = append(costs, cost)
	}
	_, min := util.Min(costs)
	return min
}

func main() {
	f, err := util.ReadInputFile(2019, 3)
	if err != nil {
		log.Fatal(err)
	}

	wire1 := CalcLines(f[0])
	wire2 := CalcLines(f[1])
	// PrintLines(wire1, wire2)
	intersections := findIntersections(wire1, wire2)
	fmt.Println("Part 1 - Manhattan Distance to closest intersection:", part1(intersections))
	fmt.Println("Part 2 - Minimal amount of steps required for closest intersection:", part2(intersections))
}
