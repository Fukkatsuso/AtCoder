package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	t_begin := -1

	for i := len(s) - 1; i >= 0; i-- {
		k := i
		for j := len(t) - 1; j >= 0; j-- {
			if !(t[j] == s[k]) && !(s[k] == '?') {
				break
			}
			k--
			if k < 0 {
				break
			}
		}
		if k == i-len(t) {
			t_begin = k + 1
			break
		}
	}

	if t_begin < 0 {
		fmt.Println("UNRESTORABLE")
		return
	}
	for i := 0; i < len(s); i++ {
		if t_begin <= i && i < t_begin+len(t) {
			fmt.Printf("%s", string(t[i-t_begin]))
		} else if s[i] == '?' {
			fmt.Printf("a")
		} else {
			fmt.Printf("%s", string(s[i]))
		}
	}
	fmt.Printf("\n")
}
