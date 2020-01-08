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

func secondMin(a, b, c int) bool {
	if a < b && b < c {
		return true
	}
	if a > b && b > c {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p := nextInts(n)

	ans := 0
	for i := 1; i < n-1; i++ {
		if secondMin(p[i-1], p[i], p[i+1]) {
			ans++
		}
	}
	fmt.Println(ans)
}
