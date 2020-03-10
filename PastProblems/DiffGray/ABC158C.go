package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	a, b := nextInt(), nextInt()
	ans := -1
	for i := 1; i <= 10000; i++ {
		f := float64(i)
		if int(math.Floor(f*0.08)) == a && int(math.Floor(f*0.1)) == b {
			ans = i
			break
		}
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
