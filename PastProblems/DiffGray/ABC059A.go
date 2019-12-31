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

func toUpper(b byte) byte {
	return byte(b - 32)
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, c := next(), next(), next()
	fmt.Printf("%c%c%c\n", toUpper(a[0]), toUpper(b[0]), toUpper(c[0]))
}
