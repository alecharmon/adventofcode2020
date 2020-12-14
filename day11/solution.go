package day11

import (
	"fmt"

	"github.com/alecharmon/adventofcode2020"
)

const (
	EMPTY    = iota
	OCCUPIED = iota
)

type seat struct {
	status int
}
type board struct {
	layout [][]*seat
}

func createSeat(s string) *seat {
	switch s {
	case "L":
		return &seat{
			status: EMPTY,
		}
	case "#":
		return &seat{
			status: OCCUPIED,
		}
	default:
		return nil
	}
}

func createBoard(input [][]string) board {
	layout := [][]*seat{}
	for _, row := range input {
		seatRow := []*seat{}
		for _, s := range row {
			seatRow = append(seatRow, createSeat(s))
		}
		layout = append(layout, seatRow)
	}
	return board{
		layout: layout,
	}
}

func adjacent(x, y int, b board) int {
	points := []struct {
		x int
		y int
	}{
		{
			x: 0,
			y: 1,
		},
		{
			x: 1,
			y: 0,
		},
		{
			x: 1,
			y: 1,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: -1,
			y: 0,
		},
		{
			x: -1,
			y: -1,
		},
		{
			x: -1,
			y: 1,
		},
		{
			x: 1,
			y: -1,
		},
	}

	sum := 0
	for _, point := range points {
		xP, yP := x+point.x, y+point.y
		if xP < 0 || yP < 0 || xP >= len(b.layout[0]) || yP >= len(b.layout) {
			continue
		}

		if s := b.layout[yP][xP]; s != nil && s.status == OCCUPIED {
			sum++
		}
	}
	return sum
}

type point struct {
	x int
	y int
}

type scaleFunc = func(next point) point

func scale(vector point) scaleFunc {
	return func(next point) point {
		return point{x: vector.x + next.x, y: vector.y + next.y}
	}
}

func adjacent2(x, y int, b board) int {

	scaleDirections := []scaleFunc{
		scale(point{
			x: 0,
			y: 1,
		}),
		scale(point{
			x: 1,
			y: 0,
		}),
		scale(point{
			x: 1,
			y: 1,
		}),
		scale(point{
			x: 0,
			y: -1,
		}),
		scale(point{
			x: -1,
			y: 0,
		}),
		scale(point{
			x: -1,
			y: -1,
		}),
		scale(point{
			x: -1,
			y: 1,
		}),
		scale(point{
			x: 1,
			y: -1,
		}),
	}

	sum := 0
	for _, scaleDir := range scaleDirections {
		next := scaleDir(point{x: x, y: y})
		for next.x >= 0 && next.y >= 0 && next.x < len(b.layout[0]) && next.y < len(b.layout) {
			s := b.layout[next.y][next.x]
			if s != nil {
				if s.status == OCCUPIED {
					sum++
				}
				break
			}
			next = scaleDir(next)
		}
	}
	return sum
}

func copyBoard(b board) board {
	layout := [][]*seat{}
	for _, row := range b.layout {
		newRow := []*seat{}
		for _, s := range row {
			if s != nil {
				newRow = append(newRow, &seat{
					status: s.status,
				})
				continue
			}
			newRow = append(newRow, nil)
		}
		layout = append(layout, newRow)
	}
	return board{
		layout: layout,
	}
}

func (b board) countOccupied() int {
	count := 0
	for _, row := range b.layout {
		for _, seat := range row {
			if seat == nil {
				fmt.Print(".")
				continue
			}
			if seat.status == OCCUPIED {
				count++
			}

		}
	}
	return count
}

func (b board) simulate() bool {
	cb := copyBoard(b)
	didChange := false
	for y, row := range b.layout {
		for x, seat := range row {
			if seat == nil {
				fmt.Print(".")
				continue
			}
			adj := adjacent(x, y, cb)

			if seat.status == OCCUPIED && adj >= 4 {
				seat.status = EMPTY
				didChange = true
			}

			if seat.status == EMPTY && adj == 0 {
				seat.status = OCCUPIED
				didChange = true
			}
			fmt.Print(adj)
		}
		fmt.Println()
	}
	return didChange
}

func (b board) simulate2() bool {
	cb := copyBoard(b)
	didChange := false
	for y, row := range b.layout {
		for x, seat := range row {
			if seat == nil {
				fmt.Print(".")
				continue
			}
			adj := adjacent2(x, y, cb)

			if seat.status == OCCUPIED && adj >= 5 {
				seat.status = EMPTY
				didChange = true
			}

			if seat.status == EMPTY && adj == 0 {
				seat.status = OCCUPIED
				didChange = true
			}
			fmt.Print(adj)
		}
		fmt.Println()
	}
	return didChange
}

func problem1(fp string) int {
	input := [][]string{}
	for line := range adventofcode2020.FileIterator(fp) {
		row := []string{}
		for _, r := range line {
			row = append(row, string(r))
		}
		input = append(input, row)
	}

	b := createBoard(input)
	count := 0
	for {
		fmt.Println("----------------")
		if !b.simulate() {
			return b.countOccupied()
		}
	}

	return count
}

func problem2(fp string) int {
	input := [][]string{}
	for line := range adventofcode2020.FileIterator(fp) {
		row := []string{}
		for _, r := range line {
			row = append(row, string(r))
		}
		input = append(input, row)
	}

	b := createBoard(input)
	count := 0
	for {
		fmt.Println("----------------")
		if !b.simulate2() {
			return b.countOccupied()
		}
	}

	return count
}
