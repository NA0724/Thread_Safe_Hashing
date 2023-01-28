package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var input Input
var entries []string

func main() {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()

		if len(text) != 0 {
			fmt.Println(text)
			//exclude lines that are comments
			if strings.Contains(text, "#") {
				continue
			}
			input = append(input, text)
		} else {
			break
		}
	}
	fmt.Println()
	mode := input[0]
	processInput(mode, input)
}

// process the input from the console
func processInput(mode string, input []string) {
	if isRandom(mode) {
		randomOperation(input)
	} else if isManual(mode) {
		manualOperation(input)
	} else {
		fmt.Println("Your input was invalid")
		return
	}
}

func manualOperation(input []string) {
	fmt.Println("Considering thread id = 0 for Manual mode ")
	for _, s := range input[1:] {
		entries = append(entries, s)
	}
	doManualOperation(entries)
}

func randomOperation(input []string) {
	var threads, operations int
	var err error
	// check if the number of threads and number of operations entered in an integer value
	if threads, err = strconv.Atoi(input[1]); err == nil {
		fmt.Println("Thread=", threads)
	} else {
		fmt.Println("Invalid format for number of threads. Please enter correct format.")
		return
	}
	if operations, err = strconv.Atoi(input[2]); err == nil {
		fmt.Println("Operations=", operations)
	} else {
		fmt.Println("Invalid format for number of opertions. Please enter correct format.")
		return
	}
	for _, line := range input[3:] {
		entries = append(entries, line)
	}
	runThreads(threads, operations, entries)
}

func isRandom(mode string) bool {
	return strings.Compare(mode, "random") == 0 || strings.Compare(mode, "Random") == 0
}

func isManual(mode string) bool {
	return strings.Compare(mode, "manual") == 0 || strings.Compare(mode, "Manual") == 0
}
