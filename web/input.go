package main

import (
	"fmt"
	"strconv"
)

var entries []string

func manualOperation(input []string) {
	fmt.Println("Considering thread id = 0 for Manual mode ")
	for _, s := range input[1:] {
		entries = append(entries, s)
	}
	doManualOperation(entries)
	fmt.Println("*************   Manual mode ended *************")
}

func randomOperation(input []string) {
	var threads, operations int
	var err error
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
	fmt.Println("*************   Random mode ended *************")
}
