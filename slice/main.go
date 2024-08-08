package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

// Slice function
func Slice(s []string, n ...int) []string {
	l := len(n)
	ls := len(s)
	if l == 1 {
		start := n[0]
		if start >= 0 {
			return s[start:]
		} else {
			return s[ls+start:]
		}
	}

	if l > 1 {

		start := n[0]
		end := n[1]

		// check if start and end are both positve
		if start >= 0 && end >= 0 {
			if start > end {
				return []string(nil)
			}
			if end > ls {
				end = ls
			}
			return s[start:end]
			// check if start and end are 0
		} else if start < 0 && end < 0 {
			if start > end {
				return []string(nil)
			}
			start = ls + start
			if start < 0 {
				return []string(nil)
			}
			end = ls + end
			if end < 0 {
				return []string(nil)
			}
			return s[start:end]
		}
	}
	return []string{}
}
