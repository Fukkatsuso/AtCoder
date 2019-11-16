// WA
package main

import (
	"fmt"
)

func main() {
	const mod int = 1000000007
	var x, y int
	fmt.Scan(&x, &y)

	a := 2*y - x
	b := 2*x - y
	if a < 0 || b < 0 {
		fmt.Println(0)
		return
	}
	if a%3 != 0 || b%3 != 0 {
		fmt.Println(0)
		return
	}
	a /= 3
	b /= 3
	if b == 0 {
		a, b = b, a
	}
	var ans uint64 = 1
	var div uint64 = uint64(b)
	for i := 1; i <= b; i++ {
		if uint64(a+i)%div == 0 {
			div--
			continue
		}
		ans = (ans * uint64(a+i))
	}
	fmt.Println(ans)
	for i := uint64(2); i <= div; i++ {
		ans /= uint64(i)
	}
	fmt.Println(ans)
}
