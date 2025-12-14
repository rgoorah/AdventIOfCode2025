package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isDouble(value int) bool {
	sValue := strconv.Itoa(value)

	end := len(sValue)
	if end%2 == 1 {
		return false
	}

	current := 0
	mid := int(end / 2)
	for current < mid {
		if sValue[current] != sValue[mid+current] {
			return false
		}
		current += 1
	}
	return true
}

func getDoubles(rangeStr string) []int {

	st := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(st[0])
	end, _ := strconv.Atoi(st[1])

	doubles := []int{}
	for i := start; i <= end; i++ {
		if isDouble(i) {
			doubles = append(doubles, i)
		}

	}
	return doubles
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
		println("Unable to open file: ", err)
	}
	defer file.Close()

	var rangeStrs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // all the data should be on a single line
		rangeStrs = strings.Split(line, ",")
	}

	sum := 0
	for _, rangeStr := range rangeStrs {
		doubles := getDoubles(rangeStr)

		// sDoubles := strings.Split(fmt.Sprint(doubles), ",")
		// fmt.Printf("Range: %s Doubles %s \n", rangeStr, sDoubles)
		for _, double := range doubles {
			sum += double
		}
	}
	fmt.Printf("Sum: %d", sum)
}
