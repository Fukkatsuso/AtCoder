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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// エラトステネスのふるい
// 1からnまでの数で
// 素数かそうでないかを全て判定し
// 真理値配列を返す
func sieve(n int) []bool {
	s := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		s[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if !s[i] {
			continue
		}
		for j := 2; i*j <= n; j++ {
			s[i*j] = false
		}
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	s := sieve(n - 1)
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] {
			ans++
		}
	}
	puts(ans)
}
