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
	} else {
		// there is a collision, get into linked-list mode
		starting_node := h.Data[index]
		fmt.Println(starting_node.next != nil)
		for ; starting_node.next != nil; starting_node = starting_node.next {
			if starting_node.key == key && starting_node.value == value {
				fmt.Println("already exists")
				return
			}
		}
		if starting_node.next == nil {
			if starting_node.key == key && starting_node.value == value {
				fmt.Println("here2")
				fmt.Println("already exists")
				return
			}
		}
		starting_node.next = &Node{key: key, value: value}
	}
	fmt.Println(h)
}

func (h *HashMap) Get(threadid int, key string, index uint) []string {
	var values []string
	if h.Data[index] != nil {
		// key is on this index, but might be somewhere in linked list
		starting_node := h.Data[index]
		for ; ; starting_node = starting_node.next {
			if starting_node.key == key {
				// key matched
				values = append(values, starting_node.value)
			}
			if starting_node.next == nil {
				break
			}
		}
		if len(values) > 0 {
			return values
		}
	}
	// key does not exists
	values = append(values, "does not exist")
	return values
}

func (h *HashMap) Delete(threadid int, key string, value string, index uint) {
	flag := false
	var prev *Node
	// key is on this index, but might be somewhere in linked list
	head := h.Data[index]
	curr_node := head
	for curr_node != nil {
		if curr_node.key == key && curr_node.value == value {
			if prev == nil {
				head = curr_node.next
			} else {
				prev.next = curr_node.next
			}
			flag = true
		} else {
			prev = curr_node
		}
		curr_node = curr_node.next
	}
	if flag {
		fmt.Println("deleted from table ", key, value)
	} else {
		fmt.Println(key, value, " does not exist in table")
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
					fmt.Println(NewDict(tablesize).Get(threadid, song, index))
				} else if val > 0.7 || val <= 0.9 {
					NewDict(tablesize).Insert(threadid, song, socket, index)
				} else {
					NewDict(tablesize).Delete(threadid, song, socket, index)
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
		runRandomOperations(i, operations, entries)
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
