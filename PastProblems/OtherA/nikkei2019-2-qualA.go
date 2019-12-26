package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ans := n / 2
	if n%2 == 0 {
		ans--
	}
	fmt.Println(ans)
}
