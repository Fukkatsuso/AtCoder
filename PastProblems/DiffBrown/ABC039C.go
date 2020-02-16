package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	// col := "WBWBWWBWBWBWWBWBWWBWBWBW"
	sounds := []string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}
	bars := []string{
		"WBWBWWBWBWBWWBWBWWBW",
		"WBWWBWBWBWWBWBWWBWBW",
		"WWBWBWBWWBWBWWBWBWBW",
		"WBWBWBWWBWBWWBWBWBWW",
		"WBWBWWBWBWWBWBWBWWBW",
		"WBWWBWBWWBWBWBWWBWBW",
		"WWBWBWWBWBWBWWBWBWWB"}
	for i := range sounds {
		if s == bars[i] {
			puts(sounds[i])
			break
		}
	}
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
