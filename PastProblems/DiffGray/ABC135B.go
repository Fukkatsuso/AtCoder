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
	n := nextInt()
	p := nextInts(n)
	cnt := 0
	for i := 0; i < n; i++ {
		if p[i] != i+1 {
			cnt++
		}
	}
	if cnt == 0 || cnt == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
