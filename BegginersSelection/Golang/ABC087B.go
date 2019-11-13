package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	ans := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			k := (x - 500*i - 100*j) / 50
			if 0 <= k && k <= c {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
