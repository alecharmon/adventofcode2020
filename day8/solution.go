package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

//offset, countAgg
type opFunc func(input int) (int, int)
type op struct {
	f     opFunc
	value int
	name  string
}

func (o op) exec() (int, int) {
	return o.f(o.value)
}

func (o op) flip() op {
	if o.name == "nop" {
		o.f = jmp
		o.name = "jmp"
	} else if o.name == "jmp" {
		o.f = nop
		o.name = "nop"
	}
	return o
}

func jmp(input int) (int, int) {
	return input, 0
}

func acc(input int) (int, int) {
	return 1, input
}

func nop(input int) (int, int) {
	return 1, 0
}

func createOp(opType string, value int) op {
	var f opFunc
	switch opType {
	case "nop":
		f = nop
	case "acc":
		f = acc
	case "jmp":
		f = jmp
	}
	return op{
		f:     f,
		value: value,
		name:  opType,
	}
}
func problem1(fp string) int {
	ops := []op{}
	for line := range adventofcode2020.FileIterator(fp) {
		data := strings.Split(line, " ")
		opType, valueString := data[0], data[1]
		value, _ := strconv.Atoi(valueString)
		ops = append(ops, createOp(opType, value))
	}

	opsVisted := make([]bool, len(ops))
	sum := 0
	current := 0
	for {
		if opsVisted[current] == true {
			break
		}
		opsVisted[current] = true
		currentOp := ops[current]
		currentOffset, sumOffset := currentOp.exec()
		sum += sumOffset
		current += currentOffset
	}
	return sum
}

func process(ops []op) (sum int, successful bool, prev []int) {
	opsVisted := make([]bool, len(ops))
	sum = 0
	current := 0
	for {
		if current >= len(ops) {
			successful = true
			return
		}
		if opsVisted[current] == true {
			successful = false
			return
		}
		opsVisted[current] = true
		currentOp := ops[current]
		currentOffset, sumOffset := currentOp.exec()
		if currentOp.name == "nop" || currentOp.name == "jmp" {
			prev = append(prev, current)
		}

		sum += sumOffset
		current += currentOffset
	}
}

func problem2(fp string) int {
	ops := []op{}
	for line := range adventofcode2020.FileIterator(fp) {
		data := strings.Split(line, " ")
		opType, valueString := data[0], data[1]
		value, _ := strconv.Atoi(valueString)
		ops = append(ops, createOp(opType, value))
	}

	_, _, prev := process(ops)
	//run through previous stack and try changing each one
	for i := len(prev) - 1; i > 0; i-- {
		toChange := prev[i]
		ops[toChange] = ops[toChange].flip()
		sum, successful, _ := process(ops)
		fmt.Println(sum, successful, toChange, i)
		if successful {
			return sum
		}
		ops[toChange] = ops[toChange].flip()
	}
	return 0

}
