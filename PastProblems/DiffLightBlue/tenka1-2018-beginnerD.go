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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func putSlice(s []int) {
	n := len(s)
	putf("%d ", n)
	for i := 0; i < n-1; i++ {
		putf("%d ", s[i])
	}
	if n > 0 {
		puts(s[n-1])
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	k := 0
	for i := 1; i*(i-1) <= 2*n; i++ {
		if i*(i-1) == 2*n {
			k = i
		}
	}
	if k == 0 {
		puts("No")
		return
	}

	puts("Yes")
	puts(k)

	ans := make([][]int, k)
	ans[0] = make([]int, 0)
	for i := 1; i < k; i++ {
		ans[0] = append(ans[0], (i*i+i)/2)
	}
	for i := 1; i < k; i++ {
		ans[i] = make([]int, 0)
		s := 1
		if i > 1 {
			s = ans[0][i-2] + 1
		}
		// ヨコ
		for s < ans[0][i-1] {
			ans[i] = append(ans[i], s)
			s++
		}
		// タテ
		for t, j := s, i; t <= n; j++ {
			ans[i] = append(ans[i], t)
			t += j
		}
	}

	for i := range ans {
		putSlice(ans[i])
	}
}
