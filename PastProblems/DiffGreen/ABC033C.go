package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	terms := strings.Split(s, "+")

	ans := 0
	for i := range terms {
		var isZero bool
		for j := 0; j < len(terms[i]); j += 2 {
			isZero = (terms[i][j] == '0')
			if isZero {
				break
			}
		}
		if !isZero {
			ans++
		}
	}
	fmt.Println(ans)
}
