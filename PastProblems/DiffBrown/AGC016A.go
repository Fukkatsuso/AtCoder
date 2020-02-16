package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	maxCnt := make([]int, 26)
	cnt := make([]int, 26)
	for _, c := range s {
		if cnt[c-'a'] > maxCnt[c-'a'] {
			maxCnt[c-'a'] = cnt[c-'a']
		}
		for j := range cnt {
			cnt[j]++
		}
		cnt[c-'a'] = 0
	}
	for i := range cnt {
		if cnt[i] > maxCnt[i] {
			maxCnt[i] = cnt[i]
		}
	}
	puts(min(maxCnt))
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
