package main

import (
	"bufio"
	"fmt"
	"math"
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

func isPrime(x int) bool {
	y := int(math.Sqrt(float64(x)))
	for i := 2; i <= y; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	x := nextInt()
	for ; !isPrime(x); x++ {
	}
	fmt.Println(x)
}
