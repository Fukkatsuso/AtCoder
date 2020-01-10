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

func div8(n int) bool {
	if n%2 == 0 {
		return false
	}
	div := 2
	for i := 3; i < n; i++ {
		if n%i == 0 {
			div++
		}
	}
	if div == 8 {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ans := 0
	for i := 105; i <= n; i++ {
		if div8(i) {
			ans++
		}
	}
	fmt.Println(ans)
}
