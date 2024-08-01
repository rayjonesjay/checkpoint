package main

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
	fmt.Println(HashCode("14 Avril 1973"))

	fmt.Println("------------------------TESTS-------------------------------")
	table := []string{"Z", "Hi!", "BB198365", "sabito", "14 Avril 1912", "zyx987bca", "		pool-2020", "965truma747", " Mercedes-AMG GT"}
	for _, s := range table {
		challenge.Function("HashCode", HashCode, solutions.HashCode, s)
	}
	fmt.Println("ok")
}

func HashCode(s string) string {
	hashed := make([]rune, len(s))
	size := len(s)
	for i, char := range s {
		hashedChar := (size + int(char)) % 127
		if isPrintable(rune(hashedChar)) {
			hashed[i] = rune(hashedChar)
		} else {
			hashed[i] = rune(hashedChar + 33)
		}
	}
	return string(hashed)
}

func isPrintable(c rune) bool {
	if c < 32 || c >= 127 {
		return false
	}
	return true
}
