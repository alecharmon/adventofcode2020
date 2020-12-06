package day3

import (
	"github.com/alecharmon/adventofcode2020"
)

func traverse(fp string, right, down int) int {
	count := 0
	current := 0
	width := 0
	tmpDown := 1
	for line := range adventofcode2020.FileIterator(fp) {
		if tmpDown > 1 {
			tmpDown--
			continue
		}
		if width == 0 {
			width = len(line)
		}
		if line[current] == '#' {
			count++
		}
		current = (current + right) % width
		if tmpDown == 1 && down > 1 {
			tmpDown = down
		}
	}
	return count
}

func problem1(fp string) int {
	return traverse(fp, 3, 1)
}

func problem2(fp string) int {
	// Right 1, down 1.
	// Right 3, down 1.
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.
	return traverse(fp, 1, 1) * traverse(fp, 3, 1) * traverse(fp, 5, 1) * traverse(fp, 7, 1) * traverse(fp, 1, 2)
}
