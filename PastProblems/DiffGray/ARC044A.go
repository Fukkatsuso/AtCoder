package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	n := int(math.Sqrt(float64(x)))
	for i := 2; i <= n; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func digitsSum(n int) int {
	ret := 0
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return ret
}

func likelyPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n%10 == 5 || n&1 == 0 || digitsSum(n)%3 == 0 {
		return false
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	if isPrime(n) || likelyPrime(n) {
		puts("Prime")
	} else {
		puts("Not Prime")
	}
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
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

func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
	}
	return slice
}

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
