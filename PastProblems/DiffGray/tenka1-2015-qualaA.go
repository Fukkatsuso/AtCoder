package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wt = bufio.NewWriter(os.Stdout)
)

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	defer wt.Flush()

	puts(654231)
}
