// 解説AC

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

	n := nextInt()
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		cnt[nextInt()]++
	}

	ok := (cnt[0] == n) || (len(cnt) == 2 && cnt[0] == n/3)
	if !ok && len(cnt) == 3 {
		ok = true
		nums := make([]int, 0)
		for k, v := range cnt {
			if v != n/3 {
				ok = false
			}
			nums = append(nums, k)
		}
		ok = ok && (nums[0]^nums[1]^nums[2] == 0)
	}

	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
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
