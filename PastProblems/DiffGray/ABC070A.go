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

func reverseString(s string) string {
	var t string
	for i := len(s) - 1; i >= 0; i-- {
		t += string(s[i])
	}
	return t
}

func main() {
	sc.Split(bufio.ScanWords)
	n := next()
	if n == reverseString(n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
