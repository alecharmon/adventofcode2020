package day3

import (
	"strconv"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

type bagEdge struct {
	to, from *bag
	qty      int
}
type bag struct {
	canBeContainedBy []*bagEdge
	name             string
	canContain       []*bagEdge
	visted           bool
}

var lookup = make(map[string]*bag)

func createEdge(from, to *bag, qty int) *bagEdge {
	return &bagEdge{
		to:   to,
		from: from,
		qty:  qty,
	}
}

func getOrCreateBag(name string) *bag {
	if b, ok := lookup[strings.TrimSpace(name)]; ok {
		return b
	}
	b := &bag{
		canBeContainedBy: []*bagEdge{},
		canContain:       []*bagEdge{},
		visted:           false,
		name:             strings.TrimSpace(name),
	}
	lookup[b.name] = b
	return b
}

func removeBagPostfix(s string) string {
	return strings.Replace(strings.Replace(s, "bags", "", -1), "bag", "", -1)
}

func createGraphFromFile(fp string) {
	for line := range adventofcode2020.FileIterator(fp) {
		data := strings.Split(line, "contain")
		b, contains := getOrCreateBag(removeBagPostfix(data[0])), strings.Split(data[1:][0], ",")
		if len(contains) == 1 && contains[0] == " no other bags." {
			continue
		}
		for _, bagName := range contains {
			childData := strings.Split(strings.TrimSpace(strings.Replace(bagName, ".", "", -1)), " ")
			bagCount, bagName := childData[0], removeBagPostfix(strings.Join(childData[1:], " "))
			bagCountInt, _ := strconv.Atoi(bagCount)
			child := getOrCreateBag(bagName)

			child.canBeContainedBy = append(child.canBeContainedBy, createEdge(child, b, bagCountInt))
			b.canContain = append(b.canContain, createEdge(b, child, bagCountInt))
		}
	}
}

func problem1(fp string) int {
	createGraphFromFile(fp)
	gold := lookup["shiny gold"]
	stack := []*bagEdge{}
	stack = append(stack, gold.canBeContainedBy...)

	count := 0
	for len(stack) > 0 {
		b := stack[len(stack)-1].to
		stack = stack[:len(stack)-1]
		if b.visted == false {
			b.visted = true
			stack = append(stack, b.canBeContainedBy...)
			count++
		}
	}
	return count
}

func problem2Helper(node, top *bag) int {
	count := 0
	if node != top {
		count++
	}

	for _, child := range node.canContain {
		b := child.to
		countRec := problem2Helper(b, top)
		count += child.qty * max(1, countRec)
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func problem2(fp string) int {
	createGraphFromFile(fp)
	gold := lookup["shiny gold"]

	return problem2Helper(gold, gold)
}
