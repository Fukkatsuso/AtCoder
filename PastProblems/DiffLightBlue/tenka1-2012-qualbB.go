// 1つWAして詰んだ
// snake -> camelで大文字を含む単語があるとNGなのを見落としたため

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func between(c, x, y byte) bool {
	return x <= c && c <= y
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	c := next()
	n := len(c)

	// 最初と最後の連続'_'列の長さ
	headUS, tailUS := 0, 0
	for i := 0; i < n && c[i] == '_'; i++ {
		headUS++
	}
	for i := n - 1; i >= 0 && c[i] == '_'; i-- {
		tailUS++
	}
	if headUS+tailUS >= n {
		puts(c)
		return
	}

	// 最初と最後の連続'_'列を除去した文字列
	s := c[headUS : n-tailUS]
	t := ""

	// snake -> camel
	if words := strings.Split(s, "_"); len(words) > 1 {
		for i, word := range words {
			if len(word) == 0 || between(word[0], '0', '9') {
				puts(c)
				return
			}
			for j := range word {
				if between(word[j], 'A', 'Z') {
					puts(c)
					return
				}
			}
			if i == 0 {
				t += word
			} else {
				t += strings.ToUpper(string(word[0])) + word[1:]
			}
		}
		puts(strings.Repeat("_", headUS) + t + strings.Repeat("_", tailUS))
		return
	}

	// camel -> snake
	if !between(s[0], 'a', 'z') {
		puts(c)
		return
	}
	for i := range s {
		if between(s[i], 'A', 'Z') {
			t += "_" + strings.ToLower(string(s[i]))
		} else {
			t += string(s[i])
		}
	}
	puts(strings.Repeat("_", headUS) + t + strings.Repeat("_", tailUS))
}
