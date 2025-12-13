package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filenamePtr := flag.String("input", "", "the input file")
	flag.Parse()

	if len(*filenamePtr) == 0 {
		println("You need to specify an input file with --input")
		return
	}

	file, err := os.Open(*filenamePtr)
	if err != nil {
		println("Unable to open file: ", err)
	}
	defer file.Close()

	sum := 50
	count := 0
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])

		if direction == 'L' {
			sum += value
		} else {
			sum -= value
		}

		if sum%100 == 0 {
			count += 1
		}
		lines += 1
	}
	fmt.Printf("Count: %d. Lines read: %d.\n", count, lines)
}
