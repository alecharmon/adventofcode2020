package day3

import (
	"fmt"
	"sort"

	"github.com/alecharmon/adventofcode2020"
)

func problem1(fp string) int {
	maxId := 0
	//there is probs a way of doing this without replicating bin search but this seems to
	//be more in the spirit of the question
	for line := range adventofcode2020.FileIterator(fp) {
		data := []rune(line)

		row, col := data[:7], data[7:]

		rowMax := 127
		rowMin := 0
		rowMid := (rowMax + rowMin) / 2
		for _, val := range row {
			if val == 'F' {
				rowMax = rowMid
			} else {
				rowMin = rowMid

			}

			rowMid = (rowMax + rowMin) / 2
			if rowMid == rowMin || rowMid == rowMax {
				rowMid++
			}
			fmt.Printf(" %s, mid %d, min %d max %d \n", string(val), rowMid, rowMin, rowMax)
		}

		colMax := 7
		colMin := 0
		colMid := (colMax + colMin) / 2
		for _, val := range col {
			if val == 'R' {
				colMin = colMid
			} else {
				colMax = colMid
			}

			colMid = ((colMax + colMin) / 2)
			if colMid == colMin || colMid == colMax {
				colMid++
			}

			fmt.Printf(" %s, mid %d, min %d max %d \n", string(val), colMid, colMin, colMax)
		}

		id := rowMid*8 + colMid
		fmt.Printf("col %d row %d : id %d \n", colMid, rowMid, id)
		maxId = max(maxId, id)
	}
	return maxId
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func problem2(fp string) int {
	ids := []int{}
	//there is probs a way of doing this without replicating bin search but this seems to
	//be more in the spirit of the question
	for line := range adventofcode2020.FileIterator(fp) {
		data := []rune(line)

		row, col := data[:7], data[7:]

		rowMax := 127
		rowMin := 0
		rowMid := (rowMax + rowMin) / 2
		for _, val := range row {
			if val == 'F' {
				rowMax = rowMid
			} else {
				rowMin = rowMid

			}

			rowMid = (rowMax + rowMin) / 2
			if rowMid == rowMin || rowMid == rowMax {
				rowMid++
			}
			fmt.Printf(" %s, mid %d, min %d max %d \n", string(val), rowMid, rowMin, rowMax)
		}

		colMax := 7
		colMin := 0
		colMid := (colMax + colMin) / 2
		for i, val := range col {
			if val == 'R' {
				colMin = colMid
			} else {
				colMax = colMid
			}

			if i == 2 {
				if val == 'R' {
					colMid = colMax
				} else {
					colMid = colMax
				}
				continue
			}

			colMid = ((colMax + colMin) / 2)
			fmt.Printf(" %s, mid %d, min %d max %d \n", string(val), colMid, colMin, colMax)
		}

		id := rowMid*8 + colMid
		fmt.Printf("col %d row %d : id %d \n", colMid, rowMid, id)

		ids = append(ids, rowMid*8+colMid)

	}

	//meb this could be done faster with a ordered dict? or insert sort
	sort.Ints(ids)
	fmt.Println(ids)
	for i, id := range ids {
		if ids[i+1]-id > 1 {
			return id + 1
		}
	}
	return -1
}
