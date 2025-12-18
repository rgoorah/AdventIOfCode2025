package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Tuple struct {
	Index   int
	Joltage int
}

func findHighest(batteries string, start int, ignore int) Tuple {
	var highest Tuple
	for i := start; i < len(batteries); i++ {

		if i == ignore {
			continue
		}

		jolt := int(batteries[i] - '0')
		if jolt > highest.Joltage {
			highest.Joltage = jolt
			highest.Index = i
		}
	}
	return highest
}

func highestJoltage(batteries string) int {
	firstHighest := findHighest(batteries, 0, -1)
	secondHighest := findHighest(batteries, 0, firstHighest.Index)

	// fmt.Printf("first: %d %d\n", firstHighest.Index, firstHighest.Joltage)
	// fmt.Printf("Second: %d %d\n", secondHighest.Index, secondHighest.Joltage)

	var output string
	if firstHighest.Index < secondHighest.Index {
		output += strconv.Itoa(firstHighest.Joltage)
		output += strconv.Itoa(secondHighest.Joltage)
	} else if firstHighest.Index != len(batteries)-1 {
		secondHighest = findHighest(batteries, firstHighest.Index+1, -1)
		output += strconv.Itoa(firstHighest.Joltage)
		output += strconv.Itoa(secondHighest.Joltage)
	} else {
		output += strconv.Itoa(secondHighest.Joltage)
		output += strconv.Itoa(firstHighest.Joltage)
	}

	o, _ := strconv.Atoi(output)

	return o
}

func main() {
	filenamePtr := flag.String("input", "", "the input file")
	flag.Parse()

	if len(*filenamePtr) == 0 {
		println("You need to specify an input file with --input")
		return
	}

	file, err := os.Open(*filenamePtr)
	if err != nil {
		println("Unable to open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		joltage := highestJoltage(line)
		// fmt.Printf("Line: %s, Joltage:%d\n", line, joltage)
		sum += joltage
	}

	fmt.Printf("Sum: %d\n", sum)
}
