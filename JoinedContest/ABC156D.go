package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func modInv(a, mod int) int {
	b := mod
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		u -= t * v
		a, b = b, a
		u, v = v, u
	}
	u = u % mod
	if u < 0 {
		u += mod
	}
	return u
}

func com(n, k, mod int) int {
	a := 1
	for i := n - k + 1; i <= n; i++ {
		a = (a * i) % mod
	}
	for i := 1; i <= k; i++ {
		a = (a * modInv(i, mod)) % mod
	}
	return a % mod
}

func main() {
	sc.Split(bufio.ScanWords)
	const mod int = 1000000007
	n, a, b := nextInt(), nextInt(), nextInt()
	ans := modPow(2, n, mod) - 1 - (com(n, a, mod)+com(n, b, mod))%mod
	if ans < 0 {
		ans += mod
	}
	puts(ans)
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
