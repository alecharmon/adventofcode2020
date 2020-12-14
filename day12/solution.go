package day12

import (
	"fmt"

	"github.com/alecharmon/adventofcode2020"
)

type direction struct {
	value    string
	opposite string
}

var dirs = map[string]direction{
	"E": direction{value: "E", opposite: "W"},
	"W": direction{value: "W", opposite: "E"},
	"N": direction{value: "N", opposite: "S"},
	"S": direction{value: "S", opposite: "N"},
}

type ship struct {
	currentDir string
	position   map[string]int
}

type instruction struct {
	action string
	value  int
}

func (s *ship) rotate(i instruction) {
	order := []string{"N", "E", "S", "W"}
	current := 0
	for i, v := range order {
		if v == s.currentDir {
			current = i
		}
	}

	deg := i.value
	if i.action == "L" {
		deg = i.value * -1
	}

	degCount := adventofcode2020.Abs(deg / 90)
	tmpCount := 0

	if i.action == "L" {
		for tmpCount != degCount {
			current--
			tmpCount++
		}
		if current < 0 {
			current = current + 4
		}
		s.currentDir = order[current]
	} else {
		for tmpCount != degCount {
			current++
			tmpCount++
		}
		if current >= 4 {
			current = current % 4
		}
		s.currentDir = order[current]
	}

}

func (s *ship) exec(i instruction) {

	switch i.action {
	case "L":
		s.rotate(i)
	case "R":
		s.rotate(i)
	case "F":
		s.moveInDir(dirs[s.currentDir], i.value)
	default:
		s.moveInDir(dirs[i.action], i.value)
	}
	fmt.Println(i, s.currentDir, s.position)
}

func (s *ship) moveInDir(dir direction, value int) {
	if s.position[dir.opposite] > 0 {
		diff := s.position[dir.opposite] - value
		if diff < 0 {
			s.position[dir.opposite] = 0
			s.position[dir.value] = adventofcode2020.Abs(diff)
		} else {
			s.position[dir.opposite] = diff
		}
	} else {
		s.position[dir.value] = s.position[dir.value] + value
	}
}

func createShip(dir string) *ship {
	return &ship{
		currentDir: dir,
		position:   make(map[string]int),
	}
}

func problem1(fp string) int {
	instructions := []instruction{}
	for line := range adventofcode2020.FileIterator(fp) {
		instructions = append(instructions,
			instruction{
				action: string(line[0]),
				value:  adventofcode2020.StringToInt(line[1:]),
			},
		)
	}

	s := createShip("E")

	for idx, i := range instructions {
		s.exec(i)
		sum := 0
		for _, v := range s.position {
			sum += v
		}
		fmt.Println(idx, sum)
	}

	sum := 0
	for _, v := range s.position {
		sum += v
	}
	return sum
}

// func problem2(fp string) int {
// 	lookup := make(map[int]bool)
// 	for line := range adventofcode2020.FileIterator(fp) {
// 		val, err := strconv.Atoi(line)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		lookup[val] = true

// 		for key, _ := range lookup {
// 			complement := 2020 - val - key
// 			if complement > 0 {
// 				if _, ok := lookup[complement]; ok {
// 					return complement * key * val, nil
// 				}
// 			}
// 		}
// 	}
// 	return -1, errors.New("No pair found")
// }
