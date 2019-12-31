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
	n := 5
	a := nextInts(n)
	k := nextInt()
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[j]-a[i] > k {
				fmt.Println(":(")
				return
			}
		}
	}
	fmt.Println("Yay!")
}
