package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

var chunkSize = 4

// generate index
func generateIndex(str string, tablesize int) uint {
	slices := divChunk(str)
	binSlice := convertStringToBin(slices)
	index := exOr(binSlice)
	size := index % uint(tablesize)
	return size
}

// xor the 4 bit chunks
func exOr(binaryVal []string) uint {
	var res uint
	for _, b := range binaryVal {
		res = res ^ binStringToDecimal(b)
	}
	return res
}

// divide the song name into 4-byte chunks
func divChunk(str string) []string {

	slices := []string{}
	lastIndex := 0
	lastI := 0
	for i := range str {
		if i-lastIndex > chunkSize {
			slices = append(slices, str[lastIndex:lastI])
			lastIndex = lastI
		}
		lastI = i
	}
	// handle leftovers at the end
	if len(str)-lastIndex > chunkSize {
		slices = append(slices, str[lastIndex:lastIndex+chunkSize], str[lastIndex+chunkSize:])
	} else {
		slices = append(slices, str[lastIndex:])
	}
	return slices
}

// convert chunks to binary, odd chunks reversed and binary
func convertStringToBin(slices []string) []string {

	//slices := divChunk(str)
	var bin []string
	var binaryVal string

	for i, s := range slices {
		// pad zeroes if length < chunksize
		for len(s) < chunkSize {
			s = s + "0"
		}
		hexValue := hex.EncodeToString([]byte(s)) // string to hex value

		if i%2 == 0 {
			binaryVal = reverseAndConvertToBin(hexValue) // reverse the odd chunks and convert to hex to binary
		} else {
			binaryVal = parseHexToBin(hexValue) // convert other chunks to hex to binary
		}
		bin = append(bin, binaryVal)
	}
	return bin
}

func parseHexToBin(s string) string {
	hexToBin, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return "error"
	}
	return fmt.Sprintf("%032b", hexToBin)
}

func reverseAndConvertToBin(hex string) string {
	var reverseBin string
	hexToBin := parseHexToBin(hex)
	for _, v := range hexToBin {
		reverseBin = string(v) + reverseBin
	}
	return reverseBin
}

func binStringToDecimal(b string) uint {
	u64, err := strconv.ParseUint(b, 2, 64) //binary string to decimal
	if err != nil {
		fmt.Println(err)
	}
	return uint(u64)
}
