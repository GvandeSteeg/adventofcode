package main

import (
	"fmt"
	"github.com/GvandeSteeg/adventofcode/util"
	"log"
	"math"
	"strconv"
)

func calculateFuel(mass int) int {
	if i := int(math.Floor(float64(mass)/3)) - 2; i > 0 {
		return i
	} else {
		return 0
	}
}

func calculateTotalFuel(mass int) int {
	var total int
	for i := calculateFuel(mass); i > 0; i = calculateFuel(i) {
		total += i
	}
	return total
}

func main() {
	f, err := util.ReadInputFile(2019, 1)
	if err != nil {
		log.Fatal(err)
	}

	final := 0
	for _, i := range f {
		v, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		final += calculateTotalFuel(v)
	}
	fmt.Println(final)
}
