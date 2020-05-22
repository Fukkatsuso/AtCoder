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

func digit2(n int) int {
	d := 0
	for n > 0 {
		n >>= 1
		d++
	}
	return d
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	d := digit2(n)

	name := []string{"Takahashi", "Aoki"}

	turn := 1
	for x := 1; x <= n; turn++ {
		// 最後の桁が来る人
		if (d-turn)%2 == 1 {
			x = 2 * x
		} else {
			x = 2*x + 1
		}
		if x > n {
			break
		}
	}
	puts(name[turn%2])
}
