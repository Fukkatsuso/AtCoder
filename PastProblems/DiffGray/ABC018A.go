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

func p(a, b, c int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := nextInt(), nextInt(), nextInt()
	switch {
	case a > b && b > c:
		p(1, 2, 3)
	case a > c && c > b:
		p(1, 3, 2)
	case b > a && a > c:
		p(2, 1, 3)
	case a < c && c < b:
		p(3, 1, 2)
	case c > a && a > b:
		p(2, 3, 1)
	case c > b && b > a:
		p(3, 2, 1)
	}
}
