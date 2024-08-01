package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	fmt.Println(ItoaBase(255, 16))  // "FF"
	fmt.Println(ItoaBase(-255, 16)) // "-FF"
	fmt.Println(ItoaBase(10, 2))    // "1010"
	fmt.Println(ItoaBase(42, 4))    // "222"
	fmt.Println(ItoaBase(0, 10))    // "0"

	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MaxInt, base)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MinInt, base)
	}
}

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return "Wrong Base"
	}
	result := ""
	// check if value is negative
	if value < 0 {
		result = "-" + ItoaBase(-value, base)
	}

	s := "0123456789ABCDEF"

	for value > 0 {
		rem := value % base
		result = string(s[rem]) + result
		value /= base
	}

	if result == "" {
		return "0"
	}

	return result
}

func PrintMemory(arr [10]byte) {
	for _, b := range arr {
		// Extract the upper 4 bits and the lower 4 bits
		upper := (b >> 4) & 0xF
		lower := b & 0xF

		// Convert these to hexadecimal characters using the digits string
		fmt.Printf("%X%X ", upper, lower)
	}
	fmt.Println()
}
