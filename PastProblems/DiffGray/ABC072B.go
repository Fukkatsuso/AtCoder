package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := ""
	for i := 0; i < len(s); i += 2 {
		ans += string(s[i])
	}
	fmt.Println(ans)
}
