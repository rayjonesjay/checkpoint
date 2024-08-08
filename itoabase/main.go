package main

import (
	"fmt"
	"math"

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

	if value == 0 {
		return "0"
	}
	
	result := ""
	
	isNegative := value < 0
	
	if isNegative {
		
		// -1 << 31 is same as math.MinInt32  -1 << 63 is same as math.MinInt64 
		if value <= -1 << 31 {
			// take care of the last digit
			lastDigit := -(value%base)
			value = -(value/base)
			return "-" + ItoaBase(value, base) + string("0123456789ABCDEF"[lastDigit])
		}
		
		value = -value
	}

	s := "0123456789ABCDEF"

	for value > 0 {
		rem := value % base
		result = string(s[rem]) + result
		value /= base
	}

	if isNegative {
		return  "-" + result 
	}
	if result == "" {
		return "0"
	}

	return result
}


