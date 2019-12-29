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

func reverse(s string, from, to int) string {
	runes := []rune(s)
	for i, j := from, to; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	length := len(s)
	n := nextInt()
	for i := 0; i < n; i++ {
		l, r := nextInt(), nextInt()
		s = s[0:l-1] + reverse(s, l-1, r-1)[l-1:r] + s[r:length]
	}
	fmt.Println(s)
}
