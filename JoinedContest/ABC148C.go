package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(lcm(a, b))
}
