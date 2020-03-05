package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"strings"
)

// https://adventofcode.com/2019/day/2
func executeIntCode(x []int, y []int) (bool, error) {
	switch x[0] {
	default:
		return false, fmt.Errorf("Invalid opcode in position 0: %d", x[0])
	case 1:
		y[x[3]] = y[x[1]] + y[x[2]]
	case 2:
		y[x[3]] = y[x[1]] * y[x[2]]
	case 99:
		return true, nil
	}
	return false, nil
}

func part1(sliceint []int) int {
	// Initialise first values
	sliceint[1] = 12
	sliceint[2] = 2

	var done bool
	var err error

	for i := 0; i < len(sliceint); i = i + 4 {
		done, err = executeIntCode(sliceint[i:i+4], sliceint)
		if err != nil {
			log.Fatal(err)
		}
		if done {
			break
		}
	}
	return sliceint[0]
}

func part2(sliceint []int) int {
	var done bool
	var err error

	tmp := make([]int, len(sliceint))

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(tmp, sliceint)
			tmp[1] = noun
			tmp[2] = verb
			for i := 0; i < len(tmp); i = i + 4 {
				done, err = executeIntCode(tmp[i:i+4], tmp)
				if err != nil {
					log.Fatal(err)
				}
				if done {
					break
				}
			}
			if tmp[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	log.Fatal("No expected output returned")
	return 0
}

func main() {
	f, err := util.ReadInputFile(2019, 2)
	if err != nil {
		log.Fatal(err)
	}

	slicestr := strings.Split(f[0], ",")
	sliceint, _ := util.ConvertSliceStrtoSliceInt(slicestr)
	fmt.Println("Part 1:", part1(sliceint))

	sliceint, _ = util.ConvertSliceStrtoSliceInt(slicestr)
	fmt.Println("Part 2:", part2(sliceint))
}
