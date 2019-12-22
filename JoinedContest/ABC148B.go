package main

import (
	"fmt"
)

func main() {
	var n int
	var s, t string
	var ans string
	fmt.Scan(&n, &s, &t)
	for i := 0; i < n; i++ {
		ans += string(s[i])
		ans += string(t[i])
	}
	fmt.Println(ans)
}
