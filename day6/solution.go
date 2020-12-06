package day3

import (
	"fmt"
	"sort"

	"github.com/alecharmon/adventofcode2020"
)

func countGroup(s string) int {
	sum := 0
	localDict := make(map[rune]bool)
	for _, char := range s {
		localDict[char] = true
	}
	fmt.Println(localDict)
	for _, _ = range localDict {
		sum++
	}
	return sum
}

func problem1(fp string) int {
	sum := 0
	group := ""
	for line := range adventofcode2020.FileIterator(fp) {
		if line == "" {
			sum += countGroup(group)
			group = ""
		} else {
			group += line
		}
	}
	sum += countGroup(group)
	return sum
}

func countGroupIndividual(s string) int {
	people := []map[rune]bool{}
	person := make(map[rune]bool)
	for _, char := range s {
		if char == '|' {
			people = append(people, person)
			person = make(map[rune]bool)
			continue
		}
		person[char] = true
	}
	if len(person) > 0 {
		people = append(people, person)
	}

	if len(people) == 1 {
		return len(people[0])
	}

	sort.SliceStable(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)

	sum := 0
	longest, remanding := people[0], people[1:]
	for k, _ := range longest {
		everoneHas := true
		for _, otherPerson := range remanding {
			if otherPerson[k] != true {
				everoneHas = false
			}
		}
		if everoneHas {
			sum++
		}
	}
	return sum
}

func problem2(fp string) int {
	sum := 0
	group := ""
	for line := range adventofcode2020.FileIterator(fp) {
		if line == "" {
			sum += countGroupIndividual(group)
			group = ""
		} else {
			group += line + "|"
		}
	}
	sum += countGroupIndividual(group)
	return sum
}
