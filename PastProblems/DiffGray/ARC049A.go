package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	for i := 0; i <= len(s); i++ {
		if i == a || i == b || i == c || i == d {
			fmt.Printf("\"")
		}
		if i < len(s) {
			fmt.Printf("%c", s[i])
		}
	}
	fmt.Printf("\n")
}
