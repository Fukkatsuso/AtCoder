package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	// sum[i][j]: iから始まりjで終わる数がいくつあるか
	sum := make([][]int, 10)
	for i := range sum {
		sum[i] = make([]int, 10)
	}
	for i := 1; i <= n; i++ {
		x := strconv.Itoa(i)
		sum[x[0]-'0'][x[len(x)-1]-'0']++
	}

	ans := 0
	for i := 1; i <= n; i++ {
		x := strconv.Itoa(i)
		ans += sum[x[len(x)-1]-'0'][x[0]-'0']
	}
	puts(ans)
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

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
