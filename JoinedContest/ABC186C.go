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

func haveSeven(s string) bool {
	for _, b := range s {
		if b == '7' {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	ans := 0
	for i := 1; i <= n; i++ {
		// 10
		ten := fmt.Sprintf("%d", i)
		// 8
		eight := fmt.Sprintf("%o", i)
		if !haveSeven(ten) && !haveSeven(eight) {
			ans++
		}
	}

	puts(ans)
}
