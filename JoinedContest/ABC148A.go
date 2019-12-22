package main

import (
	"fmt"
)

func main() {
	var a, b int
	var ans int
	fmt.Scan(&a, &b)
	if a*b == 2 {
		ans = 3
	} else if a*b == 3 {
		ans = 2
	} else {
		ans = 1
	}
	fmt.Println(ans)
}
