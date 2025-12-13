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

	current := 50
	count := 0
	lines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		value, _ := strconv.Atoi(line[1:])

		var prev int = current

		if value > 100 {
			count += value / 100         // Number of passes
			value -= 100 * (value / 100) // reset the value
		}

		if direction == 'L' {
			current -= value
		} else {
			current += value
		}

		if current < 0 {
			if prev != 0 {
				count += 1
			}
			current = 100 + current
		}
		if current > 100 {
			count += 1
			current = current - 100
		}

		if current == 0 || current == 100 {
			count += 1
			current = 0
		}

		lines += 1
	}
	fmt.Printf("Count: %d. Lines read: %d.\n", count, lines)
}
