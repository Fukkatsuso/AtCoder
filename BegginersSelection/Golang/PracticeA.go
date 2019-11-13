package main

import "fmt"

func main() {
	var a, b, c int
	var s string
	fmt.Scan(&a, &b, &c)
	fmt.Scan(&s)
	fmt.Println(a+b+c, s)
}
