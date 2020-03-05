package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInputFile(year, day int) ([]string, error) {
	f, err := os.Open(fmt.Sprintf("/Users/gvandestee/GolandProjects/src/github.com/GvandeSteeg/adventofcode/%d/day%d/input.txt", year, day))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	s := bufio.NewScanner(f)
	var j []string
	for s.Scan() {
		j = append(j, s.Text())
	}

	err = s.Err()
	if err != nil {
		return nil, err
	}

	return j, nil
}

func ConvertSliceStrtoSliceInt(slice []string) ([]int, error) {
	var intslice []int
	for _, v := range slice {
		strv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intslice = append(intslice, strv)
	}
	return intslice, nil
}

func Abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func Min(a []int) (index, min int) {
	min = a[0]
	index = 0
	for save, i := range a {
		if min > i {
			min = i
			index = save
		}
	}
	return index, min
}
