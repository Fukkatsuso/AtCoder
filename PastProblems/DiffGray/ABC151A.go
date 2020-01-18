package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	puts(string(byte(next()[0]) + 1))
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
