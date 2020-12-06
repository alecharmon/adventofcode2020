package day2

import (
	"strconv"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

type policy struct {
	least int
	max   int
	rule  string
}

func createPolicy(least, max int, rule string) policy {
	return policy{
		least: least,
		max:   max,
		rule:  rule,
	}
}

func (p policy) isValid(input string) bool {
	c := strings.Count(input, p.rule)
	return c >= p.least && c <= p.max
}

func problem1(fp string) int {
	count := 0
	for line := range adventofcode2020.FileIterator(fp) {
		s := strings.Split(line, ":")
		policyString, input := s[0], s[1]
		policyParts := strings.Split(policyString, " ")
		policyRange, policyRule := policyParts[0], strings.TrimSpace(policyParts[1])
		policyRangeParts := strings.Split(policyRange, "-")

		min, _ := strconv.Atoi(policyRangeParts[0])
		max, _ := strconv.Atoi(policyRangeParts[1])
		policy := createPolicy(min, max, policyRule)
		if policy.isValid(input) {
			count++
		}
	}
	return count
}

type policy2 struct {
	firstPosition  int
	secondPosition int
	rule           rune
}

func createPolicy2(firstPosition, secondPosition int, rule rune) policy2 {
	return policy2{
		firstPosition:  firstPosition,
		secondPosition: secondPosition,
		rule:           rule,
	}
}

func (p policy2) isValid(input string) bool {

	p1 := (p.rule == rune(input[p.firstPosition-1]))
	p2 := (p.rule == rune(input[p.secondPosition-1]))

	if !(p1 || p2) {
		return false
	}
	if p1 == p2 {
		return false
	}

	return true
}

func problem2(fp string) int {
	count := 0
	for line := range adventofcode2020.FileIterator(fp) {
		s := strings.Split(line, ":")

		policyString, input := s[0], strings.TrimSpace(s[1])
		policyParts := strings.Split(policyString, " ")
		policyRange, policyRule := policyParts[0], strings.TrimSpace(policyParts[1])
		policyRangePositions := strings.Split(policyRange, "-")

		p1, _ := strconv.Atoi(policyRangePositions[0])
		p2, _ := strconv.Atoi(policyRangePositions[1])

		policy := createPolicy2(p1, p2, []rune(policyRule)[0])
		output := policy.isValid(input)
		if output {
			count++
		}

	}
	return count
}
