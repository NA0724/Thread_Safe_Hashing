package main

import "math"

var power = 3

/*The initial default table size is 7 (the closest prime number around 23).
If the total numbers of entries is greater than roughly 110% of the table size,
you need to do rehash by doubling the table size and make the table size a closest
prime number (e.g., 17 is the closest prime for 24). If the total numbers of the entries
is less than roughly 40% of the table size, you need to do rehash by half the table size
to the closest prime number */

func rehash(noOfEntries float64) int {
	tablelen := tableSize(power)
	m1 := math.Floor(1.1 * float64(tablelen)) // 110% of table size
	m2 := math.Floor(0.4 * float64(tablelen)) // 40% of table size
	if noOfEntries > m1 {
		tablelen = tableSize(power + 1)
	} else if noOfEntries < m2 {
		tablelen = tableSize(power - 1)
	}
	return tablelen
}

func tableSize(pow int) int {
	size := findNearPrime(int(math.Pow(2, float64(pow))))
	return size
}

func isPrime(n int) bool {
	c := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	} else {
		return false
	}
}

func findNearPrime(num int) int {
	diff1 := 0
	diff2 := 0
	num1 := 0
	num2 := 0

	for i := num; ; i++ { //No end limit as when prime will be found we will break the loop.
		if isPrime(i) {
			diff1 = i - num
			num1 = i
			break
		}
	}
	for j := num; ; j-- { //No end limit as when prime will be found we will break the loop.
		if isPrime(j) {
			diff2 = num - j
			num2 = j
			break
		}
	}
	if diff1 < diff2 { //Nearest Prime number will have least difference from given number.
		return num1
	} else if diff2 < diff1 {
		return num2
	} else {
		return num1
	}

}
