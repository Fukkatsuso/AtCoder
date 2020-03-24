package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	name, kit := next(), next()

	cntName, cntKit := make([]int, 26), make([]int, 26)
	for i := 0; i < n; i++ {
		cntName[name[i]-'A']++
	}
	for i := 0; i < m; i++ {
		cntKit[kit[i]-'A']++
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if cntName[i] == 0 {
			continue
		}
		if cntKit[i] == 0 {
			puts(-1)
			return
		}
		ans = max(ans, divFloor(cntName[i], cntKit[i]))
	}

	puts(ans)
}

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

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
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
