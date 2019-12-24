package main

import (
	"fmt"
)

func fillSlice(x []int, y int) {
	for i := range x {
		x[i] = y
	}
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	x := make([]int, 26)
	y := make([]int, 26)
	fillSlice(x, -1)
	fillSlice(y, -1)

	ok := true
	for i := 0; i < len(s); i++ {
		si := int(s[i]) - 97
		ti := int(t[i]) - 97
		if x[si] < 0 {
			x[si] = ti
		}
		if y[ti] < 0 {
			y[ti] = si
		}
		if (x[si] != ti) || (y[ti] != si) {
			ok = false
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
