package adventofcode2020

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func FileIterator(fp string) chan string {
	ch := make(chan string)

	file, err := os.Open(fp)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(ch)
		defer file.Close()
		for scanner.Scan() {
			ch <- strings.Replace(scanner.Text(), "\n", "", -1)
		}
	}()
	return ch
}
