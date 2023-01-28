package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

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
		fmt.Println("Thread id: ", threadid, " put song <", key, "> at socket <", value, "> in hash table at index ", index)
	} else {
		// there is a collision, get into linked-list mode
		starting_node := h.Data[index]
		for ; starting_node.next != nil; starting_node = starting_node.next {
			if starting_node.key == key && starting_node.value == value {
				fmt.Println("Thread id: ", threadid, " put song <", key, "> at socket <", value, "> already in hash table at index ", index)
				return
			}
		}
		if starting_node.next == nil {
			if starting_node.key == key && starting_node.value == value {
				fmt.Println("Thread id: ", threadid, " put song <", key, "> at socket <", value, "> already in hash table at index ", index)
				return
			}
		}
		starting_node.next = &Node{key: key, value: value}
		fmt.Println("Thread id: ", threadid, " put song <", key, "> at socket <", value, "> in hash table at index ", index)
	}
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
			fmt.Println("Thread id: ", threadid, " get song <", key, "> can be downloaded from sockets ")
			return values
		}
	}
	// key does not exists
	values = append(values, "")
	fmt.Println("Thread id: ", threadid, " get song <", key, "> is not in the hash table")
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
		fmt.Println("Thread id: ", threadid, " delete song <", key, "> at socket <", value, "> from hash table")
	} else {
		fmt.Println("Thread id: ", threadid, " delete song <", key, "> is not in hash table")
	}

}

func runRandomOperations(wg *sync.WaitGroup, mu *sync.Mutex, threadid, operations int, entries []string) {

	fmt.Println("Inside runRandomOperations")
	var count = 1
	noOfEntries := len(entries)
	tablesize := rehash(float64(noOfEntries))
	hashtable := NewDict(tablesize)
	for j := 0; j < operations; j++ {
		if count == operations+1 {
			break
		} else {
			count++
			if j == noOfEntries {
				j = 0
			}
			//choosing operation
			val := rand.Float32()

			//fetch song and socket from entries
			s := entries[j]
			str := strings.Split(s, " ")
			song := fmt.Sprint(str[:len(str)-1])
			song = strings.Replace(song, ",", "", -1)
			socket := str[len(str)-1]

			index := generateIndex(song, tablesize)
			if val <= 0.7 {
				fmt.Println(hashtable.Get(threadid, song, index))
			} else if val > 0.7 && val <= 0.9 {
				hashtable.Insert(threadid, song, socket, index)
			} else {
				hashtable.Delete(threadid, song, socket, index)
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
	time.Sleep(time.Millisecond * 100)

}

func runThreads(threads, operations int, entries []string) {
	// implement go routoines
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= threads; i++ {
		wg.Add(threads)

		go runRandomOperations(&wg, &mu, i, operations, entries)

		time.Sleep(time.Millisecond * 100)

		wg.Done()
	}
	wg.Wait()
}

// Manual mode operations
func doManualOperation(entries []string) {
	noOfEntries := len(entries)
	tablesize := rehash(float64(noOfEntries))
	hashtable := NewDict(tablesize)

	for _, str := range entries {
		s := strings.Split(str, " ")
		op := s[0]
		socket := s[len(s)-1]
		song := fmt.Sprint(s[1 : len(s)-1])
		song = strings.Replace(song, ",", "", -1)
		index := generateIndex(song, tablesize)
		if strings.Compare(op, "put") == 0 || strings.Compare(op, "Put") == 0 || strings.Compare(op, "PUT") == 0 {
			hashtable.Insert(0, song, socket, index)
		} else if strings.Compare(op, "get") == 0 || strings.Compare(op, "Get") == 0 || strings.Compare(op, "GET") == 0 {
			fmt.Println(hashtable.Get(0, song, index))
		} else if strings.Compare(op, "delete") == 0 || strings.Compare(op, "Delete") == 0 || strings.Compare(op, "DELETE") == 0 {
			hashtable.Delete(0, song, socket, index)
		} else {
			fmt.Println("Invalid Operation")
			return
		}
	}
}
