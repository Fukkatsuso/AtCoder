package main

import (
	"fmt"
)

func main() {
	var t1, t2 int64
	var a1, a2 int64
	var b1, b2 int64
	fmt.Scan(&t1, &t2)
	fmt.Scan(&a1, &a2)
	fmt.Scan(&b1, &b2)

	p := (a1 - b1) * t1
	q := (a2 - b2) * t2
	if p > 0 {
		p *= -1
		q *= -1
	}

	if p+q == 0 {
		fmt.Println("infinity")
		return
	}

	ans := int64(0)

	if p+q < 0 {
		ans = 0
	} else {
		s := (-p) / (p + q)
		t := (-p) % (p + q)
		if t != 0 {
			ans = s*2 + 1
		} else {
			ans = s * 2
		}
	}

	fmt.Println(ans)
}
