package day1

import (
	"errors"
	"log"
	"strconv"

	"github.com/alecharmon/adventofcode2020"
)

func problem1(fp string) (int, error) {
	lookup := make(map[int]bool)
	for line := range adventofcode2020.FileIterator(fp) {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		lookup[val] = true

		complement := 2020 - val
		if _, ok := lookup[complement]; ok {
			return complement * val, nil
		}
	}
	return -1, errors.New("No pair found")
}

func problem2(fp string) (int, error) {
	lookup := make(map[int]bool)
	for line := range adventofcode2020.FileIterator(fp) {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		lookup[val] = true

		for key, _ := range lookup {
			complement := 2020 - val - key
			if complement > 0 {
				if _, ok := lookup[complement]; ok {
					return complement * key * val, nil
				}
			}
		}
	}
	return -1, errors.New("No pair found")
}
