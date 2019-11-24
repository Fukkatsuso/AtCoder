package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	var ans string

	for _, c := range s {
		ans += string((int(c)-65+n)%26 + 65)
	}
	fmt.Println(ans)
}
