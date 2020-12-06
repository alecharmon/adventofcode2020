package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecharmon/adventofcode2020"
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)

type id struct {
	store map[string]string
}

func createID(values map[string]string) id {
	// getWithDefault := func(key string) string {
	// 	val, ok := values[key]
	// 	if !ok {
	// 		return ""
	// 	}
	// 	return val
	// }

	return id{
		store: values,
	}
}

func (i id) validate() bool {
	neededKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, key := range neededKeys {

		if _, ok := i.store[key]; !ok {
			return false
		}
	}
	return true
}

func (i id) validateAndValues() bool {

	var validationRules = []struct {
		key  string
		rule func(input string) bool
	}{
		{
			"byr",
			func(input string) bool {
				val, _ := strconv.Atoi(input)
				return val >= 1920 && val <= 2002
			},
		},
		{
			"iyr",
			func(input string) bool {
				val, _ := strconv.Atoi(input)
				return val >= 2010 && val <= 2020
			},
		},
		{
			"eyr",
			func(input string) bool {
				val, _ := strconv.Atoi(input)
				return val >= 2020 && val <= 2030
			},
		},
		{
			"hgt",
			func(input string) bool {
				if strings.Contains(input, "cm") {
					num := strings.Replace(input, "cm", "", -1)
					val, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					return val >= 150 && val <= 193
				}
				if strings.Contains(input, "in") {
					num := strings.Replace(input, "in", "", -1)
					val, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					return val >= 59 && val <= 76
				}
				return false
			},
		},
		{
			"hcl",
			func(input string) bool {
				pattern := regexp.MustCompile(`#[0-9a-f]{6}`)
				return pattern.MatchString(input)
			},
		},
		{
			"ecl",
			func(input string) bool {
				fmt.Print(input, "\n")
				valid := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				for _, option := range valid {
					if input == option {
						return true
					}
				}
				return false
			},
		},
		{
			"pid",
			func(input string) bool {
				pattern := regexp.MustCompile(`[0-9]{9}`)
				return pattern.MatchString(input)
			},
		},
	}

	for _, key := range validationRules {
		value, ok := i.store[key.key]
		if !ok {
			return false
		}
		if key.rule(value) == false {
			fmt.Printf("rule for %s failed validation, value %s \n", key.key, value)
			return false
		}

	}
	return true
}

func createIdFromString(input string) id {
	re, _ := regexp.Compile(`([a-zA-Z]{3}:[#a-zA-Z0-9]+)`)
	res := re.FindAllStringSubmatch(input, -1)
	m := make(map[string]string)
	for _, match := range res {
		parts := strings.Split(match[0], ":")
		m[parts[0]] = parts[1]
	}

	idInstance := createID(m)
	fmt.Println(idInstance)
	return idInstance
}
func problem1(fp string) int {
	group := ""
	count := 0
	for line := range adventofcode2020.FileIterator(fp) {
		if line == "" {

			idInstance := createIdFromString(group)

			if idInstance.validate() {
				count++
			}
			group = ""
			continue
		}
		group += line + " "
	}

	if group != "" {
		idInstance := createIdFromString(group)

		if idInstance.validate() {
			count++
		}
	}
	return count
}

func problem2(fp string) int {
	group := ""
	count := 0
	for line := range adventofcode2020.FileIterator(fp) {
		if line == "" {

			idInstance := createIdFromString(group)

			if idInstance.validateAndValues() {
				count++
			}
			group = ""
			continue
		}
		group += line + " "
	}

	if group != "" {
		idInstance := createIdFromString(group)

		if idInstance.validateAndValues() {
			count++
		}
	}
	return count
}
