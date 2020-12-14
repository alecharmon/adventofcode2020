package day14

import (
	"regexp"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

type mask struct {
	mask []string
}

type maskSet struct {
	m   mask
	mem map[int]int
}

func (m mask) applyValue(in int) int {
	result := 0
	l := len(m.mask) - 1
	for i, c := range m.mask {

		switch c {
		case "X":
			subMask := int(1 << (l - i))
			if subMask&in == subMask {
				result += subMask
			}

		case "1":
			subMask := int(1 << (l - i))
			result += subMask

		}

	}
	return result
}

func copyMask(in mask) mask {
	m := []string{}
	for _, c := range in.mask {
		m = append(m, c)
	}
	return mask{mask: m}
}

func applyAddress(in int, m mask) []int {
	current := 0
	l := len(m.mask) - 1
	for i, c := range m.mask {
		switch c {
		case "X":
			zeroCopy := copyMask(m)
			oneCopy := copyMask(m)
			zeroCopy.mask[i] = "0"
			oneCopy.mask[i] = "1"
			results := []int{}
			results = append(results, applyAddress(in, zeroCopy)...)
			results = append(results, applyAddress(in, oneCopy)...)
			return results
		case "1":
			subMask := int(1 << (l - i))
			current += subMask

		case "U":
			subMask := int(1 << (l - i))
			if subMask&in == subMask {
				current += subMask
			}
		}

	}
	return []int{current}
}
func createMask(maskValue []string) mask {
	return mask{
		mask: maskValue,
	}
}

func problem1(fp string) int {
	maskRegex := regexp.MustCompile(`mask = ([10X]+)`)
	valuesRegex := regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)
	mem := map[int]int{}
	var current mask
	for line := range adventofcode2020.FileIterator(fp) {
		if matches := maskRegex.FindStringSubmatch(line); len(matches) > 0 {
			current = createMask(strings.Split(matches[1], ""))
			continue
		}

		matches := valuesRegex.FindStringSubmatch(line)

		addr, value := adventofcode2020.StringToInt(matches[1]), adventofcode2020.StringToInt(matches[2])
		mem[addr] = int(current.applyValue(int(value)))
	}

	sum := 0

	for _, v := range mem {
		sum += v
	}

	return sum
}

func problem2(fp string) int {
	maskRegex := regexp.MustCompile(`mask = ([10X]+)`)
	valuesRegex := regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)
	mem := map[int]int{}
	var current mask
	for line := range adventofcode2020.FileIterator(fp) {
		if matches := maskRegex.FindStringSubmatch(line); len(matches) > 0 {
			tmp := []string{}
			for _, c := range strings.Split(matches[1], "") {
				if c == "0" {
					tmp = append(tmp, "U")
				} else {
					tmp = append(tmp, c)
				}
			}
			current = createMask(tmp)
			continue
		}

		matches := valuesRegex.FindStringSubmatch(line)
		addr, value := adventofcode2020.StringToInt(matches[1]), adventofcode2020.StringToInt(matches[2])
		addresses := applyAddress(addr, current)
		for _, v := range addresses {
			mem[v] = value
		}
	}

	sum := 0

	for _, v := range mem {
		sum += v
	}

	return sum
}
