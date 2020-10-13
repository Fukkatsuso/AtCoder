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

func lowerBound(a []int, key int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] < key {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

type Query struct {
	op   string
	n, m int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	q, l := nextInt(), nextInt()
	query := make([]Query, q)
	for i := 0; i < q; i++ {
		op := next()
		qi := Query{op: op}
		if op == "Push" {
			qi.n, qi.m = nextInt(), nextInt()
		} else if op == "Pop" {
			qi.n = nextInt()
		}
		query[i] = qi
	}

	// 一度のPushを1層追加とする
	// num[i]: 下からi層目の数字
	// dp[i]: 下から層iまでの累積和
	num, dp := make([]int, 1), make([]int, 1)
	num[0], dp[0] = 1<<30, 0
	for _, qi := range query {
		// 層数
		n := len(num)
		switch qi.op {
		case "Push":
			if dp[n-1]+qi.n > l {
				puts("FULL")
				return
			}
			num = append(num, qi.m)
			dp = append(dp, dp[n-1]+qi.n)
		case "Pop":
			if qi.n > dp[n-1] {
				puts("EMPTY")
				return
			}
			res := dp[n-1] - qi.n
			idx := lowerBound(dp, res)
			num, dp = num[:idx+1], dp[:idx+1]
			dp[idx] = res
		case "Top":
			if n == 1 {
				puts("EMPTY")
				return
			}
			puts(num[n-1])
		case "Size":
			puts(dp[n-1])
		}
	}
	puts("SAFE")
}
