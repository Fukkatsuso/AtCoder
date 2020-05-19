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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// bが非負の既約分数
func cut(a, b int) (int, int) {
	g := gcd(abs(a), abs(b))
	if b < 0 {
		return -a / g, -b / g
	}
	return a / g, b / g
}

func modPow(a, n, mod int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = (ret * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return ret
}

type Pair struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n := nextInt()

	zero := 0
	m := map[Pair]Pair{}
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		if a == 0 && b == 0 {
			zero++
			continue
		} else if a == 0 {
			p := m[Pair{0, 1}]
			p.a++
			m[Pair{0, 1}] = p
			continue
		} else if b == 0 {
			p := m[Pair{0, 1}]
			p.b++
			m[Pair{0, 1}] = p
			continue
		}

		a, b = cut(a, b)
		if a < 0 {
			p := m[Pair{b, -a}]
			p.a++
			m[Pair{b, -a}] = p
		} else {
			p := m[Pair{a, b}]
			p.b++
			m[Pair{a, b}] = p
		}
	}

	ans := 1
	for _, p := range m {
		a := (modPow(2, p.a, mod) + modPow(2, p.b, mod) + mod - 1) % mod
		ans = (ans * a) % mod
	}
	ans = (ans + zero - 1 + mod) % mod

	puts(ans)
}
