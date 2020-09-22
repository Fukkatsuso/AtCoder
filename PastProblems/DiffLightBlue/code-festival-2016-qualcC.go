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
	mod            = 1000000007
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

func min(nums ...int) int {
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
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	t, a := nextInts(n), nextInts(n)

	topt, topa := make([]bool, n), make([]bool, n)
	topt[0], topa[n-1] = true, true
	for i := 1; i < n; i++ {
		topt[i] = t[i] > t[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		topa[i] = a[i] > a[i+1]
	}

	ans := 1
	for i := 0; i < n; i++ {
		if topt[i] && topa[i] {
			// 標高がt[i]==a[i]で確定するはず
			if t[i] != a[i] {
				ans = 0
				break
			}
		} else if topt[i] {
			// 標高がt[i]で確定するはず
			// t[i]がa[i]を超えているならtopa[i]のはずで矛盾する
			if t[i] > a[i] {
				ans = 0
				break
			}
		} else if topa[i] {
			if t[i] < a[i] {
				ans = 0
				break
			}
		} else {
			ans *= min(t[i], a[i])
			ans %= mod
		}
	}

	puts(ans)
}
