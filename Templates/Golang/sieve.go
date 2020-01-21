package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	n := 100
	s := sieve(n)
	for i := 1; i <= n; i++ {
		if s[i] {
			puts(i)
		}
	}
	wt.Flush()
}

var (
	wt = bufio.NewWriter(os.Stdout)
)

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
