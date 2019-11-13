package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	sub := [4]string{"dream", "dreamer", "erase", "eraser"}
	n := len(s)
	for i := n - 1; i > 0; {
		for j := 0; j < 4; j++ {
			if strings.HasSuffix(s, sub[j]) {
				k := len(sub[j])
				s = s[:(len(s) - k)]
				i -= k
				break
			} else if j == 3 {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
