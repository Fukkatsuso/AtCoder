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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func ZAlgorithm(s string) []int {
	n := len(s)
	a := make([]int, n)
	a[0] = n
	for i, j := 1, 0; i < n; {
		for i+j < n && s[j] == s[i+j] {
			j++
		}
		a[i] = j
		if j == 0 {
			i++
			continue
		}
		k := 1
		for i+k < n && k+a[k] < j {
			a[i+k] = a[k]
			k++
		}
		i += k
		j -= k
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, s := nextInt(), next()

	ans := 0
	for i := 0; i < n-1; i++ {
		a := ZAlgorithm(s[i:])
		for j := range a {
			if a[j] <= j {
				ans = max(ans, a[j])
			}
		}
	}

	puts(ans)
}
