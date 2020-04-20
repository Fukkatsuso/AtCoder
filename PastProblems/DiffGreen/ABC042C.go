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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func canPay(usable []bool, money int) bool {
	for money > 0 {
		if !usable[money%10] {
			return false
		}
		money /= 10
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	usable := make([]bool, 10)
	for i := range usable {
		usable[i] = true
	}
	for i := 0; i < k; i++ {
		d := nextInt()
		usable[d] = false
	}

	ans := n
	for ; !canPay(usable, ans); ans++ {
	}

	puts(ans)
}
