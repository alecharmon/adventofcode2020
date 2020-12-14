package day9

import (
	"log"
	"math"
	"sort"
	"strconv"

	"github.com/alecharmon/adventofcode2020"
)

func problem1(fp string, size int) int {
	input := []int{}
	for line := range adventofcode2020.FileIterator(fp) {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, v)
	}

	lookup := make(map[int]bool)
	window, input := input[:size], input[size:]
	//starting off the lookup table
	for _, v := range window {
		lookup[v] = true
	}

	for _, v := range input {
		wasFound := false
		for _, j := range window {
			diff := int(math.Abs(float64(j - v)))
			if lookup[diff] {
				wasFound = true
				delete(lookup, window[0])
				window = append(window[1:], v)
				lookup[v] = true
				break
			}
		}
		if !wasFound {
			return v
		}
	}
	return -1
}

func problem2(fp string, value int) int {
	input := []int{}
	for line := range adventofcode2020.FileIterator(fp) {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, v)
	}

	for i := 0; i < len(input); i++ {
		sum := 0
		list := []int{}
		for _, v := range input[i:] {
			sum += v
			list = append(list, v)
			if sum == value {
				sort.Ints(list)
				return list[0] + list[len(list)-1]
			}
			if sum > value {
				break
			}
		}
	}
	return -1
}
