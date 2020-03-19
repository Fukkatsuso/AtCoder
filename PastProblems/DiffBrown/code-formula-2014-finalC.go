package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := strings.Split(next(), " ")
	users := make([]string, 0)
	m := map[string]bool{}
	for i := range s {
		t := strings.Split(s[i], "@")
		for j := 1; j < len(t); j++ {
			if len(t[j]) > 0 && !m[t[j]] {
				m[t[j]] = true
				users = append(users, t[j])
			}
		}
	}

	sort.Sort(sort.StringSlice(users))

	for i := range users {
		puts(users[i])
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
