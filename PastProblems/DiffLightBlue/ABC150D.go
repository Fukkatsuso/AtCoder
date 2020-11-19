// 解説AC

package main

import (
	"bufio"
	"fmt"
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func pow2(x int) int {
	res := 0
	for x > 0 && x%2 == 0 {
		x /= 2
		res++
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	a := getInts(n)

	// a[0]が2で割り切れる回数
	// 2で割り切れる回数が全てのa[i]で等しいかを確認
	p := pow2(a[0])
	if p == 0 {
		puts(0)
		return
	}
	for i := 1; i < n; i++ {
		if pow2(a[i]) != p {
			puts(0)
			return
		}
	}

	// a[i]/2の最小公倍数
	// mを超えると答えは0
	t := 1
	for i := range a {
		t = lcm(t, a[i]/2)
		if t > m {
			puts(0)
			return
		}
	}

	// 条件を満たすxの個数をカウント
	ans := 0
	for x := 0; t*(2*x+1) <= m; x++ {
		ans++
	}
	puts(ans)
}
