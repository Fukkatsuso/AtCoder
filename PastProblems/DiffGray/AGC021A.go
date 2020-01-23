// 一部解説AC
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := next()
	x, _ := strconv.Atoi(n)
	ans := 0
	if x < 100 {
		for i := 1; i <= x; i++ {
			ans = max(digitsSum(i), ans)
		}
	} else if nines(n) {
		ans = 9*(len(n)-1) + int(n[0]-'0')
	} else {
		ans = int(n[0]-'0') - 1 + 9*(len(n)-1)
	}
	puts(ans)
	wt.Flush()
}

func nines(n string) bool {
	for i := 1; i < len(n); i++ {
		if n[i] != '9' {
			return false
		}
	}
	return true
}

func digitsSum(n int) int {
	ret := 0
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return ret
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

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
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
