// 解説AC
// '0'を1桁の整数に含めるのを忘れてWA

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
}

// nの上位d桁目の数
func digitNum(n, d int) int {
	s := strconv.Itoa(n)
	return int(s[d-1] - '0')
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	s, c := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		s[i], c[i] = nextInt(), nextInt()
	}

	ans := -1
	iMax := pow(10, n)
	iBegin := pow(10, n-1)
	if iBegin == 1 {
		iBegin = 0
	}
	for i := iBegin; i < iMax; i++ {
		ok := true
		for j := 0; j < m; j++ {
			if digitNum(i, s[j]) != c[j] {
				ok = false
				break
			}
		}
		if ok {
			ans = i
			break
		}
	}
	puts(ans)
	wt.Flush()
}

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
