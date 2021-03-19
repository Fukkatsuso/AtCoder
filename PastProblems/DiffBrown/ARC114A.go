// 解説AC

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	x := getInts(n)

	primes := make([]int, 0)
	for i := 2; i <= 50; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	m := len(primes)
	ans := 1 << 60
	for bit := 1; bit < (1 << m); bit++ {
		ok := true
		for _, xi := range x {
			// primesのある項pでxiが割り切れるか
			okx := false
			for j, p := range primes {
				if (bit>>j)&1 == 1 && xi%p == 0 {
					okx = true
					break
				}
			}
			ok = ok && okx
		}
		if !ok {
			continue
		}

		y := 1
		for j, p := range primes {
			if (bit>>j)&1 == 1 {
				y *= p
			}
		}
		ans = min(ans, y)
	}

	puts(ans)
}
