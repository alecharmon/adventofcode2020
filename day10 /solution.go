package day10

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/alecharmon/adventofcode2020"
)

func seperateValues(input []int, pivot int) (works []int, doesNotWork []int) {
	for _, v := range input {
		if v-pivot <= 3 && v-pivot > -1 {
			works = append(works, v)
			continue
		}
		doesNotWork = append(doesNotWork, v)
	}
	return
}

func keysFromMap(m map[int]bool) (keys []int) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	return
}

func copyMap(originalMap map[int]bool) map[int]bool {
	// Create the target map
	targetMap := make(map[int]bool)

	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}
	return targetMap
}

func mapKey(originalMap map[int]bool) string {
	keys := keysFromMap(originalMap)
	sort.Ints(keys)
	key := fmt.Sprint(keys)
	fmt.Println(key)
	return key
}

func findPaths(remainder map[int]bool, current []int, last int, mem map[string]int) int {
	var pivot int
	if len(current) == 0 {
		pivot = 0
	} else {
		pivot = current[len(current)-1]
	}
	if len(keysFromMap(remainder)) > 0 {
		fmt.Println(keysFromMap(remainder))
	}

	if last == pivot {
		return 1
	}

	if val, ok := mem[mapKey(remainder)]; ok {
		return val

	}
	sum := 0

	for k, _ := range remainder {
		if k-pivot <= 3 && k-pivot >= 0 {

			copy := copyMap(remainder)
			delete(copy, k)
			sum += findPaths(copy, append(current, k), last, mem)

		} else if k-pivot < 0 {
			delete(remainder, k)
		}
	}
	mem[mapKey(remainder)] = sum
	return sum
}

func problem2(fp string) int {
	input := make(map[int]bool)
	max := 0
	for line := range adventofcode2020.FileIterator(fp) {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input[v] = true
		if v > max {
			max = v
		}
	}

	return findPaths(input, []int{}, max, make(map[string]int))
}

func problem1(fp string) int {
	input := []int{}
	for line := range adventofcode2020.FileIterator(fp) {
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, v)
	}

	sort.Ints(input)

	diffs := make(map[int]int)
	diffs[1] = 0
	diffs[2] = 0
	diffs[3] = 0
	prev := 0
	for _, v := range input {
		diffs[v-prev] = diffs[v-prev] + 1
		prev = v
	}

	//for our last power adapter
	diffs[3] = diffs[3] + 1

	fmt.Println(diffs)
	return diffs[1] * diffs[3]
}
