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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, R := nextInt(), nextInt()-1
	s := []byte(next())

	ans := 0
	for i := N - 1; i >= 0; i-- {
		if s[i] == 'o' {
			if ans > 0 {
				ans++
			}
			continue
		}
		// 撃つ場所に移動
		if ans > 0 {
			// 移動コストを加算
			to := i
			from := max(0, i-R)
			ans += to - from + 1
		}
		i -= R
		// 撃つコスト
		ans++
		// puts(i, ans)
	}

	puts(ans)
}
