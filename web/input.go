package main

import (
	"fmt"
	"strconv"
)

var entries []string

func manualOperation(input []string) {
	fmt.Println("input from user", input)
	for _, s := range input[1:] {
		entries = append(entries, s)
	} //TODO: do manual operation
	fmt.Println("**************************")
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

	//runThreads(threads, operations, entries)
	//for testing
	size := rehash(float64(len(entries)))
	hashtable := NewDict(size)
	hashtable.Insert(1, "Listen to the music", "http://foo.com:54321", generateIndex("Listen to the music", size))
	hashtable.Insert(1, "Listen to the music", "http://foo.com:54321", generateIndex("Listen to the music", size))
	hashtable.Insert(1, "Time to say goodbye", "http://bar.com:12345", generateIndex("Time to say goodbye", size))
	hashtable.Insert(1, "Listen to the music", "http://ijk.com:22222", generateIndex("Listen to the music", size))
	fmt.Println(hashtable)
	fmt.Println("===========================================================================================")
	if str, ok := hashtable.Get(1, "Changing partner", generateIndex("Listen to the music", size)); ok {
		fmt.Println(str)
	}
	if xyz, ok := hashtable.Get(1, "Listen to the music", generateIndex("Listen to the music", size)); ok {
		fmt.Println(xyz)
	}
	hashtable.Delete(1, "Listen to the music", generateIndex("Listen to the music", size))
	hashtable.Delete(1, "Changing partner", generateIndex("Listen to the music", size))
	fmt.Println(hashtable)
}
