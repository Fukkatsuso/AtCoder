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
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] && s[i] != '*' && s[n-1-i] != '*' {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}