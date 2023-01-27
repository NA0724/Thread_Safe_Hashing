package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var input Input
var table map[string]string

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
	if isRandom(mode) {
		randomOperation(input)
	} else if isManual(mode) {
		manualOperation(input)
	} else {
		fmt.Println("Your input was invalid")
		return
	}
}

func isRandom(mode string) bool {
	return strings.Compare(mode, "random") == 0 || strings.Compare(mode, "Random") == 0
}

func isManual(mode string) bool {
	return strings.Compare(mode, "manual") == 0 || strings.Compare(mode, "Manual") == 0
}
