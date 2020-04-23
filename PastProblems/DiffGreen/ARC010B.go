package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func daysInMonth(m int) int {
	switch m {
	case 2:
		return 29
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

// 何日目か
func md2days(md string) int {
	s := strings.Split(md, "/")
	month, _ := strconv.Atoi(s[0])
	day, _ := strconv.Atoi(s[1])
	res := 0
	for m := 1; m < month; m++ {
		res += daysInMonth(m)
	}
	res += day
	return res
}

func max(nums []int) int {
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

	n := nextInt()
	holyday := make([]bool, 366+1)
	for i := 0; i < n; i++ {
		days := md2days(next())
		holyday[days] = true
	}

	sum := make([]int, 366+1)
	for i, f := 1, 0; i <= 366; i++ {
		if holyday[i] || i%7 <= 1 || f > 0 {
			sum[i] = sum[i-1] + 1
		} else {
			sum[i] = 0
		}
		// 振替を差し引く
		if f > 0 && !holyday[i] && i%7 > 1 {
			f--
		}
		// 振替を追加
		if holyday[i] && i%7 <= 1 {
			f++
		}
	}

	puts(max(sum))
}
