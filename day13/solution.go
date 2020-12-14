package day13

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

type direction struct {
	value    string
	opposite string
}

func problem1(fp string) int {
	busIds := []int{}

	input := adventofcode2020.FileIterator(fp)
	ts := adventofcode2020.StringToInt(<-input)

	for _, id := range strings.Split(<-input, ",") {
		if id != "x" {
			busIds = append(busIds, adventofcode2020.StringToInt(id))
		}
	}

	closet, remainder := 0, 1<<62
	for _, v := range busIds {
		plausibleTime := v
		for plausibleTime < ts {
			plausibleTime += v
		}
		tmpRemainder := plausibleTime - ts
		if tmpRemainder < remainder {
			remainder = tmpRemainder
			closet = v
		}
	}

	return closet * remainder
}

type Bus struct {
	id    int
	order int
}

func problem2(fp string) int {
	buses := []int{}

	input := adventofcode2020.FileIterator(fp)
	<-input

	for _, id := range strings.Split(<-input, ",") {
		if id == "x" {
			buses = append(buses, 1)
			continue
		}
		buses = append(buses, adventofcode2020.StringToInt(id))
	}
	timestamp := 1

	for {
		timeToSkipIfNoMatch := 1
		valid := true

		for offset := 0; offset < len(buses); offset++ {
			// A given bus will depart when the timestamp is evenly divisible by the bus ID
			if (timestamp+offset)%buses[offset] != 0 {
				// No match here; abort and we'll try the next potential timestamp
				valid = false
				break
			}
			fmt.Println(timeToSkipIfNoMatch)
		}

		// Did we find a full match?
		if valid {
			return timestamp
		}

		timestamp += timeToSkipIfNoMatch
	}
}

type Departure struct {
	b  Bus
	ts float64
}

func attemptToBuild(ts float64, list []string) bool {
	log.Printf("%.10f\n", ts)
	busIndex := 1
	for i := float64(ts) + 1; i < float64(1<<64); i++ {
		if len(list) == busIndex {
			return true
		}
		if list[busIndex] == "x" {
			busIndex++
			continue
		}

		id := float64(adventofcode2020.StringToInt(list[busIndex]))
		if math.Mod(i, id) == 0 {
			busIndex++
			continue
		}
		return false
	}
	return true
}

func cantAdd(ts float64, b Bus, schedule []Departure, list []Bus) bool {
	if len(schedule) < 2 {
		return false
	}
	tmp := schedule[0].ts
	for tmp <= ts {
		for _, departure := range schedule {
			if math.Mod(tmp, float64(departure.b.id)) == 0 && tmp != departure.ts {
				comparingToFirst := departure.b.id == list[0].id
				isLast := b.id == list[len(list)-1].id
				if comparingToFirst && isLast {
					return tmp == departure.ts
				}
				return true
			}
		}
		tmp++
	}

	return false
}

func isMonotonic(list []Departure) bool {
	prev := list[0].ts
	for _, bus := range list {
		if prev > bus.ts {
			return false
		}
		prev = bus.ts
	}
	return true
}
