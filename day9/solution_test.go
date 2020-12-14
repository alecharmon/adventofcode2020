package day9

import (
	"log"
	"testing"
)

func TestProblem1(t *testing.T) {
	log.Println(problem1("./input.txt", 25))
}

func TestProblem2(t *testing.T) {
	// log.Println(problem2("./test.txt", 127))
	log.Println(problem2("./input.txt", 507622668))
}
