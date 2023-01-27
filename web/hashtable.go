package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

var m sync.Mutex

type Node struct {
	key   string
	value string
	next  *Node
}

type HashMap struct {
	Data []*Node
}

func NewDict(size int) *HashMap {
	return &HashMap{Data: make([]*Node, size)}
}

func (n *Node) String() string {
	return fmt.Sprintf("<Key: %s, Value: %s>\n", n.key, n.value)
}

func (h *HashMap) String() string {
	var output bytes.Buffer
	fmt.Fprintln(&output, "{")
	for _, n := range h.Data {
		if n != nil {
			fmt.Fprintf(&output, "\t%s: %s\n", n.key, n.value)
			for node := n.next; node != nil; node = node.next {
				fmt.Fprintf(&output, "\t%s: %s\n", node.key, node.value)
			}
		}
	}

	fmt.Fprintln(&output, "}")

	return output.String()
}

func (h *HashMap) Insert(threadid int, key string, value string, index uint) {
	if h.Data[index] == nil {
		// index is empty, go ahead and insert
		h.Data[index] = &Node{key: key, value: value}
		fmt.Println("Thread id: ", threadid, " put <", key, "> at socket <", value, "> in hash table at index ", index)
	} else {
		// there is a collision, get into linked-list mode
		node := h.Data[index]
		if node.key == key && node.value == value {
			fmt.Println("Thread id: ", threadid, " put <", key, "> at socket <", value, "> already exists in hash table at index ", index)
		} else {
			starting_node := node
			for ; starting_node.next != nil; starting_node = starting_node.next {
				fmt.Println(starting_node.key == key)
				fmt.Println(starting_node.value)
				if starting_node.key == key {
					// the key exists, its a modifying operation
					starting_node.value = value
					fmt.Println("Thread id: ", threadid, " put <", key, "> at socket <", value, "> in hash table at index ", index)
					return
				}
			}
			starting_node.next = &Node{key: key, value: value}
		}
	}
}

func (h *HashMap) Get(threadid int, key string, index uint) (string, bool) {
	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in linked list
		starting_node := h.Data[index]
		for ; ; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				fmt.Print("Thread id: ", threadid, " get song <", key, "> at socket from hash table:")
				return starting_node.value, true
			}
			if starting_node.next == nil {
				break
			}
		}
	}
	// key does not exists
	fmt.Println("Thread id: ", threadid, " get song <", key, ">does not exist in hash table")
	return "", false
}

func (h *HashMap) Delete(threadid int, key string, index uint) {

	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in linked list
		starting_node := h.Data[index]
		value := starting_node.value
		for ; ; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				starting_node.next = starting_node.next.next
				fmt.Println("Thread id: ", threadid, " delete song <", key, "> at socket <", value, "> from hash table")
			} else {
				fmt.Println("Thread id: ", threadid, " delete song <", key, "> does not exist in hash table")
				return
			}
			if starting_node.next == nil {
				break
			}
			starting_node = starting_node.next
		}
	}
}

func runRandomOperations(threadid, operations int, entries []string) {
	fmt.Println("Inside runRandomOperations")
	m.Lock()
	var count = 1
	noOfEntries := len(entries)
	tablesize := rehash(float64(noOfEntries))
	for j := 0; j < operations; j++ {
		if count == operations {
			break
		} else {
			count++
			if j == noOfEntries {
				j = 0
			}
			//choosing operation
			val := rand.Float32()
			s := entries[j]
			if strings.Contains(s, ",") {
				str := strings.Split(s, ",")
				song := str[0]
				socket := str[1]
				index := generateIndex(s, tablesize)
				if val > 0.0 || val <= 0.7 {
					if getres, ok := NewDict(tablesize).Get(threadid, song, index); ok {
						fmt.Println(getres)
					}
				} else if val > 0.7 || val <= 0.9 {
					NewDict(tablesize).Insert(threadid, song, socket, index)
				} else {
					NewDict(tablesize).Delete(threadid, song, index)
				}
			} else {
				fmt.Println("Song socket input format is invalid, please enter correct format")
				return
			}
		}
	}
	m.Unlock()
}

func runThreads(threads, operations int, entries []string) {
	for i := threads; i > 0; i-- {
		fmt.Println("Inside run threads function")
		go runRandomOperations(i, operations, entries)
	}

}

/*
func doManualOperation(entries []string) {
	if strings.Compare(op, "put") == 0 || strings.Compare(op, "Put") == 0 || strings.Compare(op, "PUT") == 0 {
		put(song, socket, table)
	} else if strings.Compare(op, "get") == 0 || strings.Compare(op, "Get") == 0 || strings.Compare(op, "GET") == 0 {
		NewDict(tablesize).Get(song, index)
	} else if strings.Compare(op, "delete") == 0 || strings.Compare(op, "Delete") == 0 || strings.Compare(op, "DELETE") == 0 {
		put(song, socket, table)
	} else {
		fmt.Println("Invalid Operation")
		return
	}
}*/
